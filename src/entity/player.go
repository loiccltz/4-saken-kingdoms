package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)


type Player struct {

	Position  rl.Vector2
	Health    int
	//MaxHealth	int
	//Shield:	int
	//MaxShield: int
	//ShieldRechargeRate: int
	//Endurance	int
	//MaxEndurance	int
	//EnduranceRechargeRate	int
	Money     int
	Speed     float32
	Damage    int
	Inventory []item.Item

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
func (p *Player) UpdateEndurance() {
    // Détermine le montant d'endurance à ajouter par mise à jour
    increment := int(p.EnduranceRechargeRate)

    // Augmente l'endurance en fonction du taux de régénération
    if p.Endurance < p.MaxEndurance {
        p.Endurance += increment
        if p.Endurance > p.MaxEndurance {
            p.Endurance = p.MaxEndurance
        }
    }
}

func (p *Player) UpdateShield() {
    // Détermine le montant de bouclier à ajouter par mise à jour
    increment := int(p.ShieldRechargeRate)

    // Augmente le bouclier en fonction du taux de régénération
    if p.Shield < p.MaxShield {
        p.Shield += increment
        if p.Shield > p.MaxShield {
            p.Shield = p.MaxShield
        }
    }
}
