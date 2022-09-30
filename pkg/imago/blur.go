package imago

import (
	"image"
	"image/color"
)

func blurByMatrix(src *Blur, x, y int, m [][]float64) color.Color {
	// var div float64 = float64(src.r*2+1) * float64(src.r*2+1)
	var div float64 = 0
	r, g, b, a := 0., 0., 0., 0.
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			div += m[i][j]
			rr, gg, bb, aa := src.getColor(x-src.r+j, y-src.r+i).RGBA()
			r += m[i][j] * float64(rr)
			g += m[i][j] * float64(gg)
			b += m[i][j] * float64(bb)
			a += m[i][j] * float64(aa)
		}
	}

	if r < 0 {
		r = 0
	}

	if g < 0 {
		g = 0
	}

	if b < 0 {
		b = 0
	}

	if a < 0 {
		a = 0
	}

	out := color.RGBA64{
		R: uint16(r / div),
		G: uint16(g / div),
		B: uint16(b / div),
		A: uint16(a / div),
	}

	return src.ColorModel().Convert(out)
}

func sharpBlur(r int) [][]float64 {
	buf := make([][]float64, r*2+1)
	for i := 0; i < len(buf); i++ {
		buf[i] = make([]float64, r*2+1)
		for j := 0; j < len(buf[i]); j++ {
			buf[i][j] = -1
		}
	}
	buf[r][r] = float64(r*2+1) * float64(r*2+1)
	return buf
}

// var (
// 	normalMatrix = [][]float64{
// 		{0.000789, 0.006581, 0.013347, 0.006581, 0.000789},
// 		{0.006581, 0.054901, 0.111346, 0.054901, 0.006581},
// 		{0.013347, 0.111346, 0.225821, 0.111346, 0.013347},
// 		{0.006581, 0.054901, 0.111346, 0.054901, 0.006581},
// 		{0.000789, 0.006581, 0.013347, 0.006581, 0.000789},
// 	}
// )

// func defaultBlur(r int) [][]float64 {
// 	buf := make([][]float64, r*2+1)
// 	for i := 0; i < len(buf); i++ {
// 		buf[i] = make([]float64, r*2+1)
// 		for j := 0;j < len(buf[i]);j++ {
// 			buf[i][j] =
// 		}
// 	}
// }

type BlurMode string

const (
	Sharp BlurMode = "Sharp"
)

var (
	blurModes = map[BlurMode]func(int) [][]float64{
		Sharp: sharpBlur,
	}
)

type Blur struct {
	r   int
	src image.Image

	Mode BlurMode
}

func (b *Blur) getColor(x, y int) color.Color {
	rect := b.Bounds()

	if x < rect.Min.X {
		x = rect.Min.X
	}

	if x > rect.Max.X {
		x = rect.Max.X
	}

	if y < rect.Min.Y {
		y = rect.Min.Y
	}

	if y > rect.Max.Y {
		y = rect.Max.Y
	}

	return b.src.At(x, y)
}

func (b *Blur) ColorModel() color.Model {
	return b.src.ColorModel()
}

func (b *Blur) Bounds() image.Rectangle {
	return b.src.Bounds()
}

func (b *Blur) At(x, y int) color.Color {
	matrix := blurModes[b.Mode](b.r)
	return blurByMatrix(b, x, y, matrix)
}

func NewBlur(src image.Image, r int) *Blur {
	return &Blur{
		src:  src,
		r:    r,
		Mode: Sharp,
	}
}
