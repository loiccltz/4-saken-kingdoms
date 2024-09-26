package engine

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"main/src/building"
	"main/src/entity"
	"main/src/fight"
)

func (e *Engine) HomeLogic() {
	fmt.Println()
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
	// sprint qui fonctionne pour l'instant comme un noclip
	/*
		if rl.IsKeyDown(rl.KeyW) && rl.IsKeyDown(rl.KeyLeftShift) {
			e.Player.Position.Y -= e.Player.Speed - 2
		}
		if rl.IsKeyDown(rl.KeyS) && rl.IsKeyDown(rl.KeyLeftShift) {
			e.Player.Position.Y += e.Player.Speed + 2
		}
		if rl.IsKeyDown(rl.KeyA) && rl.IsKeyDown(rl.KeyLeftShift) {
			e.Player.Position.X -= e.Player.Speed - 2
		}
		if rl.IsKeyDown(rl.KeyD) && rl.IsKeyDown(rl.KeyLeftShift) {
			e.Player.Position.X += e.Player.Speed + 2
		}
	*/
	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: float32(ScreenWidth) / 2, Y: float32(ScreenHeight) / 2}

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

	}
	//e.FightCollisions()
	posX := e.Player.Position.X
	posY := e.Player.Position.Y
	if posX == 356 && posY == 200 {
		for e.Monsters[0].Health > 0 {
			e.ShootLogic()
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
	playerRect := rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 40, 40) // Je trace un rectangle autour du joueur

	for _, obj := range e.Objects { // je parcours chaque objet
		// Je trace un rectangle au coordoonées de  l'objet, avec sa taille, on fait *2-16 car notre pack d'assets est en 16x16
		objectRect := rl.NewRectangle(obj.X*2-16, obj.Y*2-16, obj.Width*2-16, obj.Height*2-16)
		if rl.CheckCollisionRecs(playerRect, objectRect) { // on peut maintenant regarder si il y a collisions entre les deux
			return true
		}
	}
	return false
}

func (e *Engine) CheckCollisions() {
	e.MobsCollisions()
	e.ShootCollisions()
	e.UpdateMobs()
	e.PnjCollisions()
	e.TowerCollisions()
	e.SellerCollisions()
	e.CheckCollisionsWithObjects()
	e.BlockCollisions()
}
func (e *Engine) BlockCollisions() {
	if e.CheckCollisionsWithObjects() {
		// si il y a une collision, on bloque le joueur
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

/*func (e *Engine) FightCollisions() {
	e.UpdateShoot()
	e.MoveShoot()
	e.UpdateShoot()
	//e.SpellCollision()
}*/

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

	for i := range e.Pnj {

		if e.Pnj[i].Position.X > e.Player.Position.X-20 &&
			e.Pnj[i].Position.X < e.Player.Position.X+20 &&
			e.Pnj[i].Position.Y > e.Player.Position.Y-20 &&
			e.Pnj[i].Position.Y < e.Player.Position.Y+20 {
			// On stock le message crypté
			if e.Pnj[i].Name == "Jack" {
				var cipherSentence string
				sentence := "Bonjour aventurier, explore ce monde et libère les princesses"
				runes := []rune(sentence)
				for _, r := range runes {
					// On décale chaque lettres de 1 ( methode césar)
					cipherRune := r + 1
					cipherSentence += string(cipherRune)
				}
				e.CypherTalk(e.Pnj[i], cipherSentence)
			}
			if e.Pnj[i].Name == "Jacky" {
				e.RenderExplanationPnj(e.Pnj[i],
					"Bonjour aventurier, je vais t'aider à traduire les messages de ce monde : '\n' Explore ce monde et libère les princesses, suis les différents chemins si tu es perdu ")
			}

			if e.Pnj[i].Name == "Michou" {
				var RobotSentence string
				sentence := "Bonjour étranger tu te dirige vers le chateau !"
				//
				for _, runes := range sentence {
					// Convertir la rune (int32) en ASCII
					asciiValue := int(runes)
					// convertir l'ASCII en chaîne binaire
					binaryString := fmt.Sprintf("%08b", asciiValue)
					// on ajoute la valeur en binaire
					RobotSentence += binaryString
				}
				// Appeler la méthode pour faire parler binaire
				e.RobotTalk(e.Pnj[i], RobotSentence)
			}

			if e.Pnj[i].Name == "Francis" {
				e.RenderExplanationPnj(e.Pnj[i], "Tu te dirige vers le chateau !")
			}

			if e.Pnj[i].Name == "Garde" {
				e.RenderExplanationPnj(e.Pnj[i], "Saluez le Roi Léo et la reine Yannette !")
			}

		}
	}
}
func (e *Engine) CheckCollisionsWithSquare() bool {
	for i, _ := range e.Shoot {
		FightRect := rl.NewRectangle(e.Shoot[i].Position.X, e.Shoot[i].Position.Y, 40, 40)
		// * 2 - 16
		for _, fight := range e.BossFight {
			objectRect := rl.NewRectangle(fight.X*2-16, fight.Y*2-16, fight.Width, fight.Height)
			if rl.CheckCollisionRecs(FightRect, objectRect) {
				rl.UnloadTexture(e.Shoot[i].Sprite)
				return true
			}
		}
	}
	return false
}

func (e *Engine) UseSelectedItem() {
	// Ajoute le montant de régénération de l'objet selectionné
	e.Player.Health += e.Player.Inventory[e.selectedIndex].Regen

	// Vérifie si la santé dépasse la santé maximale
	if e.Player.Health > e.Player.MaxHealth {
		// Si c'est le cas, remet la santé du joueur à la santé maximale
		e.Player.Health = e.Player.MaxHealth
	}
}

func (e *Engine) TowerCollisions() {
	for _, tower := range e.Tower {
		if tower.Position.X > e.Player.Position.X-100 &&
			tower.Position.X < e.Player.Position.X+100 &&
			tower.Position.Y > e.Player.Position.Y-100 &&
			tower.Position.Y < e.Player.Position.Y+100 {
			if tower.Name == "Royaume de Ran" {
				e.NormalExplanation(tower, "To save Princess Tom press J")
				if rl.IsKeyPressed(rl.KeyJ) {
					e.Player.Position.X = 3116
					e.Player.Position.Y = 5046
				}
			} else if tower.Name == "Royaume de Salkin" {
				e.NormalExplanation(tower, "To save Princess Arnaud press K")
				if rl.IsKeyPressed(rl.KeyK) {
					e.Player.Position.X = 3152
					e.Player.Position.Y = 7269
				}
			} else if tower.Name == "Royaume d'Usun" {
				e.NormalExplanation(tower, "To save Princess Yann press L")
				if rl.IsKeyPressed(rl.KeyL) {
					e.Player.Position.X = 7517
					e.Player.Position.Y = 4128
				}
			} else if tower.Name == "Royaume de Siroi" {
				e.NormalExplanation(tower, "To save Princess Léo press M")
				if rl.IsKeyPressed(rl.KeyM) {
					e.Player.Position.X = 7449
					e.Player.Position.Y = 7598
				}
			}
		}
	}
}

func (e *Engine) MobsCollisions() {

	e.Player.UpdateEndurance()

	for i := range e.Mobs {
		if e.Mobs[i].IsAlive {
			if e.Mobs[i].Position.X > e.Player.Position.X-20 &&
				e.Mobs[i].Position.X < e.Player.Position.X+20 &&
				e.Mobs[i].Position.Y > e.Player.Position.Y-20 &&
				e.Mobs[i].Position.Y < e.Player.Position.Y+20 {

				fmt.Println(e.Mobs[i].Health)

				if e.Player.IsAlive {

					e.ApplyDamageToPlayer(e.Mobs[i].Damage)

					if rl.IsKeyPressed(rl.KeyEnter) {
						if e.Player.Endurance >= e.Player.MaxEndurance {

							fight.PlayerVsMobs(&e.Player, &e.Mobs[i])

							e.Player.Endurance = 0
							fmt.Println(e.Player.Health)
						} else {
							fmt.Println("Endurance insuffisante pour attaquer")
						}
					}
				}
			}
		}
	}

	for i, monster := range e.Monsters {
		if monster.IsAlive {
			if monster.Position.X > e.Player.Position.X-300 &&
				monster.Position.X < e.Player.Position.X+300 &&
				monster.Position.Y > e.Player.Position.Y-300 &&
				monster.Position.Y < e.Player.Position.Y+300 {
				e.ShootLogic()
				fmt.Println(monster.Health)

				if e.Player.IsAlive {
					for i, _ := range e.Shoot {
						if e.Shoot[i].Position.X > e.Player.Position.X-10 &&
							e.Shoot[i].Position.X < e.Player.Position.X+10 &&
							e.Shoot[i].Position.Y > e.Player.Position.Y-10 &&
							e.Shoot[i].Position.Y < e.Player.Position.Y+10 {
							if e.Shoot[i].IsShooting && e.Player.IsAlive {
								e.ApplyDamageToPlayer(e.Shoot[i].Damage)
								e.Shoot[i].IsShooting = false
								e.Shoot = append(e.Shoot[:i], e.Shoot[i+1:]...)
							}
						}
					}
					if rl.IsKeyPressed(rl.KeyEnter) {
						if e.Player.Endurance >= e.Player.MaxEndurance {
							fight.PlayerVsMonster(&e.Player, &e.Monsters[i])
							e.Player.Endurance = 0
							fmt.Println(e.Player.Health)
						} else {
							fmt.Println("Endurance insuffisante pour attaquer")
						}
					}
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
			if shoot.IsShooting && e.Player.IsAlive {
				e.ApplyDamageToPlayer(shoot.Damage)
			}
		}
	}
}

func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}
func (e *Engine) CypherTalk(pnj entity.Pnj, sentence string) {
	e.RenderExplanationPnj(pnj, sentence)
}

func (e *Engine) RobotTalk(pnj entity.Pnj, sentence string) {
	e.RenderExplanationPnj(pnj, sentence)
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

// met à jour la position des projectiles
func (e *Engine) UpdateShoot() {
	// Parcourt tous les projectiles dans la liste e.Shoot
	for i := 0; i < len(e.Shoot); i++ {
		// Vérifie si le projectile est en train d'être tiré
		if e.Shoot[i].IsShooting {
			// Calcule la distance entre le joueur et le projectile
			distance := rl.Vector2Distance(e.Player.Position, e.Shoot[i].Position)
			// Si le projectile est à une distance suffisante du joueur, commence à le suivre
			if distance <= ChaseDistance {
				// Calcule la direction du projectile vers le joueur
				direction := rl.Vector2Subtract(e.Player.Position, e.Shoot[i].Position)

				// Normalise la direction pour obtenir un vecteur de longueur 1
				direction = rl.Vector2Normalize(direction)

				// Met à jour la position du projectile en ajoutant la direction normalisée à sa position actuelle
				e.Shoot[i].Position = rl.Vector2Add(e.Shoot[i].Position, direction)
			}
		}
	}
}
func (e *Engine) ShootLogic() {
	// Gestion des tirs existants
	rl.BeginMode2D(e.Camera)

	for _, shoot := range e.Shoot {
		// Mise à jour de la position du tir en fonction de sa direction
		switch shoot.Direction {
		case 0: // Haut
			shoot.Position.Y -= 5
		case 1: // Bas
			shoot.Position.Y += 5
		case 2: // Gauche
			shoot.Position.X -= 5
		case 3: // Droite
			shoot.Position.X += 5
		}

		/*/ Vérifier si le tir sort des limites de l'écran ou de la zone de combat
		if shoot.Position.X < e.BossFight[0].X || shoot.Position.X > e.BossFight[0].Width||
			shoot.Position.Y < e.BossFight[0].Y || shoot.Position.Y > e.BossFight[0].Height {
			// Si le tir sort de l'écran, on le supprime
			e.Shoot = append(e.Shoot[:i], e.Shoot[i+1:]...)
			i--
		} */
	}
	e.RenderShoot()
	rl.EndMode2D()
}

/*
const (

	boxSize    = 50
	ShootSpeed = 6

)

	func (e *Engine) ShootLogic() {
		boxX := float32(ScreenWidth/2 - boxSize/2)
		boxY := float32(ScreenHeight/2 - boxSize/2)
		rand.Seed(time.Now().UnixNano())
		for !rl.WindowShouldClose() {
			// Ajouter une nouvelle flèche aléatoirement
			if rand.Float32() < 0.6 { // 1% de chance par frame d'ajouter une flèche
				shoot := entity.Shoot{
					Position:   rl.Vector2{X: boxX + boxSize/2, Y: boxY + boxSize/2},
					IsShooting: true,
					Direction:  rand.Intn(4), // Direction aléatoire
					Sprite:     rl.LoadTexture("textures/fefolet.png"),
				}
				e.Shoot = append(e.Shoot, shoot)
			}

			// Mise à jour des flèches
			for i := 0; i < len(e.Shoot); i++ {
				switch e.Shoot[i].Direction {
				case 0: // Haut
					e.Shoot[i].Position.Y -= ShootSpeed
				case 1: // Bas
					e.Shoot[i].Position.Y += ShootSpeed
				case 2: // Gauche
					e.Shoot[i].Position.X -= ShootSpeed
				case 3: // Droite
					e.Shoot[i].Position.X += ShootSpeed
				}

				// Retirer les flèches qui sortent de l'écran
				if e.Shoot[i].Position.X < 0 || e.Shoot[i].Position.X > float32(ScreenWidth) ||
					e.Shoot[i].Position.Y < 0 || e.Shoot[i].Position.Y > float32(ScreenHeight) {
					e.Shoot = append(e.Shoot[:i], e.Shoot[i+1:]...)
					i--
				}
			}

			// Dessin à l'écran
			rl.BeginDrawing()
			rl.DrawRectangle(int32(boxX), int32(boxY), boxSize, boxSize, rl.DarkGray)

			// Dessiner les flèches
			for _, shoot := range e.Shoot {
				switch shoot.Direction {
				case 0: // Haut
					rl.DrawTriangle(rl.Vector2{shoot.Position.X, shoot.Position.Y}, rl.Vector2{shoot.Position.X - 10, shoot.Position.Y + 20}, rl.Vector2{shoot.Position.X + 10, shoot.Position.Y + 20}, rl.Red)
				case 1: // Bas
					rl.DrawTriangle(rl.Vector2{shoot.Position.X, shoot.Position.Y}, rl.Vector2{shoot.Position.X - 10, shoot.Position.Y - 20}, rl.Vector2{shoot.Position.X + 10, shoot.Position.Y - 20}, rl.Red)
				case 2: // Gauche
					rl.DrawTriangle(rl.Vector2{shoot.Position.X, shoot.Position.Y}, rl.Vector2{shoot.Position.X + 20, shoot.Position.Y - 10}, rl.Vector2{shoot.Position.X + 20, shoot.Position.Y + 10}, rl.Red)
				case 3: // Droite
					rl.DrawTriangle(rl.Vector2{shoot.Position.X, shoot.Position.Y}, rl.Vector2{shoot.Position.X - 20, shoot.Position.Y - 10}, rl.Vector2{shoot.Position.X - 20, shoot.Position.Y + 10}, rl.Red)
				}
			}

			rl.EndDrawing()
		}
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
*/
