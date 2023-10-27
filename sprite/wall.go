package sprite

import "github.com/hajimehoshi/ebiten/v2"

type Wall struct {
	sprite *Sprite
}

func NewWall(x, y int) *Wall {
	return &Wall{sprite: NewSpriteFromSheet(17, 10, x, y)}
}

func (w *Wall) Draw(screen *ebiten.Image) {
	w.sprite.Draw(screen)
}

func (w *Wall) Update() {
}

func (w *Wall) Size() (int, int) {
	return SpritesSize, SpritesSize
}
