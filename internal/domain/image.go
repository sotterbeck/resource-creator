package domain

type Image interface {
	IsValidRes(w int, h int) bool
}

type Resolution struct {
	Width  int
	Height int
}

var images = map[string]Image{
	"pattern": &PatternImage{},
	"tile":    &TileImage{},
}

func GetImage(name string) Image {
	if img, ok := images[name]; ok {
		return img
	}
	return &TileImage{}
}

type PatternImage struct {
}

func (p *PatternImage) IsValidRes(w int, h int) bool {
	return isPowerOfTwo(w) && isPowerOfTwo(h) && w > 0 && h > 0
}

type TileImage struct {
}

func (t *TileImage) IsValidRes(w int, h int) bool {
	return w > 0 && h > 0 && w == h && isPowerOfTwo(w)
}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
