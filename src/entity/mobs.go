package entity

import (
	"main/src/item"          // Importation du package pour gérer les objets que les mobs peuvent laisser
	"time"                   // Importation du package time pour gérer les durées et les horodatages

	rl "github.com/gen2brain/raylib-go/raylib" // Importation de Raylib pour les fonctionnalités graphiques
)

// Mobs : Structure représentant un mob dans le jeu
type Mobs struct {
	Name           string          // Nom du mob
	Position       rl.Vector2      // Position du mob sur l'écran, représentée par un vecteur 2D
	Health         int             // Points de vie du mob
 	 MaxHealth	   int
	Damage         int             // Dégâts infligés par le mob lors d'une attaque
	Loot           []item.Item     // Liste d'objets que le mob peut laisser après sa défaite
	Worth          int             // Valeur du mob, potentiellement utilisée pour des mécanismes de récompense
	CoolDown       time.Duration    // Temps d'attente entre les attaques, pour gérer la fréquence d'attaque
	IsAlive        bool            // Indicateur de l'état de vie du mob (vivant ou mort)
	Sprite         rl.Texture2D    // Texture du mob pour le rendu graphique
	LastAttackTime time.Time        // Horodatage du dernier moment où le mob a attaqué

}

// Attack : Méthode pour faire attaquer un mob un joueur
func (g *Mobs) Attack(p *Player) {
	p.Health -= g.Damage // Réduit la santé du joueur en fonction des dégâts du mob
}
