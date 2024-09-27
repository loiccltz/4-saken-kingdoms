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


		Position: rl.Vector2{X: 4648, Y: 6670},
		Health:   100,

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
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/Potion.png"),
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
		Name:     "Jack",
		Position: rl.Vector2{X: 4628, Y: 6534}, //
		IsAlive:  true,
	})

	e.Pnj = append(e.Pnj, entity.Pnj{
		Name:     "Jacky",
		Position: rl.Vector2{X: 4564, Y: 6544}, //
		IsAlive:  true,
	})

	e.Pnj = append(e.Pnj, entity.Pnj{
		Name:     "Michou",
		Position: rl.Vector2{X: 5844, Y: 5234}, //
		IsAlive:  true,
	})

	e.Pnj = append(e.Pnj, entity.Pnj{
		Name:     "Francis",
		Position: rl.Vector2{X: 5846, Y: 5290}, //
		IsAlive:  true,
	})

	e.Pnj = append(e.Pnj, entity.Pnj{
		Name:     "Garde",
		Position: rl.Vector2{X: 6286, Y: 4264}, //
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
		Regen:        10,
		Sprite:       rl.LoadTexture("textures/items/item.png"),
	})
	e.Seller.Inventory = append(e.Seller.Inventory, item.Item{
		Name:         "Epée",
		Price:        15,
		IsConsumable: true,
		IsEquippable: false,
		Regen:        10,
		Sprite:       rl.LoadTexture("textures/items/item.png"),
	})
	e.Seller.Inventory = append(e.Seller.Inventory, item.Item{
		Name:         "Bouclier",
		Price:        25,
		IsConsumable: false,
		IsEquippable: true,
		Sprite:       rl.LoadTexture("textures/items/item.png"),
	})

}

// InitMobs initialise les mobs avec leurs attributs et loot.
func (e *Engine) InitMobs() {

	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob1",
		Position: rl.Vector2{X: 4064, Y: 5080},
		Health:   20,
		MaxHealth:	20,
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
		Sprite:       rl.LoadTexture("textures/items/item.png")})

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
		Regen:  10,
		Sprite: rl.LoadTexture("textures/items/item.png")})
	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob3",
		Position: rl.Vector2{X: 4450, Y: 6850},
		Health:   3,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    25,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Mobs[len(e.Mobs)-1].Loot = append(e.Mobs[len(e.Mobs)-1].Loot, item.Item{
		Name:         "Biscuit",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Regen:        10,
		Sprite:       rl.LoadTexture("textures/items/item.png")})
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
		Sprite:       rl.LoadTexture("textures/items/item.png")}) 

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
		Sprite:       rl.LoadTexture("textures/items/item.png")}) 

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
		Sprite:       rl.LoadTexture("textures/items/item.png")}) 


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
		Sprite:       rl.LoadTexture("textures/items/item.png")}) 
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
    e.Camera = rl.Camera2D{
        Target: rl.Vector2{X: 0, Y: 0},                 // La cible que la caméra suit (souvent le joueur)
        Offset: rl.Vector2{X: float32(ScreenWidth) / 2, Y: float32(ScreenHeight) / 2},  // Centre de l'écran
        Rotation: 0.0,
        Zoom: 2.0,
    }
}
// InitMusic initialise la musique de fond pour le jeu.
func (e *Engine) InitMusic() {
	rl.InitAudioDevice()
	e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")
}
