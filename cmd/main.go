package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jbonadiman/roguelite"
	"github.com/jbonadiman/roguelite/resources/fonts"
	"github.com/jbonadiman/roguelite/text"
)

var (
	firstLevel = [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 2, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	helloWorldText *ebiten.Image
)

type Game struct {
	keys     []ebiten.Key
	level    *roguelite.Level
	gameFont text.Font
}

func init() {}

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

	g.level.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 150)

	screen.DrawImage(helloWorldText, op)
	g.level.Draw(screen)
}

func (g *Game) Layout(int, int) (screenWidth int, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("roguelite")

	snapJackFont, err := text.NewFont(12, fonts.Snapjack_otf)
	if err != nil {
		log.Fatal(err)
	}

	helloWorldText = snapJackFont.Printf(color.White, "Hello World!")

	game := &Game{
		level:    roguelite.NewLevelFromMatrix(firstLevel),
		gameFont: snapJackFont,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
