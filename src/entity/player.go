package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {

	Position  rl.Vector2
	Health    int
	Money     int
	Speed     float32
	Damage    int
	Inventory []item.Item
	//	Movemement []rl.Texture2D
	IsAlive bool

	Sprite rl.Texture2D
}

func (p *Player) AttackOfPlayer(m *Monster) {
	m.Health -= p.Damage
}
func (p *Player) AttackOfPlayerOnMobs(g *Mobs) {
	g.Health -= p.Damage
}
func (p *Player) SpendMoney(s *Seller) {
	if s.Inventory[0].Name == "Biscuit"{
		p.Money -= 5
		p.Inventory = append(p.Inventory, s.Inventory[0])
	}else if s.Inventory[1].Name == "Gateau"{
		p.Money -= 15
		p.Inventory = append(p.Inventory, s.Inventory[1])
	}else if s.Inventory[2].Name == "Bouclier"{
		p.Money -= 25
		p.Inventory = append(p.Inventory, s.Inventory[2])
}}

func (p *Player) ToString() {
	fmt.Printf(`
	Joueur:
		Vie: %d,
		Argent: %d,
		Inventaire: %+v
	
	\n`, p.Health, p.Money, p.Inventory)
}
