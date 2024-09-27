package engine

import (
	"main/src/building" // Importation du package pour gérer les bâtiments
	"main/src/entity"    // Importation du package pour gérer les entités du jeu
	"main/src/fight"     // Importation du package pour gérer les combats
	"main/src/item"      // Importation du package pour gérer les objets

	rl "github.com/gen2brain/raylib-go/raylib" // Importation de Raylib pour les fonctionnalités graphiques
)

// menu : Type énuméré représentant les différents états du menu
type menu int

// Constantes définissant les états du menu
const (
	HOME     menu = iota // État du menu d'accueil
	SETTINGS menu = iota // État du menu des paramètres
	PLAY     menu = iota // État du menu de jeu
)

// engine : Type énuméré représentant les différents états du moteur de jeu
type engine int

// Constantes définissant les états du moteur de jeu
const (
	INGAME     engine = iota // État de jeu actif
	INVENTORY  engine = iota // État de l'inventaire
	MENUSELLER engine = iota // État du menu du vendeur
	PAUSE      engine = iota // État de pause
	GAMEOVER   engine = iota // État de fin de jeu
)

// Engine : Structure principale du moteur de jeu
type Engine struct {
	Player   entity.Player        // Instance du joueur
	Monsters []entity.Monster     // Liste des monstres
	Seller   entity.Seller        // Instance du vendeur
	Mobs     []entity.Mobs        // Liste des mobs
	Item     []item.Item          // Liste des objets
	Tower    []building.Tower     // Liste des tours
	Pnj      []entity.Pnj         // Liste des PNJ
	Shoot    []entity.Shoot       // Liste des projectiles
	Maps     []fight.Fight2       // Liste des cartes de combat
	selectedIndex int             // Index de l'élément sélectionné dans un menu

	Objects    []Object           // Liste des objets dans le jeu
	BossFight  []Ranger           // Liste des boss pour le combat

	Music       rl.Music          // Musique de fond
	MusicVolume float32            // Volume de la musique

	Sprites    map[string]rl.Texture2D // Dictionnaire des sprites avec leurs textures
	Animations Animations          // Animations du jeu
	Camera     rl.Camera2D         // Caméra 2D pour le rendu

	MapJSON MapJSON                // Données JSON pour la carte

	IsRunning   bool              // Indicateur si le jeu est en cours d'exécution
	StateMenu   menu              // État actuel du menu
	StateEngine engine             // État actuel du moteur de jeu

	InventoryMenu engine           // État de l'inventaire
	SellerMenu    engine           // État du menu du vendeur

	ScreenWidth  int32             // Largeur de l'écran
    ScreenHeight int32             // Hauteur de l'écran
}

// Fight : Structure pour gérer les combats
type Fight struct {
	CurrentMonster      entity.Monster // Monstre actuel en combat
	CurrentMonsterIndex int             // Index du monstre actuel
	CurrentMobs         entity.Mobs     // Mob actuel en combat
	CurrentMobsIndex    int             // Index du mob actuel
}

// Constantes pour les distances de chasse
const (
	ChaseDistance = 100 // Distance à laquelle un monstre commence à poursuivre le joueur
)
