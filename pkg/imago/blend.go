package imago

import (
	"image/color"
	"math"
)

func blendPerChannel(dst, src color.Color, bf func(float64, float64) float64) color.Color {
	fdst, fsrc := NewRGBAf64(dst), NewRGBAf64(src)
	return &RGBAf64{
		R: bf(fdst.R, fsrc.R),
		G: bf(fdst.G, fsrc.G),
		B: bf(fdst.B, fsrc.B),
		A: fdst.A,
	}
}

func colorBurnBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		return 1 - (1-d)/s
	})
}

func vividLightBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		if s > 0.5 {
			return d / (1 - 2*(s-0.5))
		} else {
			return 1 - (1-d)/(s*2)
		}
	})
}

func multiplyBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		return d * s
	})
}

func overlayBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		if d < 0.5 {
			return 2 * d * s
		} else {
			return 1 - 2*(1-d)*(1-s)
		}
	})
}

func softLightBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		if s < 0.5 {
			return 2*d*s + d*d*(1-2*s)
		} else {
			return 2*d*(1-s) + math.Sqrt(d)*(2*s-1)
		}
	})
}

func invSoftLightBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		if d < 0.5 {
			return 2*s*d + s*s*(1-2*d)
		} else {
			return 2*s*(1-d) + math.Sqrt(s)*(2*d-1)
		}
	})
}

func pegtopSoftLightBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		return (1-2*s)*d*d + 2*s*d
	})
}

func screenBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		return 1 - (1-d)*(1-s)
	})
}

func hardLightBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		if s > 0.5 {
			return 1 - (1-d)*(1-(s-0.5))
		}
		return d * (s + 0.5)
	})
}

func linearBurnBlend(dst, src color.Color) color.Color {
	return blendPerChannel(dst, src, func(d, s float64) float64 {
		return d + s - 1
	})
}

func customBlend(dst, src color.Color) color.Color {
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
	blendModes = map[BlendMode]func(color.Color, color.Color) color.Color{
		ColorBurn:       colorBurnBlend,
		VividLight:      vividLightBlend,
		Multiply:        multiplyBlend,
		Overlay:         overlayBlend,
		SoftLight:       softLightBlend,
		InvSoftLight:    invSoftLightBlend,
		PegtopSoftLight: pegtopSoftLightBlend,
		HardLight:       hardLightBlend,
		Screen:          screenBlend,
		LinearBurn:      linearBurnBlend,
		Custom:          customBlend,
	}
)
