package building

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tower struct {
	Name     string
	Position rl.Vector2
	Worth    int 
	Sprite   rl.Texture2D
}
