package imgtools

import (
	"image"
	"os"

	"github.com/fogleman/gg"
)

func MakeHelloImage(userImage image.Image, text string) (image.Image, error) {
	circle := NewCircle(userImage)

	bgFile, err := os.Open("assets/discord/background.png")
	if err != nil {
		return nil, err
	}
	defer bgFile.Close()

	background, _, err := image.Decode(bgFile)
	if err != nil {
		return nil, err
	}

	frameFile, err := os.Open("assets/discord/frame.png")
	if err != nil {
		return nil, err
	}
	defer frameFile.Close()

	frame, _, err := image.Decode(frameFile)
	if err != nil {
		return nil, err
	}

	ctx := gg.NewContextForImage(background)

	ctx.DrawImageAnchored(circle, 375, 95, 0.5, 0.5)
	ctx.DrawImageAnchored(frame, 375, 95, 0.5, 0.5)
	err = ctx.LoadFontFace("assets/discord/fonts/Amatic-Bold.ttf", 40)
	if err != nil {
		return nil, err
	}
	ctx.SetRGB(0, 0, 0)

	ctx.DrawStringAnchored(text, 375, 200, 0.5, 0.5)
	return ctx.Image(), nil
}
