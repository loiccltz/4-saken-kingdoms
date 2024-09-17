package engine

import (
	"main/src/building"
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Rendering() {
	rl.ClearBackground(rl.Blue)
}

func (e *Engine) HomeRendering() {
	image := rl.LoadImage("4SKMENUENTRE-Photoroom.png")
    texture := rl.LoadTextureFromImage(image)
	rl.DrawTexture(texture, 0, 0, rl.White)
	
}
func(e *Engine) InventoryRendering() {
	rl.ClearBackground(rl.White)

	rl.DrawText("Invetory Menu", int32(rl.GetScreenWidth())/2-rl.MeasureText("Inventory Menu", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[I] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[I] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.Beige)
	
	rl.DrawText("ITEM 1", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 1", 20)/2, int32(rl.GetScreenHeight())/2+200, 20, rl.Black)
	rl.DrawText("ITEM 2", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 2", 20)/2, int32(rl.GetScreenHeight())/2+250, 20, rl.Black)
	rl.DrawText("ITEM 3", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 3", 20)/2, int32(rl.GetScreenHeight())/2+300, 20, rl.Black)
	
}

func (e *Engine) PauseRendering() {
	rl.ClearBackground(rl.Gray)

	rl.DrawText("Resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("Resume", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[Esc] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.Beige)
	rl.DrawText("[Q] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.Beige)

}

func(e *Engine) MenuSeller() {
	rl.ClearBackground(rl.White)

	rl.DrawText("MenuSeller", int32(rl.GetScreenWidth())/2-rl.MeasureText("MesnuSeller", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[M] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[M] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.Beige)
	
	rl.DrawText("ITEM 1", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 1", 20)/2, int32(rl.GetScreenHeight())/2+200, 20, rl.Black)
	rl.DrawText("ITEM 2", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 2", 20)/2, int32(rl.GetScreenHeight())/2+250, 20, rl.Black)
	rl.DrawText("ITEM 3", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 3", 20)/2, int32(rl.GetScreenHeight())/2+300, 20, rl.Black)
	
}


func (e *Engine) FightRendering() {
	rl.ClearBackground(rl.Blue)
	rl.BeginMode2D(e.Camera)
	e.RenderMap2()
	e.RenderShoot()
	e.RenderPlayer()
	rl.EndMode2D()
	rl.DrawText("Playing", int32(rl.GetScreenWidth())/2-rl.MeasureText("Playing", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)

}

// AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)
	rl.BeginMode2D(e.Camera)
	e.RenderMap()
	e.RenderMonsters()
	e.RenderTower()
	e.RenderShop()
	e.RenderPlayer()
	rl.EndMode2D()
	rl.DrawText("Playing", int32(rl.GetScreenWidth())/2-rl.MeasureText("Playing", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)
	e.RenderHealthBar()
	e.RenderShieldBar()
	e.RenderEnduranceBar()
	e.UpdateAndRender()
}



func (e *Engine) RenderPlayer() {
	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
}
func (e *Engine) RenderShoot() {
	for _, Shoot := range e.Shoot {
		rl.DrawTexturePro(
			Shoot.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(Shoot.Position.X, Shoot.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
	
}
func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
		rl.DrawTexturePro(
			monster.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(monster.Position.X, monster.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
}
func (e *Engine) RenderMobs() {
	for _, mobs := range e.Mobs {
		rl.DrawTexturePro(
			mobs.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(mobs.Position.X, mobs.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
}
func (e *Engine) RenderShop() {
	rl.DrawTexturePro(
		e.Shop.Sprite,
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(e.Shop.Position.X, e.Shop.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
}
func (e *Engine) RenderSeller() {
	rl.DrawTexturePro(
		e.Seller.Sprite,
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(e.Seller.Position.X, e.Seller.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
}

// AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
func (e *Engine) RenderTower() {
	for _, tower := range e.Tower {
		rl.DrawTexturePro(
			tower.Sprite,
			rl.NewRectangle(0, 0, 500, 500),
			rl.NewRectangle(tower.Position.X, tower.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
}
func (e *Engine) RenderDialog(m entity.Monster, sentence string) {
	rl.BeginMode2D(e.Camera)
	rl.DrawText(
		sentence,
		int32(m.Position.X),
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)
	rl.EndMode2D()
}

// AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
func (e *Engine) RenderExplanation(m building.Tower, sentence string) {
	rl.BeginMode2D(e.Camera)
	rl.DrawText(
		sentence,
		int32(m.Position.X),
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)
	rl.EndMode2D()
}

func (e *Engine) RenderExplanationShop(m building.Shop, sentence string) {
	rl.BeginMode2D(e.Camera)
	rl.DrawText(
		sentence,
		int32(m.Position.X),
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)
	rl.EndMode2D()
}

func (e *Engine) GameOverRendering() {
	rl.ClearBackground(rl.Black)
	rl.DrawText("You DIED", int32(rl.GetScreenWidth())/2-rl.MeasureText("Playing", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[Esc] to Exit or [Enter] to Replay", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)
}

func (e *Engine) RenderHealthBar() {
    screenHeight := int32(rl.GetScreenHeight())

    // Dimensions et positions de la barre de vie
    barWidth := int32(150)  // Largeur de la barre de vie
    barHeight := int32(20)  // Hauteur de la barre de vie
    barX := int32(20)       // 20 pixels de marge depuis le bord gauche
    barY := screenHeight - barHeight - 20 // 20 pixels de marge depuis le bord bas

    // Calcul de la largeur de la barre de vie en fonction de la santé
    healthRatio := float32(e.Player.Health) / float32(e.Player.MaxHealth) // Supposons que la vie maximale est 100
    if healthRatio > 1 {
        healthRatio = 1
    } else if healthRatio < 0 {
        healthRatio = 0
    }
    healthBarWidth := int32(float32(barWidth) * healthRatio)

    // Dessine le fond de la barre de vie
    rl.DrawRectangle(barX, barY, barWidth, barHeight, rl.Gray)

    // Dessine la barre de vie
    rl.DrawRectangle(barX, barY, healthBarWidth, barHeight, rl.Red)
}

func (e *Engine) UpdateAndRender() {
    // Mise à jour de l'endurance du joueur
    e.Player.UpdateEndurance()

    // Vérifie si la touche Q est pressée pour effectuer une action
    if rl.IsKeyPressed(rl.KeyQ) {
        if e.Player.Endurance >= e.Player.MaxEndurance {
            fmt.Println("Action effectuée !")
            e.Player.Endurance = 0
        }
    }
}

func (e *Engine) RenderEnduranceBar() {
    // Calcul de la largeur de la barre d'endurance en fonction de l'endurance
    screenWidth := int32(rl.GetScreenWidth())
    screenHeight := int32(rl.GetScreenHeight())
    barWidth := int32(150) // Largeur de la barre d'endurance
    barHeight := int32(20) // Hauteur de la barre d'endurance
    barX := screenWidth - barWidth - 20
    barY := screenHeight - barHeight - 20

    // Ratio de l'endurance pour déterminer la largeur de la barre
    enduranceRatio := float32(e.Player.Endurance) / float32(e.Player.MaxEndurance)
    if enduranceRatio > 1 {
        enduranceRatio = 1
    } else if enduranceRatio < 0 {
        enduranceRatio = 0
    }
    enduranceBarWidth := int32(float32(barWidth) * enduranceRatio)

    // Dessine le fond et la barre d'endurance
    rl.DrawRectangle(barX, barY, barWidth, barHeight, rl.Gray)
    rl.DrawRectangle(barX, barY, enduranceBarWidth, barHeight, rl.Yellow)
}

func (e *Engine) UpdateAndRender2() {
    // Mise à jour du bouclier du joueur
    e.Player.UpdateShield()

    // Vérifie si la touche Q est pressée pour effectuer une action
    if rl.IsKeyPressed(rl.KeyQ) {
        if e.Player.Shield >= e.Player.MaxShield {
            fmt.Println("Action effectuée !")
            e.Player.Shield = 0
        }
    }
}

func (e *Engine) RenderShieldBar() {
    screenHeight := int32(rl.GetScreenHeight())

    // Dimensions de la barre de bouclier
    barWidth := int32(150) // Largeur de la barre de bouclier
    barHeight := int32(20) // Hauteur de la barre de bouclier

    // Positionnement en bas à gauche avec une marge de 20 pixels
    barX := int32(20)
    barY := screenHeight - barHeight - 60 // Position modifiée pour la barre de bouclier

    // Ratio du bouclier pour déterminer la largeur de la barre
    shieldRatio := float32(e.Player.Shield) / float32(e.Player.MaxShield)
    if shieldRatio > 1 {
        shieldRatio = 1
    } else if shieldRatio < 0 {
        shieldRatio = 0
    }
    shieldBarWidth := int32(float32(barWidth) * shieldRatio)

    // Dessine le fond et la barre de bouclier
    rl.DrawRectangle(barX, barY, barWidth, barHeight, rl.Gray)
    rl.DrawRectangle(barX, barY, shieldBarWidth, barHeight, rl.Blue)

    // Optionnel: Affichage de l'état du bouclier dans la console
    fmt.Printf("Shield: %d / %d\n", e.Player.Shield, e.Player.MaxShield)
}
