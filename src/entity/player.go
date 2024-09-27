package entity

import (
	"main/src/item"                  // Importation du package pour gérer les objets dans l'inventaire du joueur
	rl "github.com/gen2brain/raylib-go/raylib" // Importation de Raylib pour les fonctionnalités graphiques
)

// Player : Structure représentant un joueur dans le jeu
type Player struct {
	Position              rl.Vector2      // Position du joueur sur l'écran, représentée par un vecteur 2D
	Health                int             // Points de vie du joueur

	PlayerSrc             rl.Rectangle     // Source de la texture du joueur, utilisée pour l'animation
	PlayerDest            rl.Rectangle     // Destination de la texture du joueur pour le rendu
	PlayerFrame           int              // Indice du cadre actuel pour l'animation
	FrameCount            int              // Nombre total de cadres dans l'animation

	MaxHealth             int              // Points de vie maximum du joueur
	Shield                int              // Points de bouclier actuels du joueur
	MaxShield             int              // Points de bouclier maximum du joueur
	ShieldRechargeRate    int              // Taux de régénération du bouclier par mise à jour
	Endurance             int              // Points d'endurance actuels du joueur
	MaxEndurance          int              // Points d'endurance maximum du joueur
	EnduranceRechargeRate int              // Taux de régénération de l'endurance par mise à jour
	Money                 int              // Montant d'argent que le joueur possède
	Speed                 float32          // Vitesse de déplacement du joueur
	Damage                int              // Dégâts infligés par le joueur
	Inventory             []item.Item      // Liste d'objets que le joueur possède dans son inventaire

	IsAlive               bool             // Indicateur de l'état de vie du joueur (vivant ou mort)

	Sprite                rl.Texture2D     // Texture du joueur pour le rendu graphique
}

// AttackOfPlayer : Méthode pour faire attaquer le joueur un monstre
func (p *Player) AttackOfPlayer(m *Monster) {
	m.Health -= p.Damage // Réduit la santé du monstre en fonction des dégâts du joueur
}

// AttackOfPlayerOnMobs : Méthode pour faire attaquer le joueur un mob
func (p *Player) AttackOfPlayerOnMobs(g *Mobs) {
	g.Health -= p.Damage // Réduit la santé du mob en fonction des dégâts du joueur
}

// SpendMoney : Méthode pour dépenser de l'argent chez un vendeur
func (p *Player) SpendMoney(s *Seller) {
	if s.Inventory[0].Name == "Biscuit" {
		p.Money -= 5                         // Déduit le coût du biscuit
		p.Inventory = append(p.Inventory, s.Inventory[0]) // Ajoute le biscuit à l'inventaire
	} else if s.Inventory[1].Name == "Gateau" {
		p.Money -= 15                        // Déduit le coût du gâteau
		p.Inventory = append(p.Inventory, s.Inventory[1]) // Ajoute le gâteau à l'inventaire
	} else if s.Inventory[2].Name == "Bouclier" {
		p.Money -= 25                        // Déduit le coût du bouclier
		p.Inventory = append(p.Inventory, s.Inventory[2]) // Ajoute le bouclier à l'inventaire
	}
}

// UpdateEndurance : Méthode pour mettre à jour l'endurance du joueur
func (p *Player) UpdateEndurance() {
    increment := int(p.EnduranceRechargeRate) // Montant d'endurance à ajouter par mise à jour

    // Augmente l'endurance en fonction du taux de régénération
    if p.Endurance < p.MaxEndurance {
        p.Endurance += increment
        if p.Endurance > p.MaxEndurance {
            p.Endurance = p.MaxEndurance // Assure que l'endurance ne dépasse pas le maximum
        }
    }
}

// UpdateShield : Méthode pour mettre à jour le bouclier du joueur
func (p *Player) UpdateShield() {
    increment := int(p.ShieldRechargeRate) // Montant de bouclier à ajouter par mise à jour

    // Augmente le bouclier en fonction du taux de régénération
    if p.Shield < p.MaxShield {
        p.Shield += increment
        if p.Shield > p.MaxShield {
            p.Shield = p.MaxShield // Assure que le bouclier ne dépasse pas le maximum
        }
    }
}
