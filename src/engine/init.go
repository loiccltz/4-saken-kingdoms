package engine

import (
	"main/src/building"
	"main/src/entity"
	"main/src/item"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Variables globales pour la largeur et la hauteur de l'écran
var (
	ScreenWidth  = rl.GetScreenWidth()
	ScreenHeight = rl.GetScreenHeight()
)

// InitWindow initialise une fenêtre temporaire puis la ferme.
func (e *Engine) InitWindow() {
	rl.InitWindow(0, 0, "textures/4SKMENUENTRE-Photoroom.png")
	rl.CloseWindow()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
	}
}

func (e *Engine) Init() {
	rl.InitWindow(int32(ScreenWidth), int32(ScreenHeight), "4saken Kingdom")
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)   // Initialise la map des textures
	e.InitEntities()                            // Initialisation des entités du jeu
	e.InitTower()                               // Initialisation des tours dans le jeu
	e.InitItem()                                // Initialisation des items
	e.InitMobs()                                // Initialisation des mobs
	e.InitMonsters()                            // Initialisation des monstres
	e.InitShoot()                               // Initialisation des tirs
	e.InitCamera()                              // Initialisation de la caméra
	e.InitMusic()                               // Initialisation de la musique de fond
	e.Load()                                    // Chargement des ressources
	e.InitMap("textures/map/tilesets/map.json") // Chargement de la carte

}

// InitPauseRendering initialise la fenêtre pour le menu pause.
func (e *Engine) InitPauseRendering() {
	rl.InitWindow(int32(ScreenWidth), int32(ScreenHeight), "4SKPAUSEMENU.png")
}

// InitTower initialise les tours avec leurs positions, noms, valeurs, et textures.
func (e *Engine) InitTower() {

	e.Tower = append(e.Tower, building.Tower{
		Name:     "Royaume de Ran",
		Position: rl.Vector2{X: 3943, Y: 4890},
		Worth:    25,
		Sprite:   rl.LoadTexture("textures/fefolet.png"),
	})

	e.Tower = append(e.Tower, building.Tower{
		Name:     "Royaume de Salkin",
		Position: rl.Vector2{X: 3950, Y: 4900},
		Worth:    50,
		Sprite:   rl.LoadTexture("textures/fefolet.png"),
	})

	e.Tower = append(e.Tower, building.Tower{
		Name:     "Royaume d'Usun",
		Position: rl.Vector2{X: 4965, Y: 6311},
		Worth:    75,
		Sprite:   rl.LoadTexture("textures/fefolet.png"),
	})

	e.Tower = append(e.Tower, building.Tower{
		Name:     "Royaume de Siroi",
		Position: rl.Vector2{X: 4050, Y: 5000},
		Worth:    100,
		Sprite:   rl.LoadTexture("textures/fefolet.png"),
	})
}

// InitEntities initialise les entités comme le joueur, le vendeur et les PNJs.
func (e *Engine) InitEntities() {

	e.Player = entity.Player{
		Position:   rl.Vector2{X: 4430, Y: 6720},
		Health:     90,
		PlayerSrc:  rl.NewRectangle(0, 0, 32, 32),
		PlayerDest: rl.NewRectangle(0, 0, -20, -10),
		MaxHealth:  100,
		Shield:     10,
		MaxShield:  100,
		// D'autres attributs du joueur
		Money:   1000,
		Speed:   10,
		Damage:  5,
		IsAlive: true,
		Sprite:  e.Player.Sprite,
	}

	// Ajout d'un item dans l'inventaire du joueur
	e.Player.Inventory = append(e.Player.Inventory, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		Regen:        10,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png"),
	})

	e.Player.Money = 10

	// Initialisation du vendeur
	e.Seller = entity.Seller{
		Name:      "Robin",
		Position:  rl.Vector2{X: 5250, Y: 5300},
		Money:     500,
		Inventory: []item.Item{},
		IsAlive:   true,
		Sprite:    rl.LoadTexture("textures/towers/TXStruct.png"),
	}

	// Ajout de plusieurs PNJs
	e.Pnj = append(e.Pnj, entity.Pnj{
		Name:     "Marie",
		Position: rl.Vector2{X: 46040, Y: 6560},
		IsAlive:  true,
	})

	e.Pnj = append(e.Pnj, entity.Pnj{
		Name:     "John1",
		Position: rl.Vector2{X: 4390, Y: 5350},
		IsAlive:  true,
	})

	e.Pnj = append(e.Pnj, entity.Pnj{
		Name:     "Jean2",
		Position: rl.Vector2{X: 5790, Y: 5210},
		IsAlive:  true,
	})

	e.Pnj = append(e.Pnj, entity.Pnj{
		Name:     "Jill3",
		Position: rl.Vector2{X: 4600, Y: 6500},
		IsAlive:  true,
	})

	e.Pnj = append(e.Pnj, entity.Pnj{
		Name:     "Judi4",
		Position: rl.Vector2{X: 4700, Y: 6500},
		IsAlive:  true,
	})
}

// InitItem initialise les items dans l'inventaire du vendeur.
func (e *Engine) InitItem() {

	e.Seller.Inventory = append(e.Seller.Inventory, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png"),
	})
	e.Seller.Inventory = append(e.Seller.Inventory, item.Item{
		Name:         "Epée",
		Price:        15,
		IsConsumable: false,
		IsEquippable: true,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png"),
	})
	e.Seller.Inventory = append(e.Seller.Inventory, item.Item{
		Name:         "Bouclier",
		Price:        25,
		IsConsumable: false,
		IsEquippable: true,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png"),
	})

}

// InitMobs initialise les mobs avec leurs attributs et loot.
func (e *Engine) InitMobs() {

	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob1",
		Position: rl.Vector2{X: 4200, Y: 6250},
		Health:   10,
		Damage:   2,
		Loot:     []item.Item{},
		Worth:    25,
		CoolDown: 5 * time.Second,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	// Ajout d'un loot au mob1
	e.Mobs[len(e.Mobs)-1].Loot = append(e.Mobs[len(e.Mobs)-1].Loot, item.Item{
		Name:         "Potion",
		Price:        5,
		Regen:        10,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")})

	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob2",
		Position: rl.Vector2{X: 4712, Y: 5750},
		Damage:   2,
		Loot:     []item.Item{},
		Worth:    25,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Mobs[len(e.Mobs)-1].Loot = append(e.Mobs[len(e.Mobs)-1].Loot, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")})
}

// InitMonsters initialise les monstres avec leurs attributs et loot.
func (e *Engine) InitMonsters() {
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:        "Loup",
		Position:    rl.Vector2{X: 2989, Y: 4850},
		MonsterSrc:  rl.NewRectangle(10, 10, 50, 50),
		MonsterDest: rl.NewRectangle(0, 0, 0, 0),

		Health:  20,
		Damage:  1,
		Loot:    []item.Item{},
		Worth:   25,
		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/boss/animation/loup.png"),
	})

	// Ajout d'un loot au monstre "Loup"
	e.Monsters[len(e.Monsters)-1].Loot = append(e.Monsters[len(e.Monsters)-1].Loot, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")})

	e.Monsters = append(e.Monsters, entity.Monster{
		// taille 119x119
		// il faut des tailles paires
		Name:        "Griffon",
		Position:    rl.Vector2{X: 3151, Y: 6900},
		MonsterSrc:  rl.NewRectangle(10, 10, 50, 50),
		MonsterDest: rl.NewRectangle(0, 0, 0, 0),
		Health:      40,
		Damage:      1,
		Loot:        []item.Item{},
		Worth:       50,
		IsAlive:     true,
		Sprite:      rl.LoadTexture("textures/entities/boss/animation/Gryphon.png"),
	})

	e.Monsters[len(e.Monsters)-1].Loot = append(e.Monsters[len(e.Monsters)-1].Loot, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:        "Crabe",
		Position:    rl.Vector2{X: 7543, Y: 3870},
		MonsterSrc:  rl.NewRectangle(10, 10, 50, 50),
		MonsterDest: rl.NewRectangle(0, 0, 0, 0),
		Health:      60,
		Damage:      1,
		Loot:        []item.Item{},
		Worth:       75,
		IsAlive:     true,
		Sprite:      rl.LoadTexture("textures/entities/boss/animation/Crabe"),
	})

	e.Monsters[len(e.Monsters)-1].Loot = append(e.Monsters[len(e.Monsters)-1].Loot, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:        "Dragon",
		Position:    rl.Vector2{X: 7634, Y: 7053},
		MonsterSrc:  rl.NewRectangle(10, 10, 50, 50),
		MonsterDest: rl.NewRectangle(0, 0, 0, 0),
		Health:      80,
		Damage:      25,
		Loot:        []item.Item{},
		Worth:       100,
		IsAlive:     true,
		Sprite:      rl.LoadTexture("textures/entities/boss/animation/dragon.png"),
	})

	e.Monsters[len(e.Monsters)-1].Loot = append(e.Monsters[len(e.Monsters)-1].Loot, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")})
}

// InitShoot initialise les tirs dans le jeu avec leur direction et dégâts.
func (e *Engine) InitShoot() {
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: 3100, Y: 4775},
		IsShooting: true,
		Direction:  3,
		Damage:     1,
		Sprite:     rl.LoadTexture("textures/fefolet.png"),
	})

	// Ajout d'autres tirs
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: 3100, Y: 4850},
		IsShooting: true,
		Direction:  2,
		Damage:     1,
		Sprite:     rl.LoadTexture("textures/fefolet.png"),
	})

	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: 3100, Y: 4950},
		IsShooting: true,
		Direction:  1,
		Damage:     1,
		Sprite:     rl.LoadTexture("textures/fefolet.png"),
	})

	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: 3150, Y: 4775},
		IsShooting: true,
		Direction:  4,
		Damage:     1,
		Sprite:     rl.LoadTexture("textures/fefolet.png"),
	})

	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: 3050, Y: 4850},
		IsShooting: true,
		Direction:  2,
		Damage:     1,
		Sprite:     rl.LoadTexture("textures/fefolet.png"),
	})

	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: 3000, Y: 4950},
		IsShooting: true,
		Direction:  1,
		Damage:     1,
		Sprite:     rl.LoadTexture("textures/fefolet.png"),
	})
}

// InitCamera initialise la caméra du jeu avec des paramètres de base.
func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D(
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0.0,
		2.0,
	)
}

// InitMusic initialise la musique de fond pour le jeu.
func (e *Engine) InitMusic() {
	rl.InitAudioDevice()
	e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")
}
