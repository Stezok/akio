package imgtools

import (
	"image"
	"image/color"
)

type Circle struct {
	src image.Image
	r   int
}

func (c *Circle) ColorModel() color.Model {
	return color.RGBAModel
}

func (c *Circle) Bounds() image.Rectangle {
	return c.src.Bounds()
}

func (c *Circle) At(x, y int) color.Color {
	maxPoint := c.src.Bounds().Max
	xx, yy := maxPoint.X/2, maxPoint.Y/2

	if c.r == 0 {
		c.r = xx
	}

	if (xx-x)*(xx-x)+(yy-y)*(yy-y) > c.r*c.r {
		return color.RGBA{0, 0, 0, 0}
	}

	return c.src.At(x, y)
}

func NewCircle(i image.Image) *Circle {
	return &Circle{
		src: i,
		r:   0,
	}
}
