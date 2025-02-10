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
}

var mimeTypes = map[string]string{
	"png":  "image/png",
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"gif":  "image/gif",
}

type OpenTextureFileResp struct {
	Name    string `json:"name"`
	ImgData string `json:"imgData"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

// OpenTextureFile opens a file dialog to select an image file
// and validates it. If the file is not valid, an error is shown.
// Returns the image as a base64 data URL
func (a *App) OpenTextureFile() (*OpenTextureFileResp, error) {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select an image",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Image Files",
				Pattern:     "*.png;*.jpg;*.jpeg;*.gif;*.bmp",
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error opening file path: %s", err)
	}

	if path == "" {
		a.showErrorDialog("Select an image file to continue")
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	img, format, ok := a.validateImageFile(f, &domain.PatternImage{})
	if !ok {
		a.showErrorDialog("Image file is not valid, make sure it's a power of 2 resolution")
		return nil, err
	}
	a.img = img

	log.Printf("Selected file: %s", path)
	enImg, err := a.encodeImage(path, format)

	return &OpenTextureFileResp{
		ImgData: enImg,
		Name:    f.Name(),
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
		a.showErrorDialog("Could not open directory")
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
			a.showErrorDialog("An error occurred while exporting. See logs for more information")
			return fmt.Errorf("error exporting: %s", err)
		}
	}
	a.showInfoDialog("Export successful")

	return nil
}

func (a *App) showErrorDialog(msg string) {
	_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:  runtime.ErrorDialog,
		Title: msg,
	})
	if err != nil {
		log.Fatalf("error showing dialog: %s", err)
	}
}

func (a *App) showInfoDialog(msg string) {
	_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:  runtime.InfoDialog,
		Title: msg,
	})
	if err != nil {
		log.Fatalf("error showing dialog: %s", err)
	}
}
