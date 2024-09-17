package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Shoot struct {
	Position rl.Vector2
	IsShooting bool
	Sprite rl.Texture2D
}

func (m *Shoot) AttackOfShoot(p *Player) {
	p.Health -= 5
}