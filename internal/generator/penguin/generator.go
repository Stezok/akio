package penguin

import (
	"NFTProject/internal/generator/penguin/rules"
	"NFTProject/internal/meta"
	"image"
	"image/draw"
	"image/png"
	"io"
	"os"
)

const partCount int = 15

type PenguinGenerator struct {
	*PartManager
}

func (pg *PenguinGenerator) GenerateWithBackground(w io.Writer, bg string) error {

	background := image.NewRGBA64(image.Rect(0, 0, 1440, 1440))
	pengu := image.NewRGBA64(image.Rect(0, 0, 1440, 1440))

	filter := Filter{
		Rules:    rules.PinguinRules{},
		Color:    meta.NoneColor,
		Filename: bg,
	}
	for i := 0; i < len(meta.Slots); i++ {
		filter.Slot = meta.Slots[i]
		frMeta := pg.GetPartFiltred(filter)
		if frMeta == nil {
			continue
		}
		if frMeta.Color != meta.NoneColor {
			filter.Color = frMeta.Color
		}
		imgFile, err := os.Open(pg.GetPartPath(frMeta))
		if err != nil {
			return err
		}

		img, _, err := image.Decode(imgFile)
		if err != nil {
			return err
		}

		if i == 0 {
			colorFocus := 65535 * 0.4
			grainNoise := NewGrainNoise(img, uint16(colorFocus), 0.6)
			grainNoise.Mode = Custom
			draw.Draw(background, background.Bounds(), grainNoise, image.Point{0, 0}, draw.Over)
			filter.Filename = ""
		} else {
			draw.Draw(pengu, pengu.Bounds(), img, image.Point{0, 0}, draw.Over)
		}

	}

	// colorFocus := 65535 * 0.6
	// grainNoise := NewGrainNoise(pengu, uint16(colorFocus), 0.2)
	// grainNoise.Mode = HardLight

	colorFocus := 65535 * 0.7
	grainNoise := NewGrainNoise(pengu, uint16(colorFocus), 0.1)
	grainNoise.Mode = Overlay

	draw.Draw(background, background.Bounds(), grainNoise, image.Point{0, 0}, draw.Over)

	return png.Encode(w, background)
}

func (pg *PenguinGenerator) GenerateRandomSingle(w io.Writer) error {

	background := image.NewRGBA64(image.Rect(0, 0, 1440, 1440))
	pengu := image.NewRGBA64(image.Rect(0, 0, 1440, 1440))

	filter := Filter{
		Rules: rules.PinguinRules{},
		Color: meta.NoneColor,
	}
	for i := 0; i < len(meta.Slots); i++ {
		filter.Slot = meta.Slots[i]
		frMeta := pg.GetPartFiltred(filter)
		if frMeta == nil {
			continue
		}
		if frMeta.Color != meta.NoneColor {
			filter.Color = frMeta.Color
		}
		imgFile, err := os.Open(pg.GetPartPath(frMeta))
		if err != nil {
			return err
		}

		img, _, err := image.Decode(imgFile)
		if err != nil {
			return err
		}

		if i == 0 {
			colorFocus := 65535 * 0.4
			grainNoise := NewGrainNoise(img, uint16(colorFocus), 0.6)
			grainNoise.Mode = Custom
			draw.Draw(background, background.Bounds(), grainNoise, image.Point{0, 0}, draw.Over)
		} else {
			draw.Draw(pengu, pengu.Bounds(), img, image.Point{0, 0}, draw.Over)
		}

	}

	// colorFocus := 65535 * 0.6
	// grainNoise := NewGrainNoise(pengu, uint16(colorFocus), 0.2)
	// grainNoise.Mode = HardLight

	colorFocus := 65535 * 0.7
	grainNoise := NewGrainNoise(pengu, uint16(colorFocus), 0.1)
	grainNoise.Mode = Overlay

	draw.Draw(background, background.Bounds(), grainNoise, image.Point{0, 0}, draw.Over)

	return png.Encode(w, background)
}

func (pg *PenguinGenerator) GeneratePromo(w io.Writer, tag meta.PromoTag) error {
	pengu := image.NewRGBA(image.Rect(0, 0, 1440, 1440))

	filter := Filter{
		Rules:    rules.PinguinRules{},
		PromoTag: tag,
		Color:    meta.NoneColor,
	}
	for i := 0; i < len(meta.Slots); i++ {
		filter.Slot = meta.Slots[i]
		frMeta := pg.GetPartFiltred(filter)
		if frMeta == nil {
			continue
		}
		if frMeta.Color != meta.NoneColor {
			filter.Color = frMeta.Color
		}
		imgFile, err := os.Open(pg.GetPartPath(frMeta))
		if err != nil {
			return err
		}

		img, _, err := image.Decode(imgFile)
		if err != nil {
			return err
		}

		draw.Draw(pengu, pengu.Bounds(), img, image.Point{0, 0}, draw.Over)
	}

	return png.Encode(w, pengu)
}

func NewPinguinGenerator(dataPath string) *PenguinGenerator {
	manager := &PartManager{
		dataPath: dataPath,
	}
	manager.LoadMetadata(dataPath + "/metadata")

	return &PenguinGenerator{
		PartManager: manager,
	}
}
