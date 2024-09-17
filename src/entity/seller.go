package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Seller struct {
	Position  rl.Vector2
	Money     int
	Inventory []item.Item
	IsAlive   bool
	Sprite    rl.Texture2D
}

func (p *Player) Buy(m *Seller) {
	m.Money += 10
}

func (p *Seller) ToString() {
	fmt.Printf(`
	Joueur: Argent: %d,	Inventaire: %+v	\n`, p.Money, p.Inventory)
}

