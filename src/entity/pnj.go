package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib" // Importation de la bibliothèque Raylib pour le rendu graphique
)

// Pnj : Structure représentant un personnage non joueur dans le jeu
type Pnj struct {
	Name     string      // Nom du PNJ
	Position rl.Vector2  // Position du PNJ sur l'écran, représentée par un vecteur 2D
	IsAlive  bool        // Indicateur de l'état de vie du PNJ (vivant ou mort)
}
