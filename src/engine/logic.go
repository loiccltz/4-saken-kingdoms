package engine

import (
	"fmt"
	"main/src/building"
	"main/src/entity"
	"main/src/fight"
	"math/rand"

	//"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) HomeLogic() {

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	//Menus
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
		rl.StopMusicStream(e.Music)

	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
}

func (e *Engine) SettingsLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyB) {
		e.StateMenu = HOME
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) InGameLogic() {
	// Mouvement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		e.Player.Position.Y -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		e.Player.Position.Y += e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		e.Player.Position.X -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		e.Player.Position.X += e.Player.Speed
	}

	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.StateEngine = PAUSE
	}
	if rl.IsKeyPressed(rl.KeyI) {
		e.StateEngine = INVENTORY
	}
	if rl.IsKeyPressed(rl.KeyR) {
		e.StateEngine = MENUSELLER
	}

	e.CheckCollisions()

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-07-Simon_s-In-There-Somewhere.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	if e.Player.Health <= 0 {
		e.Player.IsAlive = false
		e.Player.Money /= 2
		e.StateEngine = GAMEOVER
	}
	for i := range e.Mobs {
		if e.Mobs[i].Health <= 0 {
			e.Mobs[i].IsAlive = false
		}

		e.FightCollisions()
		posX := e.Player.Position.X
		posY := e.Player.Position.Y
		if posX == 356 && posY == 200 {
			for e.Monsters[0].Health > 0 {
				if rand.Intn(100) < 10 {
					e.CreateShoot()
				}
				e.MoveShoot()
				e.UpdateShoot()
				e.ShootCollision()
			}
			if e.Monsters[0].Health <= 0 {
				posX =3
				posY =3
				break
			}
		}
		if posX == 356 && posY == 200 {
			for e.Monsters[1].Health > 0 {
				if rand.Intn(100) < 10 {
					e.CreateShoot()
				}
				e.MoveShoot()
				e.UpdateShoot()
				e.ShootCollision()
			}
			if e.Monsters[1].Health <= 0 {
				posX =3
				posY =3
				break
			}
		}
		if posX == 356 && posY == 200 {
			for e.Monsters[2].Health > 0 {
				if rand.Intn(100) < 10 {
					e.CreateShoot()
				}
				e.MoveShoot()
				e.UpdateShoot()
				e.ShootCollision()
			}
			if e.Monsters[2].Health <= 0 {
				posX =3
				posY =3
				break
			}
		}
		if posX == 356 && posY == 200 {
			for e.Monsters[3].Health > 0 {
				if rand.Intn(100) < 10 {
					e.CreateShoot()
				}
				e.MoveShoot()
				e.UpdateShoot()
				e.ShootCollision()
			}
			if e.Monsters[3].Health <= 0 {
				posX =3
				posY =3
				break
			}
		}
	}
}
func (e *Engine) GameOverLogic() {
	if e.Player.Health <= 0 {
		if !rl.IsMusicStreamPlaying(e.Music) {
			e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")
			rl.PlayMusicStream(e.Music)
		}
		rl.UpdateMusicStream(e.Music)

		//Menus
		if rl.IsKeyPressed(rl.KeyEnter) {
			e.Player.IsAlive = true
			e.Player.Health = e.Player.MaxHealth
			e.StateMenu = PLAY
			e.StateEngine = INGAME
			rl.StopMusicStream(e.Music)

		}
		if rl.IsKeyPressed(rl.KeyEscape) {
			e.IsRunning = false
		}
	}
}

func (e *Engine) InventoryLogic() {
	if rl.IsKeyPressed(rl.KeyI) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyI) {
		e.InventoryMenu = INVENTORY
	}
}
func (e *Engine) SellerLogic() {
	if rl.IsKeyPressed(rl.KeyR) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyR) {
		e.SellerMenu = MENUSELLER
	}
}

func (e *Engine) CheckCollisionsWithObjects() bool {
	playerRect := rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 40, 40)
	// * 2 - 16
	for _, obj := range e.Objects {
		objectRect := rl.NewRectangle(obj.X*2-16, obj.Y*2-16, obj.Width, obj.Height)
		if rl.CheckCollisionRecs(playerRect, objectRect) {
			// fmt.Print("coli")
			// fmt.Println(objectRect)
			return true
		}
	}
	return false
}

func (e *Engine) CheckCollisions() {
	// fmt.Println(e.Player.Position.X)
	// fmt.Println(e.Player.Position.Y)
	e.MobsCollisions()
	//e.MonsterCollisions()
	e.ShootCollision()
	e.UpdateMobs()
	e.TowerCollisions()
	e.SellerCollisions()
	e.CheckCollisionsWithObjects()
	e.BlockCollisions()
}
func (e *Engine) BlockCollisions() {
	if e.CheckCollisionsWithObjects() {
		if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
			e.Player.Position.Y += e.Player.Speed
		}
		if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
			e.Player.Position.Y -= e.Player.Speed
		}
		if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
			e.Player.Position.X += e.Player.Speed
		}
		if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
			e.Player.Position.X -= e.Player.Speed
		}
	}
}
func (e *Engine) FightCollisions() {
	e.UpdateShoot()
	e.MoveShoot()
	e.UpdateShoot()
	//e.SpellCollision()
}

// GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG
func (e *Engine) SellerCollisions() {
	if e.Seller.Position.X > e.Player.Position.X-20 &&
		e.Seller.Position.X < e.Player.Position.X+20 &&
		e.Seller.Position.Y > e.Player.Position.Y-20 &&
		e.Seller.Position.Y < e.Player.Position.Y+20 {
		if e.Seller.Name == "Robin" {
			e.NormalExplanationShop(e.Seller, "We have all do you want to rescue the princess, press R to enter")
		}
	}
}
func (e *Engine) PnjCollisions() {
	if e.Pnj.Position.X > e.Player.Position.X-20 &&
		e.Pnj.Position.X < e.Player.Position.X+20 &&
		e.Pnj.Position.Y > e.Player.Position.Y-20 &&
		e.Pnj.Position.Y < e.Player.Position.Y+20 {
		if e.Pnj.Name == "Jack" {
			e.NormalExplanationPnj(e.Pnj, "We have all do you want to rescue the princess, press R to enter")
		}
	}
}
func (e *Engine) TowerCollisions() {
	for _, tower := range e.Tower {
		if tower.Position.X > e.Player.Position.X-20 &&
			tower.Position.X < e.Player.Position.X+20 &&
			tower.Position.Y > e.Player.Position.Y-20 &&
			tower.Position.Y < e.Player.Position.Y+20 {
			if tower.Name == "Royaume de Ran" {
				e.NormalExplanation(tower, "To save Princess Tom press J")
				if rl.IsKeyPressed(rl.KeyJ) {
					e.Player.Position.X = 100
					e.Player.Position.Y = 700
				}
			} else if tower.Name == "Royaume de Salkin" {
				e.NormalExplanation(tower, "To save Princess Arnaud press K")
				if rl.IsKeyPressed(rl.KeyK) {
					e.Player.Position.X = 100
					e.Player.Position.Y = 700
				}
			} else if tower.Name == "Royaume d'Usun" {
				e.NormalExplanation(tower, "To save Princess Yann press L")
				if rl.IsKeyPressed(rl.KeyL) {
					e.Player.Position.X = 100
					e.Player.Position.Y = 700
				}
			} else if tower.Name == "Royaume de Siroi" {
				e.NormalExplanation(tower, "To save Princess Léo press M")
				if rl.IsKeyPressed(rl.KeyM) {
					e.Player.Position.X = 100
					e.Player.Position.Y = 700
				}
			}
		}
	}
}

func (e *Engine) MobsCollisions() {
	for i := range e.Mobs {
		if e.Mobs[i].IsAlive {
			if e.Mobs[i].Position.X > e.Player.Position.X-20 &&
				e.Mobs[i].Position.X < e.Player.Position.X+20 &&
				e.Mobs[i].Position.Y > e.Player.Position.Y-20 &&
				e.Mobs[i].Position.Y < e.Player.Position.Y+20 {
				fmt.Println(e.Mobs[i].Health)
				if e.Mobs[i].IsAlive == true && e.Player.IsAlive {
					e.NormalTalkMobs(e.Mobs[i], "Bonjour")
					//if time.Since(e.Mobs[i].LastAttackTime) < e.Mobs[i].CoolDown {
					fight.MobsVsPlayer(&e.Player, &e.Mobs[i])
					//e.Mobs[i].LastAttackTime = time.Now()
				}
				if rl.IsKeyPressed(rl.KeyE) {
					fight.PlayerVsMobs(&e.Player, &e.Mobs[i])
					fmt.Println(e.Player.Health)
				}
			}
		}
	}
}

/*
	func (e *Engine) MonsterCollisions() {
		for i, monster := range e.Monsters {
			if monster.Position.X > e.Player.Position.X-20 &&
				monster.Position.X < e.Player.Position.X+20 &&
				monster.Position.Y > e.Player.Position.Y-20 &&
				monster.Position.Y < e.Player.Position.Y+20 {
				e.Fight.CurrentMonster = e.Monsters[i]
				e.Fight.CurrentMonsterIndex = i
				e.StateEngine = INFIGHT
				if monster.Name == "lvl 1" {
					e.NormalTalk(monster, "Je suis Tom")
					if rl.IsKeyPressed(rl.KeySpace) { //BBBBBBBBBBBBBBBBBBBB
						fight.FightMonster(e.Player, e.Monsters[i])
						e.Fight.CurrentMonster = e.Monsters[i]
						e.Fight.CurrentMonsterIndex = i
					}
				} else if monster.Name == "lvl 2" {
					e.NormalTalk(monster, "Je suis Arnaud")
					if rl.IsKeyPressed(rl.KeySpace) {
						fight.FightMonster(e.Player, e.Monsters[1])
						e.Fight.CurrentMonster = e.Monsters[1]
						e.Fight.CurrentMonsterIndex = 1
					}
				} else if monster.Name == "lvl 3" {
					e.NormalTalk(monster, "Je suis Yann")
					if rl.IsKeyPressed(rl.KeySpace) {
						fight.FightMonster(e.Player, e.Monsters[2])
						e.Fight.CurrentMonster = e.Monsters[2]
						e.Fight.CurrentMonsterIndex = 2
					}
				} else if monster.Name == "lvl 4" {
					e.NormalTalk(monster, "Je suis Léo")
					if rl.IsKeyPressed(rl.KeySpace) {
						fight.FightMonster(e.Player, e.Monsters[3])
						e.Fight.CurrentMonster = e.Monsters[3]
						e.Fight.CurrentMonsterIndex = 3
					}
				}
			}
		}
	}
*/
func (e *Engine) ShootCollisions() {
	for _, shoot := range e.Shoot {
		if shoot.Position.X > e.Player.Position.X-10 &&
			shoot.Position.X < e.Player.Position.X+10 &&
			shoot.Position.Y > e.Player.Position.Y-10 &&
			shoot.Position.Y < e.Player.Position.Y+10 {
			if shoot.IsShooting == true && e.Player.Health > 0 {
				shoot.AttackOfShoot(&e.Player)
			}
		}
	}
}

func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}
func (e *Engine) NormalTalkMobs(m entity.Mobs, sentence string) {
	e.RenderDialogMobs(m, sentence)
}
func (e *Engine) NormalExplanation(m building.Tower, sentence string) {
	e.RenderExplanation(m, sentence)
}

func (e *Engine) NormalExplanationShop(m entity.Seller, sentence string) {
	e.RenderExplanationShop(m, sentence)
}
func (e *Engine) NormalExplanationPnj(m entity.Pnj, sentence string) {
	e.RenderExplanationPnj(m, sentence)
}
/*
func (e *Engine) ComeBackLogic() {
	if rl.IsKeyPressed(rl.KeySpace) {
		e.StateEngine = INGAME
	}
}*/

func (e *Engine) PauseLogic() {
	if rl.IsKeyPressed(rl.KeyR) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.StateMenu = HOME
		rl.StopMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) Random(tab []int) int {
	index := rand.Intn(len(tab))
	return tab[index]
}

func (e *Engine) seeShoot() int {
	return len(e.Shoot)
}
func (e *Engine) NewEngine() {
	e.Player.Position = rl.Vector2{X: 350, Y: 350}
}

func (e *Engine) CreateShoot() {
	newShoot := entity.Shoot{
		Position: rl.Vector2{
			X: float32(rand.Intn(200)),
			Y: float32(rand.Intn(200)),
		},
	}
	e.Shoot = append(e.Shoot, newShoot)
}
func (e *Engine) ShootCollision() bool {
	for i, shoot := range e.Shoot {
		if shoot.Position.X > e.Player.Position.X-20 &&
			shoot.Position.X < e.Player.Position.X+20 &&
			shoot.Position.Y > e.Player.Position.Y-20 &&
			shoot.Position.Y < e.Player.Position.Y+20 {
			fight.ShootVsPlayer(&e.Player, &e.Shoot[i])
			e.Shoot = append(e.Shoot[:i], e.Shoot[i+1:]...)
			return true
		}
	}
	return false
}

func (e *Engine) MoveShoot() { //BBBBBBBBBBBBBBBBBBBB
	directions := []rl.Vector2{
		{X: 0, Y: -5},
		{X: 0, Y: 5},
		{X: -5, Y: 0},
		{X: 5, Y: 0},
	}
	for i := 0; i < len(e.Shoot); i++ {
		randomDir := directions[e.Random([]int{0, 1, 2, 3})]
		e.Shoot[i].Position.X += randomDir.X
		e.Shoot[i].Position.Y += randomDir.Y
		if e.Shoot[i].Position.X < 0 || e.Shoot[i].Position.X > float32(200) ||
			e.Shoot[i].Position.Y < 0 || e.Shoot[i].Position.Y > float32(300) {
			e.Shoot = append(e.Shoot[:i], e.Shoot[i+1:]...)
			i--
		}
	}
}

func (e *Engine) UpdateMobs() {
	for i := 0; i < len(e.Mobs); i++ {
		if e.Mobs[i].IsAlive {
			distance := rl.Vector2Distance(e.Player.Position, e.Mobs[i].Position)
			if distance <= ChaseDistance {
				direction := rl.Vector2Subtract(e.Player.Position, e.Mobs[i].Position)
				direction = rl.Vector2Normalize(direction)
				e.Mobs[i].Position = rl.Vector2Add(e.Mobs[i].Position, direction)
			}

		}
	}
}

func (e *Engine) UpdateShoot() {
	for i := 0; i < len(e.Shoot); i++ {
		if e.Shoot[i].IsShooting {
			distance := rl.Vector2Distance(e.Player.Position, e.Shoot[i].Position)
			if distance <= ChaseDistance {
				direction := rl.Vector2Subtract(e.Player.Position, e.Shoot[i].Position)
				direction = rl.Vector2Normalize(direction)
				e.Shoot[i].Position = rl.Vector2Add(e.Shoot[i].Position, direction)
			}

		}
	}
}
