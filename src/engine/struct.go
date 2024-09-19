package engine

import (
	"main/src/building"
	"main/src/entity"
	"main/src/fight"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type menu int

const (
	HOME     menu = iota
	SETTINGS menu = iota
	PLAY     menu = iota
	FIGHT    menu = iota
)

type engine int

const (
	INGAME     engine = iota
	INVENTORY  engine = iota
	MENUSELLER engine = iota
	PAUSE      engine = iota
	INFIGHT    engine = iota
	GAMEOVER   engine = iota
)

type Engine struct {
	Player   entity.Player
	Monsters []entity.Monster
	Seller   entity.Seller
	Mobs     []entity.Mobs
	Item     []item.Item
	Tower    []building.Tower
	Shop     building.Shop
	Shoot    []entity.Shoot
	Maps     []fight.Fight2
	selectedIndex int 

	Objects  []Object

	Music       rl.Music
	MusicVolume float32

	Sprites map[string]rl.Texture2D
	Animations Animations
	Camera rl.Camera2D

	MapJSON MapJSON

	IsRunning     bool
	StateMenu     menu
	StateEngine   engine
	StateFight    menu
	InventoryMenu engine
	SellerMenu    engine
}
type Fight struct {
	CurrentMonster      entity.Monster
	CurrentMonsterIndex int
	CurrentMobs         entity.Mobs
	CurrentMobsIndex    int
}

const(
	ChaseDistance = 100
)
