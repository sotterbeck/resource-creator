package internal

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"resource-creator/internal/domain"
	"resource-creator/internal/service"
)

// App struct
type App struct {
	ctx context.Context
	img image.Image
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	runtime.EventsOn(a.ctx, "pattern-file-drop", func(optionalData ...interface{}) {
		paths := optionalData[0].([]interface{})
		if len(paths) == 0 {
			return
		}
		path := paths[0].(string)
		file, err := a.createTextureFile("pattern", path)
		if err != nil {
			return
		}
		runtime.EventsEmit(a.ctx, "pattern-file-drop-response", file)
	})
}

var mimeTypes = map[string]string{
	"png":  "image/png",
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"gif":  "image/gif",
}

type TextureFile struct {
	Name    string `json:"name"`
	ImgData string `json:"imgData"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

// OpenTextureFile opens a file dialog to select an image file
// and validates it. If the file is not valid, an error is shown.
// Returns the image as a base64 data URL
func (a *App) OpenTextureFile(validator string) (*TextureFile, error) {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select an image",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Image Files",
				Pattern:     "*.png",
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error opening file path: %s", err)
	}

	if path == "" {
		a.showDialog(runtime.WarningDialog, "No file selected", "Please select an image file to continue.")
		return nil, err
	}

	return a.createTextureFile(validator, path)
}

func (a *App) OpenTextureFiles(validator string) ([]*TextureFile, error) {
	paths, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select multiple images",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Image Files",
				Pattern:     "*.png",
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error opening file paths: %s", err)
	}

	if len(paths) == 0 {
		a.showDialog(runtime.WarningDialog, "No files selected", "Please select one or more image files to continue.")
		return nil, err
	}

	files := make([]*TextureFile, 0)
	for _, path := range paths {
		f, err := a.createTextureFile(validator, path)
		if err != nil {
			a.showDialog(runtime.ErrorDialog, "Error opening file", "An error occurred while opening the file. Please try again.")
			return nil, err
		}
		files = append(files, f)
	}

	return files, err
}

func (a *App) createTextureFile(validator string, path string) (*TextureFile, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	img, format, ok := a.validateImageFile(f, domain.GetImage(validator))
	if !ok {
		a.showDialog(runtime.ErrorDialog, "Invalid Image", fmt.Sprintf("The image %s is not valid. Make sure the selected image resolution is a power of 2.", filepath.Base(path)))
		return nil, err
	}
	a.img = img

	log.Printf("Selected file: %s", path)
	enImg, err := a.encodeImage(path, format)

	return &TextureFile{
		ImgData: enImg,
		Name:    filepath.Base(path),
		Width:   img.Bounds().Dx(),
		Height:  img.Bounds().Dy(),
	}, err
}

func (a *App) validateImageFile(r io.Reader, validator domain.Image) (image.Image, string, bool) {
	img, format, err := image.Decode(r)
	if err != nil {
		return nil, "", false
	}

	bounds := img.Bounds()
	if !validator.IsValidRes(bounds.Dx(), bounds.Dy()) {
		return nil, "", false
	}

	return img, format, true
}

func (a *App) encodeImage(path string, format string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading file: %s", err)
	}
	mime, ok := mimeTypes[format]
	if !ok {
		return "", fmt.Errorf("unsupported image format: %s", format)
	}
	return fmt.Sprintf("data:%s;base64,%s", mime, base64.StdEncoding.EncodeToString(data)), nil
}

// ExportPatternCTM exports all split images and the CTM properties file to the given directory.
func (a *App) ExportPatternCTM(material string, tileRes int) error {
	dir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return err
	}

	if dir == "" {
		return nil
	}

	exporters := make([]service.Exporter, 0)
	exporters = append(exporters,
		service.NewCTMExporter(material, tileRes, domain.Resolution{Width: a.img.Bounds().Dx(), Height: a.img.Bounds().Dy()}),
		service.NewImageExporter(a.img, tileRes),
	)

	for _, exporter := range exporters {
		if err := exporter.Export(dir); err != nil {
			a.showDialog(runtime.ErrorDialog, "Export Error", "See the logs for more information")
			return fmt.Errorf("error exporting: %s", err)
		}
	}
	a.showDialog(runtime.InfoDialog, "Export Complete", "All files have been exported successfully")

	return nil
}

func (a *App) showDialog(typ runtime.DialogType, title, msg string) {
	_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    typ,
		Title:   title,
		Message: msg,
	})
	if err != nil {
		log.Fatalf("error showing dialog: %s", err)
	}
}
