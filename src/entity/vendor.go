package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Vendor struct {
	Name     string
	Position rl.Vector2
	Health   int
	Damage   int
	Loot     []item.Item
	Worth    int //valeur en argent quand tué

	IsAlive bool

	Sprite rl.Texture2D
}


func (v *Vendor) ToString() {
	fmt.Printf("Je suis un vendeur avec %d points de vie\n", v.Health)
}
