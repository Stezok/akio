package clrchanger

import (
	"NFTProject/internal/meta"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type ColorProxy struct {
	src  image.Image
	spec *meta.ColorSpec

	textures map[string]image.Image
}

func (cp *ColorProxy) SetSpec(spec *meta.ColorSpec) error {
	cp.spec = spec
	log.Printf("%+v", cp.spec)
	textures := make(map[string]image.Image)
	for _, bind := range spec.Binds {
		if bind.OutType() == meta.OutTexture {
			file, err := os.Open(bind.Out[2:])
			if err != nil {
				return err
			}
			img, err := png.Decode(file)
			if err != nil {
				file.Close()
				return err
			}
			file.Close()
			textures[bind.Src] = img
		}
	}
	cp.textures = textures
	return nil
}

func (cp *ColorProxy) ColorModel() color.Model {
	return cp.src.ColorModel()
}

func (cp *ColorProxy) Bounds() image.Rectangle {
	return cp.src.Bounds()
}

func (cp *ColorProxy) At(x, y int) color.Color {
	srcColor := cp.src.At(x, y)
	if _, _, _, a := srcColor.RGBA(); a == 0 {
		return srcColor
	}

	bind := cp.spec.GetBinded(srcColor)
	if bind.Src == "" {
		// clr, _ := meta.StringToRGBA(cp.spec.Main)
		// return clr
		return srcColor
	}

	if bind.OutType() == meta.OutTexture {
		return cp.textures[bind.Src].At(x, y)
	} else {
		return bind.GetOutColor()
	}
}

func NewColorProxy(src image.Image) *ColorProxy {
	return &ColorProxy{
		src: src,
	}
}
