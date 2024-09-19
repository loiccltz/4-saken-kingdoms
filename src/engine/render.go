package engine

import (
	"main/src/building"
	"main/src/entity"
	"fmt"

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

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)
	rl.BeginMode2D(e.Camera)
	e.RenderMap()
	e.RenderTower()
	e.RenderMobs()
	e.RenderPlayer()
	rl.EndMode2D()
	
	rl.DrawText("Playing", int32(rl.GetScreenWidth())/2-rl.MeasureText("Playing", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)
	rl.DrawText("[Q]/[A] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)
	
	e.RenderHealthBar()
	e.RenderEnduranceBar()
	e.UpdateAnimation()
	e.RenderShieldBar()
	e.UpdateAndRenderEndurance()
	e.UpdateAndRenderShield()
	e.RenderSeller()
	e.RenderMonsters()
}

func(e *Engine) InventoryRendering() {
	rl.ClearBackground(rl.Black)

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

func(e *Engine) SellerRendering() {
	rl.ClearBackground(rl.Beige)

	rl.DrawText("MenuSeller", int32(rl.GetScreenWidth())/2-rl.MeasureText("MesnuSeller", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[M] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[M] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.White)
	
	rl.DrawText("ITEM 1", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 1", 20)/2, int32(rl.GetScreenHeight())/2+200, 20, rl.Black)
	rl.DrawText("ITEM 2", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 2", 20)/2, int32(rl.GetScreenHeight())/2+250, 20, rl.Black)
	rl.DrawText("ITEM 3", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 3", 20)/2, int32(rl.GetScreenHeight())/2+300, 20, rl.Black)
	
}



func (e *Engine) InFightRendering() {
	rl.ClearBackground(rl.Gray)

	e.RenderHealthBar()
	e.RenderMobs()
	e.RenderTower()
	e.RenderSeller()
	e.RenderMonsters()
	e.RenderShoot()
	rl.EndMode2D()
	rl.DrawText("Playing", int32(rl.GetScreenWidth())/2-rl.MeasureText("Playing", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)
}

func (e *Engine) RenderPlayer(){
	rl.BeginMode2D(e.Camera)
	//rl.DrawTexturePro(e.Player.Sprite, e.Player.PlayerSrc, e.Player.PlayerDest, rl.Vector2{X: 0, Y:0}, 0, rl.White)
	rl.DrawTexturePro(
		e.Player.Sprite,
		e.Player.PlayerSrc,
		// e.Payer.Dest ne marche pas alors que c la mm chose ????
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 48, 48),
		rl.NewVector2(e.Player.PlayerDest.Width, e.Player.PlayerDest.Height),
		0,
		rl.White,
	)
	rl.EndMode2D()

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

func (e *Engine) RenderDialogMobs(m entity.Mobs, sentence string) {
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

	// Texte "GAME OVER" en rouge avec contour blanc
	title := "GAME OVER"
	titleWidth := int32(rl.MeasureText(title, 80))
	screenWidth := int32(rl.GetScreenWidth()) // Utiliser int32 ici
	xPos := (screenWidth - titleWidth) / 2    // Pas besoin de conversion

	// Dessin du texte
	rl.DrawText(title, xPos, 280, 80, rl.Red)
	rl.DrawText(title, xPos+2, 282, 80, rl.White)

	// Instructions en dessous
	instructions1 := "Press ENTER to restart"
	instructions2 := "Press ESC to leave"
	instruction1Width := int32(rl.MeasureText(instructions1, 32))
	instruction2Width := int32(rl.MeasureText(instructions2, 32))

	rl.DrawText(instructions1, (screenWidth-instruction1Width)/2, 600, 32, rl.White)
	rl.DrawText(instructions2, (screenWidth-instruction2Width)/2, 640, 32, rl.White)
}

func (e *Engine) RenderHealthBar() {

    screenHeight := int32(rl.GetScreenHeight())
    barWidth := int32(250)  
    barHeight := int32(20)  
    barX := int32(20)    
    barY := screenHeight - barHeight - 20 
     healthRatio := float32(e.Player.Health) / float32(e.Player.MaxHealth)
    if healthRatio > 1 {
        healthRatio = 1
    } else if healthRatio < 0 {
        healthRatio = 0
    }
    healthBarWidth := int32(float32(barWidth) * healthRatio)
    rl.DrawRectangle(barX, barY, barWidth, barHeight, rl.Gray)
    rl.DrawRectangle(barX, barY, healthBarWidth, barHeight, rl.Red)
}


func (e *Engine) UpdateAndRenderEndurance() {
	
    e.Player.UpdateEndurance()
    if rl.IsKeyPressed(rl.KeyEnter) {
        if e.Player.Endurance >= e.Player.MaxEndurance {
            fmt.Println("Action effectuée !")
            e.Player.Endurance = 0
        }
    }
}

func (e *Engine) RenderEnduranceBar() {

    screenWidth := int32(rl.GetScreenWidth())
    screenHeight := int32(rl.GetScreenHeight())
    barWidth := int32(150)
    barHeight := int32(20)
    barX := screenWidth - barWidth - 20
    barY := screenHeight - barHeight - 20
    enduranceRatio := float32(e.Player.Endurance) / float32(e.Player.MaxEndurance)
    if enduranceRatio > 1 {
        enduranceRatio = 1
    } else if enduranceRatio < 0 {
        enduranceRatio = 0
    }
    enduranceBarWidth := int32(float32(barWidth) * enduranceRatio)
    rl.DrawRectangle(barX, barY, barWidth, barHeight, rl.Gray)
    rl.DrawRectangle(barX, barY, enduranceBarWidth, barHeight, rl.Yellow)
	
}

func (e *Engine) UpdateAndRenderShield() {
	
    e.Player.UpdateShield()
    if rl.IsKeyPressed(rl.KeyQ) {
        if e.Player.Shield >= e.Player.MaxShield {
            fmt.Println("Action effectuée !")
            e.Player.Shield = 0
        }
    }
}

func (e *Engine) RenderShieldBar() {
    screenHeight := int32(rl.GetScreenHeight())
    barWidth := int32(100)
    barHeight := int32(20)
    barX := int32(20)
    barY := screenHeight - barHeight - 60 
    shieldRatio := float32(e.Player.Shield) / float32(e.Player.MaxShield)
    if shieldRatio > 1 {
        shieldRatio = 1
    } else if shieldRatio < 0 {
        shieldRatio = 0
    }
    shieldBarWidth := int32(float32(barWidth) * shieldRatio)
    rl.DrawRectangle(barX, barY, barWidth, barHeight, rl.Gray)
    rl.DrawRectangle(barX, barY, shieldBarWidth, barHeight, rl.Blue)

	
}



