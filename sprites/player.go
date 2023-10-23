package sprites

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	sprite *Sprite
}

func NewPlayer(x, y int) *Player {
	return &Player{sprite: NewSpriteFromSheet(27, 1, x, y)}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.sprite.Draw(screen)
}

func (p *Player) Update() {
}
