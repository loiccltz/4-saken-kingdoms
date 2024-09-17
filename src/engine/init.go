package engine

import (
	"main/src/building"
	"main/src/entity"
	"main/src/item"
	"main/src/fight"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 2500
	ScreenHeight = 1500
)
const (
	ScreenWidthF  = 600
	ScreenHeightF = 1400
)

const (
	ScreenWidthS  = 800
	ScreenHeightS = 800
)

func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "4saken Kingdom")
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)
	e.InitEntities()
	e.InitCamera()
	e.InitMusic()
	e.InitMobs()
	e.InitMap("textures/map/tilesets/map.json")

}
func (e *Engine) Init_lvl1() {
	rl.CloseWindow()
	rl.InitWindow(ScreenWidthF, ScreenHeightF, "Tower")
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)
	e.InitEntities()
	e.InitCamera()
	e.InitMusicF()
	e.InitShoot()
	e.InitMonsters1()
	e.InitMap("textures/map/tilesets/map.json")
}
func (e *Engine) Init_lvl2() {
	rl.CloseWindow()
	rl.InitWindow(ScreenWidthF, ScreenHeightF, "Tower")
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)
	e.InitEntities()
	e.InitCamera()
	e.InitMusicF()
	e.InitShoot()
	e.InitMonsters2()
	e.InitMap("textures/map/tilesets/map.json")
}
func (e *Engine) Init_lvl3() {
	rl.CloseWindow()
	rl.InitWindow(ScreenWidthF, ScreenHeightF, "Tower")
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)
	e.InitEntities()
	e.InitCamera()
	e.InitMusicF()
	e.InitShoot()
	e.InitMonsters3()
	e.InitMap("textures/map/tilesets/map.json")
}
func (e *Engine) Init_lvl4() {
	rl.CloseWindow()
	rl.InitWindow(ScreenWidthF, ScreenHeightF, "Tower")
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)
	e.InitEntities()
	e.InitCamera()
	e.InitMusicF()
	e.InitShoot()
	e.InitMonsters4()
	e.InitMap("textures/map/tilesets/map.json")
}
func (e *Engine) Init_Shop() {
	rl.CloseWindow()
	rl.InitWindow(ScreenWidthF, ScreenHeightF, "Tower")
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)
	e.InitEntitiesShop()
	e.InitCamera()
	e.InitMusicF()
	e.InitMap("textures/map/tilesets/map.json")
}

func (e *Engine) InitEntities() {
	e.Player = entity.Player{
		Position:  rl.Vector2{X: 2000, Y: 2000},
		Health:    100,
		//MaxHealth:	100,
		//Shield: 10,
		//MaxShield: 100,
		//ShieldRechargeRate: 1,
		//Endurance:	10,
		//MaxEndurance:	100,
		//EnduranceRechargeRate:	1,
		Money:     1000,
		Speed:     6,
		Damage:    10,
		Inventory: []item.Item{},

		IsAlive: true,

		Sprite: e.Player.Sprite,
	}
	

	e.Shop = building.Shop{
		Name:     "Sharp Sword",
		Position: rl.Vector2{X: -100, Y: -20},
		Sprite:   rl.LoadTexture("textures/items/itemschelou.png"),
	}

	e.Tower = append(e.Tower, building.Tower{
		Name:     "Royaume de Ran",
		Position: rl.Vector2{X: -25, Y: -20},
		Worth:    25,
		Sprite:   rl.LoadTexture("textures/towers/TXStruct.png"),
	})

	e.Tower = append(e.Tower, building.Tower{
		Name:     "Royaume de Salkin",
		Position: rl.Vector2{X: -50, Y: -20},
		Worth:    50,
		Sprite:   rl.LoadTexture("textures/towers/TXStruct.png"),
	})

	e.Tower = append(e.Tower, building.Tower{
		Name:     "Royaume d'Usun",
		Position: rl.Vector2{X: -75, Y: -20},
		Worth:    75,
		Sprite:   rl.LoadTexture("textures/towers/TXStruct.png"),
	})

	e.Tower = append(e.Tower, building.Tower{
		Name:     "Royaume de Siroi",
		Position: rl.Vector2{X: -100, Y: -20},
		Worth:    100,
		Sprite:   rl.LoadTexture("textures/towers/TXStruct.png"),
	})

	e.Player.Money = 10
}

func (e *Engine) InitEntitiesShop() {
	e.Seller = entity.Seller{
		Position:  rl.Vector2{X: 300, Y: 300},
		Money:     500,
		Inventory: []item.Item{},
		IsAlive:   true,
		Sprite:    rl.LoadTexture("textures/towers/TXStruct.png"),
	}
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
func (e *Engine) InitMaps(){
	e.Maps = append(e.Maps, fight.Fight2{
		Map:     1,
		Sprite:  rl.LoadTexture("textures/items/itemschelou.png"),
	})
	e.Maps = append(e.Maps, fight.Fight2{
		Map:     2,
		Sprite:  rl.LoadTexture("textures/items/itemschelou.png"),
	})
	e.Maps = append(e.Maps, fight.Fight2{
		Map:     4,
		Sprite:  rl.LoadTexture("textures/items/itemschelou.png"),
	})
	e.Maps = append(e.Maps, fight.Fight2{
		Map:     4,
		Sprite:  rl.LoadTexture("textures/items/itemschelou.png"),
	})
}

func (e *Engine) InitMobs() {
	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob1",
		Position: rl.Vector2{X: 1000, Y: -20},
		Health:   20,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    25,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob2",
		Position: rl.Vector2{X: 900, Y: -20},
		Health:   20,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    25,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob3",
		Position: rl.Vector2{X: 800, Y: -20},
		Health:   20,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    25,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob4",
		Position: rl.Vector2{X: 700, Y: -20},
		Health:   20,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    25,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob5",
		Position: rl.Vector2{X: 600, Y: -20},
		Health:   20,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    25,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob6",
		Position: rl.Vector2{X: 500, Y: -20},
		Health:   20,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    25,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob7",
		Position: rl.Vector2{X: 400, Y: -20},
		Health:   20,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    25,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Mobs = append(e.Mobs, entity.Mobs{
		Name:     "mob8",
		Position: rl.Vector2{X: 300, Y: -20},
		Health:   20,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    25,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
}

func (e *Engine) InitMonsters1() {
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "lvl 1",
		Position: rl.Vector2{X: 1000, Y: -20},
		Health:   20,
		Damage:   1,
		Loot:     []item.Item{},
		Worth:    25,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
}

func (e *Engine) InitMonsters2() {
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
}

func (e *Engine) InitMonsters3() {
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
}

func (e *Engine) InitMonsters4() {
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
}
func (e *Engine) InitShoot() {
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
	})
	e.Shoot = append(e.Shoot, entity.Shoot{
		Position:   rl.Vector2{X: -75, Y: -20},
		IsShooting: true,
		Sprite:     rl.LoadTexture("textures/towers/TXStruct.png"),
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

func (e *Engine) InitMusicF() {
	rl.InitAudioDevice()

	e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")

	rl.PlayMusicStream(e.Music)
}
