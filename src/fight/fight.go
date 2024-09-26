package fight

import (
	"main/src/entity" // Importation du package entity pour utiliser les entités du jeu
	rl "github.com/gen2brain/raylib-go/raylib" // Importation de la bibliothèque Raylib pour le rendu graphique
)

// Fight2 : Structure représentant une instance de combat
type Fight2 struct {
	Map    int           // Indice de la carte sur laquelle le combat a lieu
	Sprite rl.Texture2D  // Texture utilisée pour représenter la scène de combat
}

// fight : Type énuméré pour les états de combat
type fight int

// Constantes représentant les différents tours de combat
const (
	PLAYER_TURN  fight = iota // Tour du joueur
	MONSTER_TURN fight = iota   // Tour du monstre
	MOBS_TURN    fight = iota   // Tour des mobs
)

// MonsterVsPlayer : Gère le combat entre un joueur et un monstre
func MonsterVsPlayer(player *entity.Player, monster *entity.Monster) {
	if player.Health <= 0 { // Vérifie si le joueur est mort
		player.IsAlive = false // Marque le joueur comme mort
		return
	}
	if monster.Health <= 0 { // Vérifie si le monstre est mort
		player.Inventory = append(player.Inventory, monster.Loot[0]) // Ajoute le butin du monstre à l'inventaire du joueur
		monster.IsAlive = false // Marque le monstre comme mort
		player.Money += monster.Worth // Ajoute la valeur du monstre à l'argent du joueur
		rl.UnloadTexture(monster.Sprite) // Décharge la texture du monstre pour libérer de la mémoire
	}
	monster.AttackOfMonster(player) // Le monstre attaque le joueur
}

// PlayerVsMonster : Gère le combat entre un joueur et un monstre, version inversée
func PlayerVsMonster(player *entity.Player, monster *entity.Monster) {
	if player.Health <= 0 { // Vérifie si le joueur est mort
		player.IsAlive = false // Marque le joueur comme mort
		return
	}
	if monster.Health <= 0 { // Vérifie si le monstre est mort
		player.Inventory = append(player.Inventory, monster.Loot[0]) // Ajoute le butin du monstre à l'inventaire du joueur
		monster.IsAlive = false // Marque le monstre comme mort
		player.Money += monster.Worth // Ajoute la valeur du monstre à l'argent du joueur
		rl.UnloadTexture(monster.Sprite) // Décharge la texture du monstre pour libérer de la mémoire
	}
	player.AttackOfPlayer(monster) // Le joueur attaque le monstre
}

// MobsVsPlayer : Gère le combat entre un joueur et un mob
func MobsVsPlayer(player *entity.Player, mobs *entity.Mobs) {
	if player.Health <= 0 { // Vérifie si le joueur est mort
		player.IsAlive = false // Marque le joueur comme mort
	}
	if mobs.Health <= 0 { // Vérifie si le mob est mort
		player.Inventory = append(player.Inventory, mobs.Loot[0]) // Ajoute le butin du mob à l'inventaire du joueur
		player.Money += mobs.Worth // Ajoute la valeur du mob à l'argent du joueur
		mobs.IsAlive = false // Marque le mob comme mort
		rl.UnloadTexture(mobs.Sprite) // Décharge la texture du mob pour libérer de la mémoire
	}
	mobs.Attack(player) // Le mob attaque le joueur
}

// PlayerVsMobs : Gère le combat entre un joueur et un mob, version inversée
func PlayerVsMobs(player *entity.Player, mobs *entity.Mobs) {
	if player.Health <= 0 { // Vérifie si le joueur est mort
		player.IsAlive = false // Marque le joueur comme mort
	}
	if mobs.Health <= 0 { // Vérifie si le mob est mort
		player.Inventory = append(player.Inventory, mobs.Loot[0]) // Ajoute le butin du mob à l'inventaire du joueur
		player.Money += mobs.Worth // Ajoute la valeur du mob à l'argent du joueur
		mobs.IsAlive = false // Marque le mob comme mort
		rl.UnloadTexture(mobs.Sprite) // Décharge la texture du mob pour libérer de la mémoire
	}
	player.AttackOfPlayerOnMobs(mobs) // Le joueur attaque le mob
}

// ShootVsPlayer : Gère le combat entre un joueur et un projectile
func ShootVsPlayer(player *entity.Player, shoot *entity.Shoot) {
	if player.Health <= 0 { // Vérifie si le joueur est mort
		player.IsAlive = false // Marque le joueur comme mort
		return
	}
	shoot.AttackOfShoot(player) // Le projectile attaque le joueur
}
