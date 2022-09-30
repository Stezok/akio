package imago

import "image/color"

const (
	maxComponent16 = 65535
)

type RGBAf64 struct {
	R, G, B, A float64
}

func (c *RGBAf64) RGBA() (r, g, b, a uint32) {
	return uint32(c.R * maxComponent16), uint32(c.G * maxComponent16), uint32(c.B * maxComponent16), uint32(c.A * maxComponent16)
}

func NewRGBAf64(c color.Color) RGBAf64 {
	r, g, b, a := c.RGBA()

	return RGBAf64{
		R: float64(r) / maxComponent16,
		G: float64(g) / maxComponent16,
		B: float64(b) / maxComponent16,
		A: float64(a) / maxComponent16,
	}
}
