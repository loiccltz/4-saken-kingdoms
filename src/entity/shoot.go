package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Shoot struct {
	Position   rl.Vector2
	IsShooting bool
	Direction  int // 0=haut, 1=bas, 2=gauche, 3=droite
	Damage int
	Sprite     rl.Texture2D
}

func (m *Shoot) AttackOfShoot(p *Player) {
	p.Health -= m.Damage
}