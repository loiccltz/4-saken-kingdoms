package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Monster struct {
	Name     string
	Position rl.Vector2
	Health   int
	MonsterSrc rl.Rectangle
	MonsterDest rl.Rectangle
	MonsterFrame int
	FrameCount int
	Damage   int
	Loot     []item.Item
	Worth    int //valeur en argent quand tué

	IsAlive bool

	Sprite rl.Texture2D
}

func (m *Monster) AttackOfMonster(p *Player) {
	p.Health -= m.Damage
}



func (m *Monster) ToString() {
	fmt.Printf("%d/100 de vie\n", m.Health)
}
