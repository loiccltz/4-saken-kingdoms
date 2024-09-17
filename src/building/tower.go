package building

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tower struct {
	Name     string
	Position rl.Vector2
	Worth    int //valeur en argent quand tué
	Sprite   rl.Texture2D
}
