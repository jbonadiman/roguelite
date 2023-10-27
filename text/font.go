package text

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	dpi = 72
)

type Font struct {
	font font.Face
}

func loadFontFromFile(fontData []byte, options opentype.FaceOptions) (
	font.Face,
	error,
) {
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

func NewFont(size float64, fontData []byte) (Font, error) {
	newFont, err := loadFontFromFile(
		fontData, opentype.FaceOptions{
			Size:    size,
			DPI:     dpi,
			Hinting: font.HintingFull,
		},
	)
	if err != nil {
		return Font{}, err
	}

	return Font{
		font: newFont,
	}, nil
}

func (f *Font) Printf(
	color color.Color,
	format string,
	a ...any,
) *ebiten.Image {
	formattedText := fmt.Sprintf(format, a...)

	rect, _ := font.BoundString(f.font, formattedText)
	sizeX := rect.Max.X.Ceil() - rect.Min.X.Floor()
	sizeY := rect.Max.Y.Ceil() - rect.Min.Y.Floor()

	textImage := ebiten.NewImage(sizeX, sizeY)

	op := &ebiten.DrawImageOptions{}

	op.ColorScale.ScaleWithColor(color)
	op.GeoM.Translate(0, float64(sizeY))
	op.Filter = ebiten.FilterNearest

	text.DrawWithOptions(textImage, formattedText, f.font, op)

	return textImage
}
