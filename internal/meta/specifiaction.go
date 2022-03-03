package meta

import (
	"encoding/json"
	"fmt"
	"io"
)

type Slot string

const (
	Back          Slot = "Back"
	Shadow        Slot = "Shadow"
	BackHand      Slot = "BackHand"
	Tail          Slot = "Tail"
	Body          Slot = "Body"
	Stom          Slot = "Stom"
	Hair          Slot = "Hair"
	Forehead      Slot = "Forehead"
	Eye           Slot = "Eye"
	Beak          Slot = "Beak"
	EyeAccessory  Slot = "EyeAccessory"
	StomAccessory Slot = "StomAccessory"
	HandAccessory Slot = "HandAccessory"
	Paws          Slot = "Paws"
	Watermark     Slot = "Watermark"
)

var (
	Slots = []Slot{
		Back,
		Shadow,
		BackHand,
		Tail,
		Body,
		Stom,
		Hair,
		Forehead,
		Eye,
		Beak,
		EyeAccessory,
		StomAccessory,
		HandAccessory,
		Paws,
		Watermark,
	}
)

func (s Slot) ID() int {
	for i, slot := range Slots {
		if slot == s {
			return i
		}
	}
	return -1
}

func (s Slot) GetData() string {
	return "slot:" + string(s)
}

type Color string

const (
	NoneColor   Color = "none"
	PinkColor   Color = "pink"
	CyanColor   Color = "cyan"
	RedColor    Color = "red"
	CosmicColor Color = "cosmic"
)

func (c Color) GetData() string {
	return "color:" + string(c)
}

type Rarity string

const (
	Common    Rarity = "common"
	Rare      Rarity = "rare"
	Epic      Rarity = "epic"
	Legendary Rarity = "legendary"
	Promo     Rarity = "promo"
)

func (r Rarity) GetData() string {
	return "rarity:" + string(r)
}

type PromoTag string

const (
	OlympicGame  PromoTag = "olympic_games"
	RickAndMorty PromoTag = "rick_and_morty"
)

var (
	PromoTags = []PromoTag{
		OlympicGame,
		RickAndMorty,
	}
)

func (pt PromoTag) GetData() string {
	return "promotag:" + string(pt)
}

type FragmentMetadata struct {
	FileName string   `json:"filename"`
	Slot     Slot     `json:"slot"`
	Color    Color    `json:"color"`
	Rarity   Rarity   `json:"rarity"`
	PromoTag PromoTag `json:"promoTag"`
}

func (fmd *FragmentMetadata) GetFileName() string {
	return fmt.Sprintf("images/%d/%s", fmd.Slot.ID()+1, fmd.FileName)
}

func (fmd *FragmentMetadata) WriteJson(w io.Writer) error {
	tmp := fmd.FileName
	defer func(s string) { fmd.FileName = s }(tmp)
	fmd.FileName = fmt.Sprintf("%s.png", fmd.FileName)
	data, err := json.Marshal(fmd)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}
