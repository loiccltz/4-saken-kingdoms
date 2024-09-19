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
	
	e.Monsters[0].Sprite = rl.LoadTexture("textures/entities/boss/animation/loup.png")
	e.Monsters[1].Sprite = rl.LoadTexture("textures/entities/boss/animation/Gryphon.png")
	e.Monsters[2].Sprite = rl.LoadTexture("textures/entities/boss/animation/Crabe.png")
	e.Monsters[3].Sprite = rl.LoadTexture("textures/entities/boss/animation/dragon.png")
	
}


func (e *Engine) UpdateAnimation() {
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

	// animation boss loup/////////////////////////////
	/*if e.Monsters[0].FrameCount % 2 == 1 {e.Monsters[0].MonsterFrame++}

	e.Monsters[0].FrameCount++

	if e.Monsters[0].MonsterFrame > 3 { e.Monsters[0].MonsterFrame = 0}

	e.Monsters[0].MonsterSrc.X = e.Player.PlayerSrc.Width * float32(e.Player.PlayerFrame)
	e.Monsters[0].MonsterSrc.Y = e.Player.PlayerSrc.Height * float32(e.Player.PlayerFrame)*/
	////////////////////////////////////////////////////
}

func loadTextures(paths []string) []rl.Texture2D {
	var textures []rl.Texture2D
	for _, path := range paths {
		texture := rl.LoadTexture(path)
		textures = append(textures, texture)
	}
	return textures
}

func (e *Engine) Unload() {
	rl.UnloadTexture(e.Player.Sprite)

	for _, sprite := range e.Sprites {
		rl.UnloadTexture(sprite)
	}

	for _, monster := range e.Monsters {
		rl.UnloadTexture(monster.Sprite)
	}
}
