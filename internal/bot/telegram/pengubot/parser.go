package pengubot

import (
	"strconv"
	"strings"
)

const (
	backgroundOffset = len("Фон: ")
	countOffset      = len("Количество: ")
)

func (bot *PenguBot) ParseCustomGenerationData(text string) CustomGenParams {

	data := strings.Split(text, "\n")

	count, _ := strconv.Atoi(data[1][countOffset:])

	return CustomGenParams{
		Count:    count,
		Filename: data[0][backgroundOffset:],
	}
}
