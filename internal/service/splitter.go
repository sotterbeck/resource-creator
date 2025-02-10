package service

import (
	"fmt"
	"image"
	"log"
	"resource-creator/internal/domain"
)

type ImageSplitter struct {
	validator domain.Image
}

func (s *ImageSplitter) SplitImage(img image.Image, res int) ([]image.Image, error) {
	if !s.validator.IsValidRes(img.Bounds().Dx(), img.Bounds().Dy()) {
		return nil, fmt.Errorf("invalid source image resolution: %d x %d", img.Bounds().Dx(), img.Bounds().Dy())
	}

	if res > img.Bounds().Dx() || res > img.Bounds().Dy() ||
		res <= 0 ||
		img.Bounds().Dx()%res != 0 || img.Bounds().Dy()%res != 0 {
		return nil, fmt.Errorf("invalid split resolution: %dx", res)
	}

	var imgs []image.Image
	for y := 0; y < img.Bounds().Dy(); y += res {
		for x := 0; x < img.Bounds().Dx(); x += res {
			imgs = append(imgs, img.(interface {
				SubImage(r image.Rectangle) image.Image
			}).SubImage(image.Rect(x, y, x+res, y+res)))
		}
	}

	log.Printf("Split image into %d parts", len(imgs))
	return imgs, nil
}
