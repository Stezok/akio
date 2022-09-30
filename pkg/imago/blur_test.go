package imago_test

import (
	"NFTProject/pkg/imago"
	"image/jpeg"
	"image/png"
	"os"
	"testing"
)

func TestBlur(t *testing.T) {
	file, err := os.Open("C:/Users/Админ/Desktop/nft/project/rat.jpeg")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		t.Error(err)
	}

	blur := imago.NewBlur(img, 1)
	blur.Mode = imago.Sharp

	out, err := os.Create("output.png")
	if err != nil {
		t.Error(err)
	}
	defer out.Close()
	err = png.Encode(out, blur)
	if err != nil {
		t.Error(err)
	}
}
