package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Pnj struct {
	Name	  string
	Position  rl.Vector2
	IsAlive   bool
}