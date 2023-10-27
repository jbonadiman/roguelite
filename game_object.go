package roguelite

import "github.com/hajimehoshi/ebiten/v2"

type GameObject interface {
	Draw(screen *ebiten.Image)
	Update()
	Size() (int, int)
}
