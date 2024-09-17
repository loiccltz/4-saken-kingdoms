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

func FightMonster(player entity.Player, monster entity.Monster) {

	for { // Boucle infinie
		// Check si le joueur ou le monstre est vaincu. Si c'est le cas, on sort de la boucle
		if player.Health <= 0 {
			player.IsAlive = false
			break
		}
		if monster.Health <= 0 {
			player.Inventory = append(player.Inventory, monster.Loot[0])
			player.Money += monster.Worth
			break
		}

		player.AttackOfPlayer(&monster)
		monster.AttackOfMonster(&player)
	}
}

func FightMobs(player entity.Player, mobs entity.Mobs) {

	for { // Boucle infinie
		// Check si le joueur ou le monstre est vaincu. Si c'est le cas, on sort de la boucle
		if player.Health <= 0 {
			player.IsAlive = false
			break
		}
		if mobs.Health <= 0 {
			player.Inventory = append(player.Inventory, mobs.Loot[0])
			player.Money += mobs.Worth
			break
		}
		player.AttackOfPlayerOnMobs(&mobs)
		mobs.Attack(&player)
	}
}

func FightShoot(player *entity.Player, shoot *entity.Shoot) {

	for { // Boucle infinie
		// Check si le joueur ou le monstre est vaincu. Si c'est le cas, on sort de la boucle
		if player.Health <= 0 {
			player.IsAlive = false
			break
		}
		shoot.AttackOfShoot(player)
	}
}
