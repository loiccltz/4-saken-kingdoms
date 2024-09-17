package engine

import (
	"fmt"
	"main/src/building"
	"main/src/entity"
	"main/src/fight"
	"math/rand"

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

// AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
func (e *Engine) FightLogic() {
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

	if e.Player.Health <= 0{
		e.Player.IsAlive = false
		e.Player.Money /= 2
		e.StateEngine = GAMEOVER
	}else if e.Fight.CurrentMonster.Health <= 0{
		e.Monsters = append(e.Monsters[:e.Fight.CurrentMonsterIndex], e.Monsters[e.Fight.CurrentMonsterIndex+1:]...)
		e.StateMenu = PLAY
		e.StateEngine = INGAME		
		e.Player.Inventory = append(e.Player.Inventory, e.Fight.CurrentMonster.Loot...)
		e.Player.Money += e.Fight.CurrentMonster.Worth
	}else{
		if rl.IsKeyPressed(rl.KeyE){
			e.Player.AttackOfPlayer(&e.Fight.CurrentMonster)
			e.Fight.CurrentMonster.AttackOfMonster(&e.Player)
		}
	}
	
}
func (e *Engine) InFightLogic() {
	// Mouvement
	
	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.StateEngine = PAUSE
	}

	e.FightCollisions()

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-07-Simon_s-In-There-Somewhere.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	//Menus
	if rl.IsKeyPressed(rl.KeyJ) || rl.IsKeyPressed(rl.KeyK) || rl.IsKeyPressed(rl.KeyL) || rl.IsKeyPressed(rl.KeyM) {
		e.StateFight = FIGHT
		e.StateEngine = INFIGHT
		rl.StopMusicStream(e.Music)

	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
}


func (e *Engine) GameOverLogic() {
	if e.Player.Health <= 0{
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	//Menus
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.Player.IsAlive = true
		e.StateMenu = PLAY
		e.StateEngine = INGAME
		rl.StopMusicStream(e.Music)

	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
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
	if rl.IsKeyPressed(rl.KeyQ) {
		e.StateEngine = MENUSELLER
	}

	e.CheckCollisions()

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-07-Simon_s-In-There-Somewhere.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
}




func (e *Engine) InventoryLogic() {
	if rl.IsKeyPressed(rl.KeyI) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyI){
		e.InventoryMenu = INVENTORY
	}
}
func (e Engine) MenuSellerLogic () {
	if rl.IsKeyPressed(rl.KeyQ) {
		e.SellerMenu = MENUSELLER
	}
}

func (e *Engine) CheckCollisionsWithObjects()  bool{
    playerRect := rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 40, 40)
	// * 2 - 16
    for _, obj := range e.Objects {
		objectRect := rl.NewRectangle(obj.X * 2 - 16, obj.Y * 2 - 16, obj.Width, obj.Height)
        if rl.CheckCollisionRecs(playerRect, objectRect) {
			fmt.Print("coli")
			fmt.Println(objectRect)
			return true
		} 
    }
	return false
}	
func (e *Engine) CheckCollisions() {
	e.MobsCollisions()
	e.TowerCollisions()
	e.ShopCollisions()
}
func (e *Engine) FightCollisions() {
	e.MonsterCollisions()
	//e.SpellCollision()
	e.ShootCollision()
	e.CheckCollisionsWithObjects()
}

func (e *Engine) ShopCollisions() {
	if e.Shop.Position.X > e.Player.Position.X-20 &&
		e.Shop.Position.X < e.Player.Position.X+20 &&
		e.Shop.Position.Y > e.Player.Position.Y-20 &&
		e.Shop.Position.Y < e.Player.Position.Y+20 {
		if e.Shop.Name == "Sharp Sword" {
			e.NormalExplanationShop(e.Shop, "We have all do you want to rescue the princess, press n to enter")
			if rl.IsKeyPressed(rl.KeyN) {
				e.Init_Shop()
			}
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
					e.Init_lvl1()
				}
			} else if tower.Name == "Royaume de Salkin" {
				e.NormalExplanation(tower, "To save Princess Arnaud press K")
				if rl.IsKeyPressed(rl.KeyK) {
					e.Init_lvl2()
				}
			} else if tower.Name == "Royaume d'Usun" {
				e.NormalExplanation(tower, "To save Princess Yann press L")
				if rl.IsKeyPressed(rl.KeyL) {
					e.Init_lvl3()
				}
			} else if tower.Name == "Royaume de Siroi" {
				e.NormalExplanation(tower, "To save Princess Léo press M")
				if rl.IsKeyPressed(rl.KeyM) {
					e.Init_lvl4()
				}
			}
		}
	}
}

func (e *Engine) MobsCollisions() {
	for _, mobs := range e.Mobs {
		if mobs.Position.X > e.Player.Position.X-20 &&
			mobs.Position.X < e.Player.Position.X+20 &&
			mobs.Position.Y > e.Player.Position.Y-20 &&
			mobs.Position.Y < e.Player.Position.Y+20 {
			if mobs.Name == "mob1" {
			}
		}
	}
}

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
					if rl.IsKeyPressed(rl.KeySpace) {
						//lancer un combat ?
					}
				} else if monster.Name == "lvl 2" {
					e.NormalTalk(monster, "Je suis Arnaud")
					if rl.IsKeyPressed(rl.KeySpace) {
						//lancer un combat ?
					}
				} else if monster.Name == "lvl 3" {
					e.NormalTalk(monster, "Je suis Yann")
					if rl.IsKeyPressed(rl.KeySpace) {
						//lancer un combat ?
					}
				} else if monster.Name == "lvl 4" {
					e.NormalTalk(monster, "Je suis Léo")
					if rl.IsKeyPressed(rl.KeySpace) {
						//lancer un combat ?
					}
				}
			} else {
				//...
			}
		}
	}


func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}

func (e *Engine) NormalExplanation(m building.Tower, sentence string) {
	e.RenderExplanation(m, sentence)
}

func (e *Engine) NormalExplanationShop(m building.Shop, sentence string) {
	e.RenderExplanationShop(m, sentence)
}
func (e *Engine) ComeBackLogic() {
	if rl.IsKeyPressed(rl.KeySpace) {
		e.StateEngine = INGAME
	}
}
func (e *Engine) AutoAttack(){
	if rl.IsKeyPressed(rl.KeyE) {
		e.StateEngine = INFIGHT
	}
}
// AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
func (e *Engine) PauseLogic() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyF) {
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

func (e *Engine) ShootCollision() bool {
	for i, shoot := range e.Shoot {
		if shoot.Position.X > e.Player.Position.X-20 &&
				shoot.Position.X < e.Player.Position.X+20 &&
				shoot.Position.Y > e.Player.Position.Y-20 &&
				shoot.Position.Y < e.Player.Position.Y+20 {
			fight.FightShoot(&e.Player , &e.Shoot[i])
			
			return true
		} else {
			return false
		}
	}
	return false
}

/*func (e *Engine) MoveShoot() {
	for i := 0; i < len(e.Shoot); i++ {
		e.Shoot[i].Position.Y += 5 
		if e.Shoot[i].Position.Y > e.ScreenWidthF {
			e.Shoot = append(e.Shoot[:i], e.Shoot[i+1:]...)
			i--
		}
	}
}
*/