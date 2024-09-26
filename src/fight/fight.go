package fight

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)
type Fight2 struct {
	Map     int
	Sprite  rl.Texture2D
}

type fight int

const (
	PLAYER_TURN  fight = iota
	MONSTER_TURN fight = iota
	MOBS_TURN    fight = iota
)

func MonsterVsPlayer(player *entity.Player, monster *entity.Monster) {
	if player.Health <= 0 {
		player.IsAlive = false
		return
	}
	if monster.Health <= 0 {
		player.Inventory = append(player.Inventory, monster.Loot[0])
		player.Money += monster.Worth
		rl.UnloadTexture(monster.Sprite)
	}
	monster.AttackOfMonster(player)
}
func PlayerVsMonster(player *entity.Player, monster *entity.Monster) {
	if player.Health <= 0 {
		player.IsAlive = false
		return
	}
	if monster.Health <= 0 {
		player.Inventory = append(player.Inventory, monster.Loot[0])
		player.Money += monster.Worth
		rl.UnloadTexture(monster.Sprite)

	}
	player.AttackOfPlayer(monster)
}

func MobsVsPlayer(player *entity.Player, mobs *entity.Mobs) {
	if player.Health <= 0 {
		player.IsAlive = false
	
	}
	if mobs.Health <= 0 {
		player.Inventory = append(player.Inventory, mobs.Loot[0])
		player.Money += mobs.Worth
		mobs.IsAlive= false

		rl.UnloadTexture(mobs.Sprite)

	}
	mobs.Attack(player)
}
func PlayerVsMobs(player *entity.Player, mobs *entity.Mobs) {
	if player.Health <= 0 {
		player.IsAlive = false
	}
	if mobs.Health <= 0 {
		player.Inventory = append(player.Inventory, mobs.Loot[0])
		player.Money += mobs.Worth
		mobs.IsAlive= false
		rl.UnloadTexture(mobs.Sprite)

	}
	player.AttackOfPlayerOnMobs(mobs)

}

func ShootVsPlayer(player *entity.Player, shoot *entity.Shoot) {
	if player.Health <= 0 {
		player.IsAlive = false
		return
	}
	shoot.AttackOfShoot(player)
}
