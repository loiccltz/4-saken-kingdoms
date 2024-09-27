package engine

import (
	"fmt"
	"math/rand"
	"time"
	"main/src/building"
	"main/src/entity"
	"main/src/fight"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// HomeLogic gère la logique de l'écran d'accueil (menu principal)
func (e *Engine) HomeLogic() {

	// Musique : démarre la musique si elle n'est pas déjà en cours
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")
		rl.PlayMusicStream(e.Music)
	}

	// Mise à jour de la musique
	rl.UpdateMusicStream(e.Music)

	// Menus : commence le jeu ou ferme l'application selon les touches pressées
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY          // Passage à l'état de jeu
		e.StateEngine = INGAME      // Moteur en mode de jeu
		rl.StopMusicStream(e.Music) // Arrête la musique
	}

	// Quitter le jeu selon les touches pressées
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
}

// SettingsLogic gère la logique du menu des paramètres
func (e *Engine) SettingsLogic() {

	// Menus : retour à l'écran d'accueil si la touche B est pressée
	if rl.IsKeyPressed(rl.KeyB) {
		e.StateMenu = HOME
	}

	// Mise à jour de la musique
	rl.UpdateMusicStream(e.Music)
}

// InGameLogic gère la logique du jeu en cours (déplacements, interactions, etc.)
func (e *Engine) InGameLogic() {

	// Mouvement du joueur selon les touches directionnelles
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		e.Player.Position.Y -= e.Player.Speed // Déplacement vers le haut
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		e.Player.Position.Y += e.Player.Speed // Déplacement vers le bas
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		e.Player.Position.X -= e.Player.Speed // Déplacement vers la gauche
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		e.Player.Position.X += e.Player.Speed // Déplacement vers la droite
	}


	// Mise à jour de la caméra pour suivre le joueur
	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: float32(ScreenWidth) / 2, Y: float32(ScreenHeight) / 2}

	// Menus : navigation vers le menu pause, inventaire, ou vendeur
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.StateEngine = PAUSE
	}
	if rl.IsKeyPressed(rl.KeyI) {
		e.StateEngine = INVENTORY
	}
	if rl.IsKeyPressed(rl.KeyR) {
		e.StateEngine = MENUSELLER
	}

	// Gestion des collisions dans le jeu
	e.CheckCollisions()

	// Musique : démarre la musique si elle n'est pas en cours
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-07-Simon_s-In-There-Somewhere.mp3")
		rl.PlayMusicStream(e.Music)
	}

	// Mise à jour de la musique
	rl.UpdateMusicStream(e.Music)

	// Vérification si le joueur est mort
	if e.Player.Health <= 0 {
		e.Player.IsAlive = false
		e.Player.Money /= 2      // Pénalité en argent
		e.StateEngine = GAMEOVER // Transition à l'écran de Game Over
	}

	// Vérification si les mobs sont morts
	for i := range e.Mobs {
		if e.Mobs[i].Health <= 0 {
			e.Mobs[i].IsAlive = false
		}
	}

	// Vérifie si le monstre 0 est en vie, puis déclenche la logique de tir
	if e.Monsters[0].Health > 0 {
		e.ShootLogic()
	}
	// Si le monstre 0 est mort, téléporte le joueur à une position spécifique
	if e.Monsters[0].Health <= 0 {
		e.Player.Position.X = 4648
		e.Player.Position.Y = 6670
	}

	// Même logique pour le monstre 1
	if e.Monsters[1].Health > 0 {
		e.ShootLogic()
	}
	if e.Monsters[1].Health <= 0 {
		e.Player.Position.X = 4648
		e.Player.Position.Y = 6670
	}

	// Même logique pour le monstre 2
	if e.Monsters[2].Health > 0 {
		e.ShootLogic()
	}
	if e.Monsters[2].Health <= 0 {
		e.Player.Position.X = 4648
		e.Player.Position.Y = 6670
	}

	// Même logique pour le monstre 3
	if e.Monsters[3].Health > 0 {
		e.ShootLogic()
	}
	if e.Monsters[3].Health <= 0 {
		e.Player.Position.X = 4648
		e.Player.Position.Y = 6670
	}
}

// Logique de fin de jeu
func (e *Engine) GameOverLogic() {
	// Si la santé du joueur est à 0 ou moins, démarrer la musique de fin
	if e.Player.Health <= 0 {
		if !rl.IsMusicStreamPlaying(e.Music) {
			e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")
			rl.PlayMusicStream(e.Music)
		}
		rl.UpdateMusicStream(e.Music)

		// Logique des menus : si la touche Entrée est pressée, relancer le joueur
		if rl.IsKeyPressed(rl.KeyEnter) {
			e.Player.IsAlive = true
			e.Player.Health = e.Player.MaxHealth
			e.StateMenu = PLAY
			e.StateEngine = INGAME
			rl.StopMusicStream(e.Music)
		}

		// Si la touche Échap est pressée, quitter le jeu
		if rl.IsKeyPressed(rl.KeyEscape) {
			e.IsRunning = false
		}
	}
}

// Logique d'ouverture de l'inventaire
func (e *Engine) InventoryLogic() {
	if rl.IsKeyPressed(rl.KeyI) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyI) {
		e.InventoryMenu = INVENTORY
	}
}

// Logique d'ouverture du menu du vendeur
func (e *Engine) SellerLogic() {
	if rl.IsKeyPressed(rl.KeyR) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyR) {
		e.SellerMenu = MENUSELLER
	}
}

// Vérifie les collisions entre le joueur et les objets
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

// Gère les différentes collisions dans le jeu
func (e *Engine) CheckCollisions() {
	e.MobsCollisions()
	e.ShootCollisions()
	e.UpdateMobs()
	e.UpdateShoot()
	e.PnjCollisions()
	e.TowerCollisions()
	e.SellerCollisions()
	e.CheckCollisionsWithObjects()
	e.BlockCollisions()
}

// Gère les collisions entre le joueur et les obstacles qui le bloquent
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

// Gère les collisions entre le joueur et le vendeur
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

// Gère les collisions avec les PNJ et les interactions associées
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

// Vérifie les collisions entre les tirs et un carré
func (e *Engine) CheckCollisionsWithSquare() bool {
	for i, _ := range e.Shoot {
		FightRect := rl.NewRectangle(e.Shoot[i].Position.X, e.Shoot[i].Position.Y, 40, 40)
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
	// Génère un nombre aléatoire entre 0 et 99
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)

	// Si la chance est, par exemple, inférieure à 10, le joueur meurt
	if chance < 10 {
		e.Player.Health = 1
	}

	// Sinon, la potion régénère la santé comme d'habitude
	e.Player.Health += e.Player.Inventory[e.selectedIndex].Regen

	// Vérifie si la santé dépasse la santé maximale
	if e.Player.Health > e.Player.MaxHealth {
		// Si c'est le cas, remet la santé du joueur à la santé maximale
		e.Player.Health = e.Player.MaxHealth
	}
}

// Gère les collisions entre le joueur et les tours
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
				if rl.IsKeyPressed(rl.KeyN) {
					e.Player.Position.X = 7449
					e.Player.Position.Y = 7598
				}
			}
		}
	}
}

// Gère les collisions entre les mobs (ennemis) et le joueur
func (e *Engine) MobsCollisions() {

	// Met à jour l'endurance du joueur (peut-être utilisée pour les attaques)
	e.Player.UpdateEndurance()

	// Boucle à travers tous les mobs pour vérifier s'ils sont en vie et proches du joueur
	for i := range e.Mobs {
		if e.Mobs[i].IsAlive {
			// Vérifie si un mob est suffisamment proche du joueur pour une interaction
			if e.Mobs[i].Position.X > e.Player.Position.X-20 &&
				e.Mobs[i].Position.X < e.Player.Position.X+20 &&
				e.Mobs[i].Position.Y > e.Player.Position.Y-20 &&
				e.Mobs[i].Position.Y < e.Player.Position.Y+20 {

				// Si le joueur est vivant, applique les dégâts infligés par le mob
				if e.Player.IsAlive {
					e.ApplyDamageToPlayer(e.Mobs[i].Damage)

					// Si le joueur appuie sur "Entrée", déclenche un combat entre le joueur et le mob
					if rl.IsKeyPressed(rl.KeyEnter) {
						// Vérifie si le joueur a assez d'endurance pour attaquer
						if e.Player.Endurance >= e.Player.MaxEndurance {
							// Appelle la fonction qui gère le combat entre le joueur et le mob
							fight.PlayerVsMobs(&e.Player, &e.Mobs[i])
							// Réinitialise l'endurance du joueur après l'attaque
							e.Player.Endurance = 0
						} else {
							// Message d'erreur si l'endurance est insuffisante
							fmt.Println("Endurance insuffisante pour attaquer")
						}
					}
				}
			}
		}
	}

	// Boucle pour gérer les collisions avec les "monstres" (probablement des boss ou ennemis plus puissants)
	for i, monster := range e.Monsters {
		if monster.IsAlive {
			// Vérifie si un monstre est dans une certaine distance du joueur (300 unités ici)
			if monster.Position.X > e.Player.Position.X-300 &&
				monster.Position.X < e.Player.Position.X+300 &&
				monster.Position.Y > e.Player.Position.Y-300 &&
				monster.Position.Y < e.Player.Position.Y+300 {

				// Déclenche la logique de tir si un monstre est proche du joueur
				e.ShootLogic()

				// Si le joueur est vivant, vérifie si un tir touche le joueur
				if e.Player.IsAlive {
					for i, _ := range e.Shoot {
						if e.Shoot[i].Position.X > e.Player.Position.X-10 &&
							e.Shoot[i].Position.X < e.Player.Position.X+10 &&
							e.Shoot[i].Position.Y > e.Player.Position.Y-10 &&
							e.Shoot[i].Position.Y < e.Player.Position.Y+10 {

							// Si un tir est en cours et touche le joueur, applique des dégâts
							if e.Shoot[i].IsShooting && e.Player.IsAlive {
								e.ApplyDamageToPlayer(e.Shoot[i].Damage)
								// Désactive le tir une fois qu'il a touché le joueur
								e.Shoot[i].IsShooting = false
								// Supprime le tir de la liste des tirs
								e.Shoot = append(e.Shoot[:i], e.Shoot[i+1:]...)
							}
						}
					}

					// Si le joueur appuie sur "Entrée", vérifie à nouveau l'endurance pour attaquer le monstre
					if rl.IsKeyPressed(rl.KeyEnter) {
						if e.Player.Endurance >= e.Player.MaxEndurance {
							// Gère le combat entre le joueur et le monstre
							fight.PlayerVsMonster(&e.Player, &e.Monsters[i])
							// Réinitialise l'endurance du joueur après l'attaque
							e.Player.Endurance = 0
						} else {
							// Message d'erreur si l'endurance est insuffisante
							fmt.Println("Endurance insuffisante pour attaquer")
						}
					}
				}
			}
		}
	}
}

// Gère les collisions entre les tirs et le joueur
func (e *Engine) ShootCollisions() {
	// Parcourt chaque tir en cours dans le jeu
	for _, shoot := range e.Shoot {
		// Vérifie si un tir touche le joueur en comparant leurs positions
		if shoot.Position.X > e.Player.Position.X-10 &&
			shoot.Position.X < e.Player.Position.X+10 &&
			shoot.Position.Y > e.Player.Position.Y-10 &&
			shoot.Position.Y < e.Player.Position.Y+10 {
			// Si un tir est en cours et touche un joueur vivant, applique des dégâts
			if shoot.IsShooting && e.Player.IsAlive {
				e.ApplyDamageToPlayer(shoot.Damage)
			}
		}
	}
}

// Affiche un dialogue normal pour les monstres
func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}

// Affiche un dialogue chiffré pour les PNJ (système de cryptage de message)
func (e *Engine) CypherTalk(pnj entity.Pnj, sentence string) {
	e.RenderExplanationPnj(pnj, sentence)
}

func (e *Engine) RobotTalk(pnj entity.Pnj, sentence string) {
	e.RenderExplanationPnj(pnj, sentence)
}

// Affiche un dialogue normal pour les mobs (petits ennemis)
func (e *Engine) NormalTalkMobs(m entity.Mobs, sentence string) {
	e.RenderDialogMobs(m, sentence)
}

// Affiche une explication normale pour une tour (probablement des structures dans le jeu)
func (e *Engine) NormalExplanation(m building.Tower, sentence string) {
	e.RenderExplanation(m, sentence)
}

// Affiche une explication normale pour un vendeur
func (e *Engine) NormalExplanationShop(m entity.Seller, sentence string) {
	e.RenderExplanationShop(m, sentence)
}
// Gère la logique de pause du jeu (retour au menu ou reprise)
func (e *Engine) PauseLogic() {
	// Si la touche "R" est appuyée, le jeu reprend
	if rl.IsKeyPressed(rl.KeyR) {
		e.StateEngine = INGAME
	}
	// Si "Échap" est appuyé, retourne au menu principal
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.StateMenu = HOME
		rl.StopMusicStream(e.Music)
	}
	// Mise à jour de la musique en mode pause
	rl.UpdateMusicStream(e.Music)
}

// Met à jour les positions des mobs en fonction de la distance par rapport au joueur
func (e *Engine) UpdateMobs() {
	// Parcourt chaque mob vivant
	for i := 0; i < len(e.Mobs); i++ {
		if e.Mobs[i].IsAlive {
			// Calcule la distance entre le joueur et le mob
			distance := rl.Vector2Distance(e.Player.Position, e.Mobs[i].Position)
			// Si le mob est suffisamment proche, il poursuit le joueur
			if distance <= ChaseDistance {
				direction := rl.Vector2Subtract(e.Player.Position, e.Mobs[i].Position)
				direction = rl.Vector2Normalize(direction)
				e.Mobs[i].Position = rl.Vector2Add(e.Mobs[i].Position, direction)
			}
		}
	}
}


// Met à jour la position des tirs en fonction de la distance par rapport au joueur
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

// Gère la logique des tirs (mise à jour de leurs positions et affichage)
func (e *Engine) ShootLogic() {
	// Commence un mode 2D pour gérer les tirs
	rl.BeginMode2D(e.Camera)

	// Parcourt chaque tir et met à jour sa position selon sa direction
	for _, shoot := range e.Shoot {
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
	}
	// Affiche les tirs à l'écran
	e.RenderShoot()
	// Fin du mode 2D
	rl.EndMode2D()
}
