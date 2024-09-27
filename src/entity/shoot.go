package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib" // Importation de la bibliothèque Raylib pour le rendu graphique
)

// Shoot : Structure représentant un projectile dans le jeu
type Shoot struct {
	Position   rl.Vector2  // Position du projectile sur l'écran, représentée par un vecteur 2D
	IsShooting bool        // Indicateur si le projectile est en train d'être tiré
	Direction  int         // Direction du tir (0=haut, 1=bas, 2=gauche, 3=droite)
	Damage     int         // Dégâts infligés par le projectile
	Sprite     rl.Texture2D // Texture utilisée pour représenter le projectile sur l'écran
}

// AttackOfShoot : Méthode pour infliger des dégâts au joueur lorsqu'un projectile le touche
func (m *Shoot) AttackOfShoot(p *Player) {
	p.Health -= m.Damage // Réduit la santé du joueur en fonction des dégâts du projectile
}
