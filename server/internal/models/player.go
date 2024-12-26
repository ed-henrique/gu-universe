package models

import "math"

type Player struct {
	X           uint8
	Y           uint8
	Health      uint8
	Mana        uint8
	recoverFunc func(health, mana uint8) (newHealth, newMana uint8)
}

func DefaultRegenFunc(health, mana uint8) (newHealth, newMana uint8) {
	if health+5 >= math.MaxUint8 {
		newHealth = math.MaxUint8
	} else {
		newHealth = health + 5
	}

	if mana+5 >= math.MaxUint8 {
		newMana = math.MaxUint8
	} else {
		newMana = mana + 5
	}

	return
}

func NewPlayer(recoverFunc func(health, mana uint8) (uint8, uint8)) Player {
	p := Player{
		X:      0,
		Y:      0,
		Health: 255,
		Mana:   255,
	}

	if recoverFunc == nil {
		p.recoverFunc = DefaultRegenFunc
	} else {
		p.recoverFunc = recoverFunc
	}

	return p
}

func (p *Player) SetPosition(x, y uint8) {
	p.X = x
	p.Y = y
}

func (p *Player) Recover() {
	p.Health, p.Mana = p.recoverFunc(p.Health, p.Mana)
}
