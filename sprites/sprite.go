package sprites

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jbonadiman/roguelite/resources/images"
)

var (
	spritesheet *ebiten.Image
)

const (
	SPRITES_SIZE = 16
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
	spritesheet = ebiten.NewImageFromImage(img)
}

func NewSpriteFromSheet(row, column, x, y int) *Sprite {
	image := spritesheet.SubImage(
		image.Rect(
			(column * SPRITES_SIZE),
			(row * SPRITES_SIZE),
			((column + 1) * SPRITES_SIZE),
			((row + 1) * SPRITES_SIZE))).(*ebiten.Image)

	return &Sprite{
		X:     x,
		Y:     y,
		Image: image,
	}
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.X), float64(s.Y))
	screen.DrawImage(s.Image, op)
}
