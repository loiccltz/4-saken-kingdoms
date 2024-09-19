package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Load() {
	e.Player.Sprite = rl.LoadTexture("textures/entities/player/animation/S.png")
	e.Animations.Player.Walk = loadTextures([]string{
		"textures/entities/player/animation/S.png",
        "textures/entities/player/animation/Z.png",
        "textures/entities/player/animation/D.png",
		"textures/entities/player/animation/Q.png",
    })	
}
func (e *Engine) UpdateAnimation() {
	// Sélectionner la bonne texture en fonction de la direction et de la frame
	switch {
	case rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp):
		e.Player.Sprite = e.Animations.Player.Walk[1] //   haut

	case rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown):
		e.Player.Sprite = e.Animations.Player.Walk[0] //   bas

	case rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft):
		e.Player.Sprite = e.Animations.Player.Walk[3] //  gauche

	case rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight):
		e.Player.Sprite = e.Animations.Player.Walk[2] //  droite

	}
	 if e.Player.FrameCount % 3 == 1 {e.Player.PlayerFrame++}

	e.Player.FrameCount++

	if e.Player.PlayerFrame > 3 { e.Player.PlayerFrame = 0}

	e.Player.PlayerSrc.X = e.Player.PlayerSrc.Width * float32(e.Player.PlayerFrame)
	e.Player.PlayerSrc.Y = e.Player.PlayerSrc.Height * float32(e.Player.PlayerFrame)
}
func loadTextures(paths []string) []rl.Texture2D {
	// convertir en rl 2texture2d
	var textures []rl.Texture2D
	for _, path := range paths {
		texture := rl.LoadTexture(path)
		textures = append(textures, texture)
	}
	return textures
}

func (e *Engine) Unload() {
	// On libère les textures chargées, le joueur, la map, les monstres, etc...
	rl.UnloadTexture(e.Player.Sprite)

	for _, sprite := range e.Sprites {
		rl.UnloadTexture(sprite)
	}

	for _, monster := range e.Monsters {
		rl.UnloadTexture(monster.Sprite)
	}
}
