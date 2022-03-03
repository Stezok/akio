package penguin

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

const (
	maxComponent16 = 65535
)

type RGBAf64 struct {
	R, G, B, A float64
}

func (c *RGBAf64) RGBA() (r, g, b, a uint32) {
	return uint32(c.R * maxComponent16), uint32(c.G * maxComponent16), uint32(c.B * maxComponent16), uint32(c.A * maxComponent16)
}

func color2f64(c color.Color) RGBAf64 {
	r, g, b, a := c.RGBA()

	return RGBAf64{
		R: float64(r) / maxComponent16,
		G: float64(g) / maxComponent16,
		B: float64(b) / maxComponent16,
		A: float64(a) / maxComponent16,
	}
}

func blendPerChannel(dst, src color.Color, bf func(float64, float64) float64) color.Color {
	fdst, fsrc := color2f64(dst), color2f64(src)
	return &RGBAf64{
		R: bf(fdst.R, fsrc.R),
		G: bf(fdst.G, fsrc.G),
		B: bf(fdst.B, fsrc.B),
		// A: bf(fdst.A, fsrc.A),
		A: fdst.A,
	}
}

func ColorBurnBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		return 1 - (1-d)/s
	})
}

func VividLightBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		if s > 0.5 {
			return d / (1 - 2*(s-0.5))
		} else {
			return 1 - (1-d)/(s*2)
		}
	})
}

func MultiplyBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		return d * s
	})
}

func OverlayBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		if d < 0.5 {
			return 2 * d * s
		} else {
			return 1 - 2*(1-d)*(1-s)
		}
	})
}

func SoftLightBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		if s < 0.5 {
			return 2*d*s + d*d*(1-2*s)
		} else {
			return 2*d*(1-s) + math.Sqrt(d)*(2*s-1)
		}
	})
}

func InvSoftLightBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		if d < 0.5 {
			return 2*s*d + s*s*(1-2*d)
		} else {
			return 2*s*(1-d) + math.Sqrt(s)*(2*d-1)
		}
	})
}

func PegtopSoftLightBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		return (1-2*s)*d*d + 2*s*d
	})
}

func ScreenBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		return 1 - (1-d)*(1-s)
	})
}

func HardLightBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		if s > 0.5 {
			return 1 - (1-d)*(1-(s-0.5))
		}
		return d * (s + 0.5)
	})
}

func LinearBurnBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		return d + s - 1
	})
}

func CustomBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		if s > 0.5 && d < 0.5 {
			return 2 * d * s
		}

		if s > 0.5 {
			return 1 - (1-d)*(1-(s-0.5))
		}
		return d * (s + 0.5)
	})
}

type BlendMode string

const (
	ColorBurn       BlendMode = "ColorBurn"
	VividLight      BlendMode = "VividLight"
	Multiply        BlendMode = "Multiply"
	Overlay         BlendMode = "Overlay"
	SoftLight       BlendMode = "SoftLight"
	InvSoftLight    BlendMode = "InvSoftLight"
	PegtopSoftLight BlendMode = "PegtopSoftLight"
	HardLight       BlendMode = "HardLight"
	Screen          BlendMode = "Screen"
	LinearBurn      BlendMode = "LinearBurn"
	Custom          BlendMode = "Custom"
)

var (
	modes = map[BlendMode]func(color.Color, color.Color) color.Color{
		ColorBurn:       ColorBurnBlend,
		VividLight:      VividLightBlend,
		Multiply:        MultiplyBlend,
		Overlay:         OverlayBlend,
		SoftLight:       SoftLightBlend,
		InvSoftLight:    InvSoftLightBlend,
		PegtopSoftLight: PegtopSoftLightBlend,
		HardLight:       HardLightBlend,
		Screen:          ScreenBlend,
		LinearBurn:      LinearBurnBlend,
		Custom:          CustomBlend,
	}
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
	noise := image.NewGray16(rect)
	for y := 0; y < rect.Dy(); y++ {
		for x := 0; x < rect.Dx(); x++ {
			grainChance := rand.Intn(100)
			if grainChance < int(gn.grainFreq*100) {
				noise.Set(x, y, randomNoisePixelGray16(gn.colorFocus))
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

	grainChance := rand.Intn(1000)
	if grainChance < int(1000*gn.grainFreq) {
		// noiseColor := randomNoisePixelGray16(gn.colorFocus)
		noiseColor := randomNoisePixelColor()
		return modes[gn.Mode](gn.src.At(x, y), noiseColor)
	}
	return gn.src.At(x, y)
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
