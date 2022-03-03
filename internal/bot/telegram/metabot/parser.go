package metabot

import (
	"NFTProject/internal/meta"
	"strings"
)

// "Текущие атрибуты\nЦвет: %s\nРедкость: %s\nПромо-коллекция: %s

const (
	nameOffset     = len("Имя: ")
	slotOffset     = len("Слот: ")
	colorOffset    = len("Цвет: ")
	rarityOffset   = len("Редкость: ")
	promoTagOffset = len("Промо-коллекция: ")
)

func ParseMetaAttributes(s string) meta.FragmentMetadata {
	data := strings.Split(s, "\n")

	var fragmentMetadata meta.FragmentMetadata
	fragmentMetadata.FileName = data[1][nameOffset:]
	fragmentMetadata.Slot = meta.Slot(data[2][slotOffset:])
	fragmentMetadata.Color = meta.Color(data[3][colorOffset:])
	fragmentMetadata.Rarity = meta.Rarity(data[4][rarityOffset:])
	if len(data) == 6 {
		fragmentMetadata.PromoTag = meta.PromoTag(data[5][promoTagOffset:])
	}

	return fragmentMetadata
}
