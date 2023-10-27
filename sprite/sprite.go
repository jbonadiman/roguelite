package sprite

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jbonadiman/roguelite/resources/images"
)

var (
	spriteSheet *ebiten.Image
)

const (
	SpritesSize = 16
)

type Sprite struct {
	X, Y  int
	Image *ebiten.Image
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.Tilemap_png))
	if err != nil {
		log.Fatal(err)
	}
	spriteSheet = ebiten.NewImageFromImage(img)
}

func NewSpriteFromSheet(row, column, x, y int) *Sprite {
	img := spriteSheet.SubImage(
		image.Rect(
			column*SpritesSize,
			row*SpritesSize,
			(column+1)*SpritesSize,
			(row+1)*SpritesSize,
		),
	).(*ebiten.Image)

	return &Sprite{
		X:     x,
		Y:     y,
		Image: img,
	}
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.X), float64(s.Y))
	screen.DrawImage(s.Image, op)
}
