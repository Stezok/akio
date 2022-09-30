package meta

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func ColorToString(clr color.Color) string {
	r, g, b, a := clr.RGBA()
	return fmt.Sprintf("%d, %d, %d, %d", r, g, b, a)
}

var (
	InvalidColorString = errors.New("Invalid color string")
)

func StringToRGBA(strColor string) (color.RGBA, error) {
	colorArr := strings.Split(strColor, ", ")
	if len(colorArr) != 4 {
		return color.RGBA{}, InvalidColorString
	}

	r, err := strconv.Atoi(colorArr[0])
	if err != nil {
		return color.RGBA{}, InvalidColorString
	}

	g, err := strconv.Atoi(colorArr[1])
	if err != nil {
		return color.RGBA{}, InvalidColorString
	}

	b, err := strconv.Atoi(colorArr[2])
	if err != nil {
		return color.RGBA{}, InvalidColorString
	}

	a, err := strconv.Atoi(colorArr[3])
	if err != nil {
		return color.RGBA{}, InvalidColorString
	}

	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}, nil
}

type Bind struct {
	Src string `json:"src"`
	Out string `json:"out"`
}

type OutType string

const (
	OutColor   OutType = "color"
	OutTexture OutType = "texture"
)

func (b *Bind) OutType() OutType {
	if len(b.Out) >= 2 && b.Out[:2] == `\t` {
		return OutTexture
	}

	return OutColor
}

func (b *Bind) GetOutColor() color.RGBA {
	if b.OutType() != OutColor {
		return color.RGBA{}
	}
	clr, _ := StringToRGBA(b.Out)
	return clr
}

type ColorSpec struct {
	Type  string `json:"type"`
	Main  string `json:"main"`
	Binds []Bind `json:"binds"`
}

func (cs *ColorSpec) GetBinded(clr color.Color) Bind {
	colorStr := ColorToString(clr)
	for _, bind := range cs.Binds {
		rgba, _ := StringToRGBA(bind.Src)
		if colorStr == ColorToString(rgba) {
			return bind
		}
	}

	return Bind{}
}

func LoadColorSpec(r io.Reader) (ColorSpec, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return ColorSpec{}, err
	}

	var colorSpec ColorSpec
	err = json.Unmarshal(buf, &colorSpec)
	return colorSpec, err
}
