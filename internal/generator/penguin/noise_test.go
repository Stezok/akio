package penguin_test

import (
	"NFTProject/pkg/imago"
	"image"
	"image/png"
	"os"
	"testing"
)

func TestNoise(t *testing.T) {
	f, err := os.Open("../../../tmp/back_Gold1.png")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		t.Error(err)
	}

	w, _ := os.Create("output.png")
	grainNoise := imago.NewGrainNoise(img, 0, 0.1)
	grainNoise.Mode = imago.Overlay
	err = png.Encode(w, grainNoise)
	if err != nil {
		t.Error(err)
	}
}
