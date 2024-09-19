package entity

import (
	"main/src/item"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Mobs struct {
	Name           string
	Position       rl.Vector2
	Health         int
	Damage         int
	Loot           []item.Item
	Worth          int
	CoolDown       time.Duration
	IsAlive        bool
	Sprite         rl.Texture2D
	LastAttackTime time.Time
}

func (g *Mobs) Attack(p *Player) {
	p.Health -= g.Damage
}
