package engine

import (
	"main/src/building"
	"main/src/entity"
	"main/src/item"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 720
)


	func (e *Engine) InitWindow() {
		rl.InitWindow(0, 0, "textures/4SKMENUENTRE-Photoroom.png")
		rl.CloseWindow()
		for !rl.WindowShouldClose() {
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
		}
	}

func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "4saken Kingdom")
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)
	e.InitEntities()
	e.InitTower()
	e.InitItem()
	e.InitMobs()
	e.InitMonsters()
	e.InitShoot()
	e.InitCamera()
	e.InitMusic()
	e.Load()
	e.InitMap("textures/map/tilesets/map.json")

}
func (e *Engine) InitPauseRendering() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "4SKPAUSEMENU.png")
}
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
func (e *Engine) InitEntities() {
	e.Player = entity.Player{

		Position:              rl.Vector2{X: 4430, Y: 6720},
		Health:                100,
		PlayerSrc: rl.NewRectangle(0, 0, 32, 32),
		PlayerDest: rl.NewRectangle(0, 0, -20, -10,),

    
		MaxHealth:             100,
		Shield:                10,
		MaxShield:             100,
		ShieldRechargeRate:    1,
		Endurance:             10,
		MaxEndurance:          100,
		EnduranceRechargeRate: 1,
		Money:                 1000,
		Speed:                 6,
		Damage:                5,
		Inventory:             []item.Item{},
		IsAlive:               true,
		Sprite:                e.Player.Sprite,
	}
	e.Player.Money = 10
	//fmt.Println(e.Player.Position.X)
	//fmt.Println(e.Player.Position.Y)
	e.Seller = entity.Seller{
		Name:      "Robin",
		Position:  rl.Vector2{X: 4400, Y: 6700},
		Money:     500,
		Inventory: []item.Item{},
		IsAlive:   true,
		Sprite:    rl.LoadTexture("textures/towers/TXStruct.png"),
	}
	e.Pnj = entity.Pnj{
		Name:      "Jack",
		Position:  rl.Vector2{X: 4400, Y: 6700},
		IsAlive:   true,
		Sprite:    rl.LoadTexture("textures/towers/TXStruct.png"),
	}

}

func (e *Engine) InitItem() {


	e.Seller.Inventory = append(e.Seller.Inventory, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png"), // Il faut changer la texture pour La Potion
	})
	e.Seller.Inventory = append(e.Seller.Inventory, item.Item{
		Name:         "Epée",
		Price:        15,
		IsConsumable: false,
		IsEquippable: true,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png"), // Il faut changer la texture pour L'Epée
	})
	e.Seller.Inventory = append(e.Seller.Inventory, item.Item{
		Name:         "Bouclier",
		Price:        25,
		IsConsumable: false,
		IsEquippable: true,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png"), // Il faut changer la texture pour Le bouclier
	})

}

func (e *Engine) InitMobs() {
	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob1",
		Position: rl.Vector2{X: 4430, Y: 6880}, // Il faut changer les coordonnées pour Le mob
		Health:   20,
		Damage:   2,
		Loot:     []item.Item{},
		Worth:    25,
		CoolDown: 5 * time.Second,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"), // Il faut changer la texture pour Le mob
	})

	e.Mobs[len(e.Mobs)-1].Loot = append(e.Mobs[len(e.Mobs)-1].Loot, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")}) // Il faut changer la texture pour La potion

	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob2",
		Position: rl.Vector2{X: 4400, Y: 6800}, // Il faut changer les coordonnées pour Le mob
		Health:   20,
		Damage:   2,
		Loot:     []item.Item{},
		Worth:    25,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"), // Il faut changer la texture pour Le mob
	})
	e.Mobs[len(e.Mobs)-1].Loot = append(e.Mobs[len(e.Mobs)-1].Loot, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")}) // Il faut changer la texture pour La potion
}

func (e *Engine) InitMonsters() {
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Loup",
		Position: rl.Vector2{X: 2989, Y: 4850},
		MonsterSrc: rl.NewRectangle(10, 10, 50, 50),
		MonsterDest: rl.NewRectangle(0, 0, 0, 0),

		Health:   20,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    25,
		IsAlive: true,
		Sprite: rl.LoadTexture("textures/entities/boss/animation/loup.png"),
	})
	e.Monsters[len(e.Monsters)-1].Loot = append(e.Monsters[len(e.Monsters)-1].Loot, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")}) // Il faut changer la texture pour La Potion

	e.Monsters = append(e.Monsters, entity.Monster{
		// taille 119x119
		// il faut des tailles paires
		Name:     "Griffon",
		Position: rl.Vector2{X: 3151, Y: 6900},

		MonsterSrc: rl.NewRectangle(10, 10, 50, 50),
		MonsterDest: rl.NewRectangle(0, 0, 0, 0),

		Health:   40,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    50,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/boss/animation/Gryphon.png"),
	})
	e.Monsters[len(e.Monsters)-1].Loot = append(e.Monsters[len(e.Monsters)-1].Loot, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")}) // Il faut changer la texture pour La Potion

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Crabe",
		Position: rl.Vector2{X: 7543, Y: 3870},

		MonsterSrc: rl.NewRectangle(10, 10, 50, 50),
		MonsterDest: rl.NewRectangle(0, 0, 0, 0),

		Health:   60,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    75,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/boss/animation/Crabe"),
	})
	e.Monsters[len(e.Monsters)-1].Loot = append(e.Monsters[len(e.Monsters)-1].Loot, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")}) // Il faut changer la texture pour La Potion

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Dragon",
		Position: rl.Vector2{X: 7634, Y: 7053},

		MonsterSrc: rl.NewRectangle(10, 10, 50, 50),
		MonsterDest: rl.NewRectangle(0, 0, 0, 0),

		Health:   80,
		Damage:   25,
		Loot:     []item.Item{},
		Worth:    100,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/boss/animation/dragon.png"),
	})
	e.Monsters[len(e.Monsters)-1].Loot = append(e.Monsters[len(e.Monsters)-1].Loot, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")}) // Il faut changer la texture pour La Potion

}

func (e *Engine) InitShoot() {
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: 4420, Y: 6880},
		IsShooting: true,
		Damage:     5,
		Sprite:     rl.LoadTexture("textures/tilesets/TX Tileset Stone Ground.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: 4450, Y: 6880},
		IsShooting: true,
		Damage:     5,
		Sprite:     rl.LoadTexture("textures/tilesets/TX Tileset Stone Ground.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: 4400, Y: 6880},
		IsShooting: true,
		Damage:     5,
		Sprite:     rl.LoadTexture("textures/tilesets/TX Tileset Stone Ground.png"),
	})
}
func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D(
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0.0,
		2.0,
	)
}

func (e *Engine) InitMusic() {
	rl.InitAudioDevice()

	e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")

	rl.PlayMusicStream(e.Music)
}
