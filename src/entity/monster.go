package entity

import (
	"main/src/item"                // Importation du package pour gérer les objets que les monstres peuvent laisser

	rl "github.com/gen2brain/raylib-go/raylib" // Importation de Raylib pour les fonctionnalités graphiques
)

// Monster : Structure représentant un monstre dans le jeu
type Monster struct {
	Name        string          // Nom du monstre
	Position    rl.Vector2      // Position du monstre sur l'écran, représentée par un vecteur 2D
	Health      int             // Points de vie du monstre
	MonsterSrc  rl.Rectangle     // Source de la texture du monstre, utilisée pour l'animation
	MonsterDest rl.Rectangle     // Destination de la texture du monstre pour le rendu
	MonsterFrame int            // Indice du cadre actuel pour l'animation
	FrameCount  int             // Nombre total de cadres dans l'animation
	Damage      int             // Dégâts infligés par le monstre lors d'une attaque
	Loot        []item.Item     // Liste d'objets que le monstre peut laisser après sa défaite
	Worth       int             // Valeur du monstre, potentiellement utilisée pour des mécanismes de récompense (en argent)

	IsAlive     bool            // Indicateur de l'état de vie du monstre (vivant ou mort)

	Sprite      rl.Texture2D    // Texture du monstre pour le rendu graphique
}

// AttackOfMonster : Méthode pour faire attaquer un monstre un joueur
func (m *Monster) AttackOfMonster(p *Player) {
	p.Health -= m.Damage // Réduit la santé du joueur en fonction des dégâts du monstre
}
