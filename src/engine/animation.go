package engine

import rl "github.com/gen2brain/raylib-go/raylib"

// animations du joueur
type PlayerAnimations struct {
	Walk []rl.Texture2D
}

// animations des monstres
type MonsterAnimations struct {
	Walk []rl.Texture2D
}

// animations du monde
type WorldAnimations struct {
	Tiles []rl.Texture2D
}

// avoir toutes les animations
type Animations struct {
	Player   PlayerAnimations
	Monsters MonsterAnimations
	World    WorldAnimations
}
