package imago

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/rand"
)

func randomNoisePixelGray16(mid int) color.Gray16 {
	rv := rand.Intn(maxComponent16 / 4)

	if rv >= maxComponent16/8 {
		rv = maxComponent16/8 - rv
	}

	v := uint16((mid + rv) % maxComponent16)

	return color.Gray16{
		Y: v,
	}
}

func randomNoisePixelColor() color.RGBA64 {
	r, g, b := rand.Intn(maxComponent16+1), rand.Intn(maxComponent16+1), rand.Intn(maxComponent16+1)
	return color.RGBA64{uint16(r), uint16(g), uint16(b), maxComponent16}
}

type GrainNoise struct {
	colorFocus int
	grainFreq  float64

	src   image.Image
	noise image.Image

	Mode BlendMode
}

func (gn *GrainNoise) generateNoiseImage(rect image.Rectangle) {
	noise := image.NewRGBA(rect)
	for y := 0; y < rect.Dy(); y++ {
		for x := 0; x < rect.Dx(); x++ {
			grainChance := rand.Intn(100)
			if grainChance < int(gn.grainFreq*100) {
				noise.Set(x, y, randomNoisePixelColor())
			}
		}
	}
	gn.noise = noise
}

func (gn *GrainNoise) ColorModel() color.Model {
	return color.RGBAModel
}

func (gn *GrainNoise) Bounds() image.Rectangle {
	return gn.src.Bounds()
}

func (gn *GrainNoise) At(x, y int) color.Color {

	if _, _, _, a := gn.src.At(x, y).RGBA(); a != maxComponent16 {
		return gn.src.At(x, y)
	}

	noiseColor := gn.noise.At(x, y)
	if _, _, _, a := noiseColor.RGBA(); a != maxComponent16 {
		return gn.src.At(x, y)
	}
	return blendModes[gn.Mode](gn.src.At(x, y), noiseColor)
}

func NewGrainNoise(src image.Image, colorFocus uint16, freq float64) *GrainNoise {
	grainNoise := &GrainNoise{
		colorFocus: int(colorFocus),
		grainFreq:  freq,
		src:        src,
	}
	grainNoise.generateNoiseImage(src.Bounds())
	return grainNoise
}

func NewGrainNoiseFromReader(r io.Reader, freq float64) (*GrainNoise, error) {
	img, err := png.Decode(r)
	if err != nil {
		return nil, err
	}
	return NewGrainNoise(img, 0, freq), nil
}
