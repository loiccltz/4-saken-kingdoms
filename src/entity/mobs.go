package entity

import (
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Mobs struct {
	Name     string
	Position rl.Vector2
	Health   int
	Damage   int
	Loot     []item.Item
	Worth    int //valeur en argent quand tué

	IsAlive bool

	Sprite rl.Texture2D
}
func (g *Mobs) Attack(p *Player) {
	p.Health -= 10
}