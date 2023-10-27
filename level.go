package roguelite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jbonadiman/roguelite/sprite"
)

const (
	Wall   = 1
	Player = 2
)

type Level struct {
	objects       []GameObject
	width, height int
}

func NewLevelFromVector(levelRepresentation []int, width int) *Level {
	height := len(levelRepresentation) / width

	gameObjects := make([]GameObject, 0)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			x, y := j*sprite.SpritesSize, i*sprite.SpritesSize
			var obj GameObject

			switch levelRepresentation[i*width+j] {
			case Wall:
				obj = sprite.NewWall(x, y)
			case Player:
				obj = sprite.NewPlayer(x, y)
			}

			if obj != nil {
				gameObjects = append(gameObjects, obj)
			}
		}
	}
	return &Level{
		objects: gameObjects,
		width:   width,
		height:  height,
	}
}

func NewLevelFromMatrix(levelRepresentation [][]int) *Level {
	gameObjects := make([]GameObject, 0)
	for i := 0; i < len(levelRepresentation); i++ {
		for j := 0; j < len(levelRepresentation[i]); j++ {
			x, y := j*sprite.SpritesSize, i*sprite.SpritesSize
			var obj GameObject

			switch levelRepresentation[i][j] {
			case Wall:
				obj = sprite.NewWall(x, y)
			case Player:
				obj = sprite.NewPlayer(x, y)
			}

			if obj != nil {
				gameObjects = append(gameObjects, obj)
			}
		}
	}
	return &Level{
		objects: gameObjects,
		width:   len(levelRepresentation[0]),
		height:  len(levelRepresentation),
	}
}

func (l *Level) Draw(screen *ebiten.Image) {
	for _, obj := range l.objects {
		obj.Draw(screen)
	}
}

func (l *Level) Update() {
	for _, obj := range l.objects {
		obj.Update()
	}
}

func (l *Level) Size() (int, int) {
	return l.width, l.height
}
