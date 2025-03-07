package service

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"resource-creator/internal/domain"
	"time"
)

type CTMExporter struct {
	material   string
	tileRes    int
	patternRes domain.Resolution
}

func NewCTMExporter(material string, tileRes int, patternRes domain.Resolution) *CTMExporter {
	return &CTMExporter{
		material:   material,
		tileRes:    tileRes,
		patternRes: patternRes,
	}
}

func (ce *CTMExporter) Export(dir string) error {
	if err := ce.checkRepeatResolution(); err != nil {
		return err
	}

	_, name, err := domain.ParseMaterial(ce.material)
	if err != nil {
		return fmt.Errorf("failed to parse material: %v", err)
	}

	filePath := filepath.Join(dir, fmt.Sprintf("block_%s.properties", name))
	f, err := os.Create(filePath)
	defer f.Close()

	width, height, tileAmount := ce.calculateTileAmount()
	ctm, err := domain.NewCTMProps("repeat", 0, tileAmount-1, map[string]string{
		"width":  fmt.Sprintf("%d", width),
		"height": fmt.Sprintf("%d", height),
	})

	if err != nil {
		return err
	}

	pw := domain.NewPropertiesWriter(ctm.GetProps())
	for _, c := range ce.getComments() {
		pw.AddComment(c)
	}
	if _, err := pw.WriteTo(f); err != nil {
		return fmt.Errorf("failed to write properties: %v", err)
	}

	return nil
}

func (ce *CTMExporter) getComments() []string {
	t := time.Now()
	return []string{
		"Auto-generated properties file by Resource Creator (https://github.com/sotterbeck/resource-creator)",
		"On " + t.Format(time.RFC1123),
	}
}

func (ce *CTMExporter) checkRepeatResolution() error {
	if ce.tileRes > ce.patternRes.Width || ce.tileRes > ce.patternRes.Height {
		return fmt.Errorf("tile resolution is larger than pattern resolution")
	}

	tileValidator := domain.TileImage{}
	if !tileValidator.IsValidRes(ce.tileRes, ce.tileRes) {
		return fmt.Errorf("invalid tile resolution: %d x %d", ce.tileRes, ce.tileRes)
	}

	patternValidator := domain.PatternImage{}
	if !patternValidator.IsValidRes(ce.patternRes.Width, ce.patternRes.Height) {
		return fmt.Errorf("invalid pattern resolution: %d x %d", ce.patternRes.Width, ce.patternRes.Height)
	}
	return nil
}

func (ce *CTMExporter) calculateTileAmount() (w int, h int, len int) {
	w = ce.patternRes.Width / ce.tileRes
	h = ce.patternRes.Height / ce.tileRes
	len = w * h
	return
}

type ImageExporter struct {
	img     image.Image
	tileRes int
}

func NewImageExporter(img image.Image, tileRes int) *ImageExporter {
	return &ImageExporter{
		img:     img,
		tileRes: tileRes,
	}
}

func (ie *ImageExporter) Export(dir string) error {
	is := ImageSplitter{validator: &domain.PatternImage{}}
	imgs, err := is.SplitImage(ie.img, ie.tileRes)
	if err != nil {
		return err
	}

	for i, img := range imgs {
		filePath := filepath.Join(dir, fmt.Sprintf("%d.png", i))
		f, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		if err := png.Encode(f, img); err != nil {
			return fmt.Errorf("failed to encode image: %v", err)
		}
		f.Close()
	}

	return nil
}
