package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Load charges les textures pour le joueur et les monstres
func (e *Engine) Load() {
	// Charge la texture initiale du joueur
	e.Player.Sprite = rl.LoadTexture("textures/entities/player/animation/S.png")
	// Charge les animations de marche pour le joueur (haut, bas, gauche, droite)
	e.Animations.Player.Walk = loadTextures([]string{
		"textures/entities/player/animation/S.png",
		"textures/entities/player/animation/Z.png",
		"textures/entities/player/animation/D.png",
		"textures/entities/player/animation/Q.png",
	})
	// Charge les textures des monstres (loup, griffon, crabe, dragon)
	e.Monsters[0].Sprite = rl.LoadTexture("textures/entities/boss/animation/loup.png")
	e.Monsters[1].Sprite = rl.LoadTexture("textures/entities/boss/animation/Gryphon.png")
	e.Monsters[2].Sprite = rl.LoadTexture("textures/entities/boss/animation/Crabe.png")
	e.Monsters[3].Sprite = rl.LoadTexture("textures/entities/boss/animation/dragon.png")

}

// UpdateAnimation met à jour l'animation du joueur selon les touches pressées
func (e *Engine) UpdateAnimation() {
	switch {

	// Si la touche W (ou la flèche haut) est pressée, l'animation est dirigée vers le haut
	case rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp):
		e.Player.Sprite = e.Animations.Player.Walk[1] //   haut

	// Si la touche S (ou la flèche bas) est pressée, l'animation est dirigée vers le bas
	case rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown):
		e.Player.Sprite = e.Animations.Player.Walk[0] // bas

	// Si la touche A (ou la flèche gauche) est pressée, l'animation est dirigée vers la gauche
	case rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft):
		e.Player.Sprite = e.Animations.Player.Walk[3] // gauche

	// Si la touche D (ou la flèche droite) est pressée, l'animation est dirigée vers la droite
	case rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight):
		e.Player.Sprite = e.Animations.Player.Walk[2] // droite
	}

	// Mise à jour du compteur d'images pour changer de frame d'animation toutes les 3 frames
	if e.Player.FrameCount%3 == 1 {
		e.Player.PlayerFrame++
	}


	if e.Player.FrameCount % 3 == 1 {e.Player.PlayerFrame++}	// vitesse de l'animation

	e.Player.FrameCount++

	if e.Player.PlayerFrame > 3 { e.Player.PlayerFrame = 0} 


	// Mise à jour de la source de l'image en fonction de la frame d'animation actuelle
	e.Player.PlayerSrc.X = e.Player.PlayerSrc.Width * float32(e.Player.PlayerFrame)
	e.Player.PlayerSrc.Y = e.Player.PlayerSrc.Height * float32(e.Player.PlayerFrame)

}

// loadTextures charge une série de textures à partir de chemins donnés
func loadTextures(paths []string) []rl.Texture2D {
	var textures []rl.Texture2D
	for _, path := range paths {
		// Charge chaque texture depuis le chemin spécifié
		texture := rl.LoadTexture(path)
		// Ajoute la texture à la liste des textures
		textures = append(textures, texture)
	}
	return textures
}

// Unload décharge toutes les textures chargées dans le jeu
func (e *Engine) Unload() {
	// Décharge la texture du joueur
	rl.UnloadTexture(e.Player.Sprite)

	// Décharge toutes les textures dans le dictionnaire de sprites
	for _, sprite := range e.Sprites {
		rl.UnloadTexture(sprite)
	}

	// Décharge les textures des projectiles tirés
	for _, shoot := range e.Shoot {
		rl.UnloadTexture(shoot.Sprite)
	}

	// Décharge les textures des mobs
	for _, mobs := range e.Mobs {
		rl.UnloadTexture(mobs.Sprite)
	}

	// Décharge les textures des monstres
	for _, monster := range e.Monsters {
		rl.UnloadTexture(monster.Sprite)
	}
}
