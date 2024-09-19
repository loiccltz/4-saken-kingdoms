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

/*
	func (e *Engine) InitWindow() {
		rl.InitWindow(0, 0, "textures/4SKMENUENTRE-Photoroom.png")
		rl.CloseWindow()
		for !rl.WindowShouldClose() {
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
		}
	}
*/
func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "4saken Kingdom")
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)
	e.InitEntities()
	e.InitBuilding()
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
func (e *Engine) InitBuilding() {
	e.Shop = building.Shop{
		Name:     "Sharp Sword",
		Position: rl.Vector2{X: -100, Y: -20},
		Sprite:   rl.LoadTexture("textures/items/itemschelou.png"),
	}
	e.Tower = append(e.Tower, building.Tower{
		Name:     "Royaume de Ran",
		Position: rl.Vector2{X: 3943, Y: 4890},
		Worth:    25,
		Sprite:   rl.LoadTexture("textures/towers/TXStruct.png"),
	})

	e.Tower = append(e.Tower, building.Tower{
		Name:     "Royaume de Salkin",
		Position: rl.Vector2{X: 3950, Y: 4900},
		Worth:    50,
		Sprite:   rl.LoadTexture("textures/towers/TXStruct.png"),
	})

	e.Tower = append(e.Tower, building.Tower{
		Name:     "Royaume d'Usun",
		Position: rl.Vector2{X: 4000, Y: 4950},
		Worth:    75,
		Sprite:   rl.LoadTexture("textures/towers/TXStruct.png"),
	})

	e.Tower = append(e.Tower, building.Tower{
		Name:     "Royaume de Siroi",
		Position: rl.Vector2{X: 4050, Y: 5000},
		Worth:    100,
		Sprite:   rl.LoadTexture("textures/towers/TXStruct.png"),
	})
}
func (e *Engine) InitEntities() {
	e.Player = entity.Player{
		Position:              rl.Vector2{X: 4430, Y: 6720},
		Health:                100,
		PlayerSrc:             rl.NewRectangle(0, 0, 32, 32),
		PlayerDest:            rl.NewRectangle(0, 0, -20, -10),
		MaxHealth:             100,
		Shield:                10,
		MaxShield:             100,
		ShieldRechargeRate:    1,
		Endurance:             10,
		MaxEndurance:          100,
		EnduranceRechargeRate: 1,
		Money:                 1000,
		Speed:                 9,
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

}

func (e *Engine) InitItem() {

	e.Seller.Inventory = append(e.Seller.Inventory, item.Item{
		Name:         "Biscuit",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png"),
	})
	e.Seller.Inventory = append(e.Seller.Inventory, item.Item{
		Name:         "Gateau",
		Price:        15,
		IsConsumable: true,
		IsEquippable: false,
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

func (e *Engine) InitMobs() {
	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob1",
		Position: rl.Vector2{X: 4430, Y: 6880},
		Health:   20,
		Damage:   2,
		Loot:     []item.Item{},
		Worth:    25,
		CoolDown: 5 * time.Second,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Mobs[len(e.Mobs)-1].Loot = append(e.Mobs[len(e.Mobs)-1].Loot, item.Item{
		Name:         "Biscuit",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")})

	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob2",
		Position: rl.Vector2{X: 4400, Y: 6800},
		Health:   20,
		Damage:   2,
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
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")})
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
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")})
}

func (e *Engine) InitMonsters() {
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "lvl 1",
		Position: rl.Vector2{X: 4450, Y: 6900},
		Health:   20,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    25,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters[len(e.Monsters)-1].Loot = append(e.Monsters[len(e.Monsters)-1].Loot, item.Item{
		Name:         "Biscuit",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "lvl 2",
		Position: rl.Vector2{X: 1000, Y: 620},
		Health:   40,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    50,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters[len(e.Monsters)-1].Loot = append(e.Monsters[len(e.Monsters)-1].Loot, item.Item{
		Name:         "Biscuit",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "lvl 3",
		Position: rl.Vector2{X: -100, Y: 620},
		Health:   60,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    75,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters[len(e.Monsters)-1].Loot = append(e.Monsters[len(e.Monsters)-1].Loot, item.Item{
		Name:         "Biscuit",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "lvl 4",
		Position: rl.Vector2{X: -100, Y: -20},
		Health:   80,
		Damage:   25,
		Loot:     []item.Item{},
		Worth:    100,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters[len(e.Monsters)-1].Loot = append(e.Monsters[len(e.Monsters)-1].Loot, item.Item{
		Name:         "Biscuit",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png")})
}

func (e *Engine) InitShoot() {
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: 0, Y: 0},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/tilesets/TX Tileset Stone Ground.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: 5, Y: 5},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/tilesets/TX Tileset Stone Ground.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: 10, Y: 10},
		IsShooting: true,
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
