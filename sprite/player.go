package sprite

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	sprite *Sprite
}

func NewPlayer(x, y int) *Player {
	return &Player{sprite: NewSpriteFromSheet(0, 25, x, y)}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.sprite.Draw(screen)
}

func (p *Player) Update() {
}

func (p *Player) Size() (int, int) {
	return SpritesSize, SpritesSize
}
