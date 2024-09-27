package entity

import (
	"fmt" // Importation du package fmt pour les opérations d'entrée/sortie formatées
	"main/src/item" // Importation du package item pour utiliser des éléments dans l'inventaire du vendeur

	rl "github.com/gen2brain/raylib-go/raylib" // Importation de la bibliothèque Raylib pour le rendu graphique
)

// Seller : Structure représentant un vendeur dans le jeu
type Seller struct {
	Name      string      // Nom du vendeur
	Position  rl.Vector2  // Position du vendeur sur l'écran, représentée par un vecteur 2D
	Money     int         // Montant d'argent que le vendeur a
	Inventory []item.Item // Liste des articles disponibles à la vente
	IsAlive   bool        // Indicateur de l'état de vie du vendeur (vivant ou mort)
	Sprite    rl.Texture2D // Texture utilisée pour représenter le vendeur sur l'écran
}

// Buy : Méthode pour acheter un article auprès du vendeur
func (p *Player) Buy(m *Seller) {
	m.Money += 10 // Ajoute 10 à la somme d'argent du vendeur lorsqu'un joueur achète un article
}

// ToString : Méthode pour afficher les informations du vendeur
func (p *Seller) ToString() {
	fmt.Printf(`
	Joueur: Argent: %d,	Inventaire: %+v	\n`, p.Money, p.Inventory)
}
