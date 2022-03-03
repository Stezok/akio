package rules

import (
	"NFTProject/internal/generator"
	"NFTProject/internal/meta"
)

type PinguinRules struct {
}

type Action int

const (
	Pass Action = iota
	Continue
)

func (pr *PinguinRules) PassWithChance(chance int) Action {
	if generator.Chance(chance) {
		return Pass
	} else {
		return Continue
	}
}

func (pr *PinguinRules) DecidePass(slot meta.Slot) Action {
	return Pass
}

func (pr *PinguinRules) Decide(slot meta.Slot) Action {
	switch slot {
	case meta.EyeAccessory:
		return pr.PassWithChance(20)
	case meta.StomAccessory:
		return pr.PassWithChance(50)
	case meta.HandAccessory:
		return pr.PassWithChance(20)
	case meta.Tail:
		return pr.PassWithChance(15)
	case meta.Forehead:
		return pr.PassWithChance(33)
	default:
		return Pass
	}
}
