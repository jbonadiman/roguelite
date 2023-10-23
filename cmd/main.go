package main

import (
	"fmt"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jbonadiman/roguelite"
	"github.com/jbonadiman/roguelite/resources/fonts"
	"github.com/jbonadiman/roguelite/sprites"

	_ "image/png"
)

const (
	MAP_WIDTH  = 10
	MAP_HEIGHT = 8
)

var (
	snapJackFont font.Face
	textImage    = ebiten.NewImage(400, 300)
	player       *sprites.Player
	gameMap      = []int{
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	}
	gameObjects []roguelite.GameObject
)

type Game struct {
	keys []ebiten.Key
}

func loadFontFromFile(fontData []byte, options opentype.FaceOptions) (font.Face, error) {
	tt, err := opentype.Parse(fontData)
	if err != nil {
		return nil, err
	}

	loadedFont, err := opentype.NewFace(tt, &options)
	if err != nil {
		return nil, err
	}

	return loadedFont, nil
}

func init() {
	var err error

	snapJackFont, err = loadFontFromFile(fonts.Snapjack_otf, opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < MAP_HEIGHT; i++ {
		for j := 0; j < MAP_WIDTH; j++ {
			if gameMap[i*MAP_WIDTH+j] == 1 {
				gameObjects = append(
					gameObjects,
					sprites.NewWall(i*sprites.SPRITES_SIZE, j*sprites.SPRITES_SIZE))
			}
		}
	}

	player = sprites.NewPlayer(20, 20)
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	var keyStrs []string
	for _, k := range g.keys {
		keyStrs = append(keyStrs, k.String())
		fmt.Println(keyStrs)
	}

	// if len(keyStrs) > 0 {
	// 	if keyStrs[0] == "ArrowUp" {
	// 		y -= 10
	// 	} else if keyStrs[0] == "ArrowDown" {
	// 		y += 10
	// 	} else if keyStrs[0] == "ArrowLeft" {
	// 		x -= 10
	// 	} else if keyStrs[0] == "ArrowRight" {
	// 		x += 10
	// 	}
	// }

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// text.Draw(textImage, "Hello Ebiten!", snapJackFont, 20, 40, color.White)

	for _, obj := range gameObjects {
		obj.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
