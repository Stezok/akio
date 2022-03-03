package generator

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"os"
)

type Pixel struct {
	Point image.Point
	Color color.Color
}

func CoinFlip() bool {
	return rand.Intn(2) == 1
}

func Chance(chance int) bool {
	return rand.Intn(100) < chance
}

func TestPixels(pixels []*Pixel) error {
	var uniqColors []color.Color
	for _, pixel := range pixels {

		var was bool = false
		for _, color := range uniqColors {
			if color == pixel.Color {
				was = true
				break
			}
		}

		if !was {
			uniqColors = append(uniqColors, pixel.Color)
		}
	}

	file, err := os.Create("pixelTestResult")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, color := range uniqColors {
		fmt.Fprintf(file, "%+v\n", color)
	}

	return nil
}

func DecodePixelsFromImage(img image.Image, offsetX, offsetY int) []*Pixel {
	pixels := []*Pixel{}
	for y := 0; y <= img.Bounds().Max.Y; y++ {
		for x := 0; x <= img.Bounds().Max.X; x++ {
			p := &Pixel{
				Point: image.Point{x + offsetX, y + offsetY},
				Color: img.At(x, y),
			}
			pixels = append(pixels, p)
		}
	}
	return pixels
}

func IsClearPixel(pixel *Pixel) bool {
	r, g, b, a := pixel.Color.RGBA()
	return r == 0 && g == 0 && b == 0 && a == 0
	// return (r == 0 && g == 0 && b == 0) || a <= 128
}

func MergePNGPixels(img *image.RGBA, pixels []*Pixel) {
	for _, pixel := range pixels {
		if !IsClearPixel(pixel) {
			// log.Print(pixel)
			img.Set(
				pixel.Point.X,
				pixel.Point.Y,
				pixel.Color,
			)
		}
	}
}
