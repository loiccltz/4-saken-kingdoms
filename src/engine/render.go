package engine

import (
	"main/src/building"
	"fmt"
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

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)
	rl.BeginMode2D(e.Camera)
	e.RenderMap()
	e.RenderMobs()
	e.RenderShoot()
	e.RenderWolf()
	e.RenderPlayer()

	rl.EndMode2D()
	
	rl.DrawText("[Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Pause", 340)/4, int32(rl.GetScreenHeight())/2-320, 20, rl.Black)
	rl.DrawText("[I] to Inventory", int32(rl.GetScreenWidth())/2-rl.MeasureText("[I] to Inventory", 420)/4, int32(rl.GetScreenHeight())/2-280, 20, rl.Black)
	
	e.RenderHealthBar()
	e.RenderEnduranceBar()

	e.RenderTower()
	e.UpdateAnimation()
	e.RenderSeller()
	e.RenderEnduranceBar()
	e.RenderHealthBar()
	e.RenderShieldBar()

	e.UpdateAndRenderShield()

	e.RenderSeller()

	e.RenderWolf()
	e.RenderCrabe()
	e.RenderDragon()
	e.RenderGriffon()
	e.RenderPlayer()
	rl.EndMode2D()
	rl.DrawText("[Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Pause", 340)/4, int32(rl.GetScreenHeight())/2-320, 20, rl.Black)
	rl.DrawText("[I] to Inventory", int32(rl.GetScreenWidth())/2-rl.MeasureText("[I] to Inventory", 420)/4, int32(rl.GetScreenHeight())/2-280, 20, rl.Black)

}

func (e *Engine) InventoryRendering() {
    screenWidth := int32(rl.GetScreenWidth())

    title := "Inventaire"
    titleFontSize := int32(80)
    titleWidth := int32(rl.MeasureText(title, int32(titleFontSize)))
    titleXPos := (screenWidth - titleWidth) / 2
    titleYPos := int32(50)

    rl.DrawText(title, titleXPos-2, titleYPos-2, titleFontSize, rl.Blue)
    rl.DrawText(title, titleXPos+2, titleYPos+2, titleFontSize, rl.Blue)
    rl.DrawText(title, titleXPos, titleYPos, titleFontSize, rl.White)

    itemSize := int32(80)
    itemSpacing := int32(10)

    inventoryWidth := int32(7)*(itemSize+itemSpacing) + itemSpacing
    inventoryHeight := int32(120)
    inventoryXPos := (screenWidth - inventoryWidth) / 2
    inventoryYPos := titleYPos + titleFontSize + 50

    rl.DrawRectangle(inventoryXPos, inventoryYPos, inventoryWidth, inventoryHeight, rl.Gray)

    if rl.IsKeyPressed(rl.KeyQ) {
        e.selectedIndex--  
        if e.selectedIndex < 0 {
            e.selectedIndex = 6
        }
    } else if rl.IsKeyPressed(rl.KeyE) {
        e.selectedIndex++
        if e.selectedIndex > 6 {
            e.selectedIndex = 0
        }
    }


	if rl.IsKeyPressed(rl.KeyEnter) {
		e.UseSelectedItem()
	}
	if len(e.Player.Inventory) > 0 {
		item := e.Player.Inventory[0]
		
	
		rl.DrawTexture(item.Sprite, 50, 50, rl.White)
	
		rl.DrawText(fmt.Sprintf("x%d", item.Quantity), 50, 110, 20, rl.White)
	}

    for i := 0; i < 7; i++ {
        itemXPos := inventoryXPos + int32(i)*(itemSize + itemSpacing) + itemSpacing
        itemYPos := inventoryYPos + (inventoryHeight-itemSize)/2

        if i == e.selectedIndex {
            rl.DrawRectangle(itemXPos-5, itemYPos-5, itemSize+10, itemSize+10, rl.White)
        }

		

        rl.DrawRectangle(itemXPos, itemYPos, itemSize, itemSize, rl.Black)

        if i < len(e.Player.Inventory) {
            item := e.Player.Inventory[i]
            rl.DrawText(item.Name, itemXPos+10, itemYPos+40, 20, rl.White)
        }
    }

    if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyI) {
        e.StateEngine = INGAME
    }
}



func (e *Engine) PauseRendering() {

	image := rl.LoadImage("4SKPAUSEMENU.png")
    texture := rl.LoadTextureFromImage(image)
	rl.DrawTexture(texture, 0, 0, rl.White)
}

func(e *Engine) SellerRendering() {
	rl.ClearBackground(rl.Beige)

	rl.DrawText("MenuSeller", int32(rl.GetScreenWidth())/2-rl.MeasureText("MesnuSeller", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[R] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[R] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.White)
	
	rl.DrawText("ITEM 1", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 1", 20)/2, int32(rl.GetScreenHeight())/2+200, 20, rl.Black)
	rl.DrawText("ITEM 2", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 2", 20)/2, int32(rl.GetScreenHeight())/2+250, 20, rl.Black)
	rl.DrawText("ITEM 3", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 3", 20)/2, int32(rl.GetScreenHeight())/2+300, 20, rl.Black)
	
}


func (e *Engine) RenderPlayer(){
	rl.BeginMode2D(e.Camera)

	rl.DrawTexturePro(
		e.Player.Sprite,
		e.Player.PlayerSrc,
	
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 48, 48),
		rl.NewVector2(e.Player.PlayerDest.Width, e.Player.PlayerDest.Height),
		0,
		rl.White,
	)
	rl.EndMode2D()
}

func (e *Engine) RenderWolf() {
	rl.BeginMode2D(e.Camera)
		rl.DrawTexturePro(
			e.Monsters[0].Sprite,
			e.Monsters[0].MonsterSrc,
			rl.NewRectangle(e.Monsters[0].Position.X, e.Monsters[0].Position.Y, 150, 150),
			rl.NewVector2(e.Monsters[0].MonsterDest.Width, e.Monsters[0].MonsterDest.Height),
			0,
			rl.White,
		)
	rl.EndMode2D()
}


func (e *Engine) RenderGriffon() {
	rl.BeginMode2D(e.Camera)
		rl.DrawTexturePro(
			e.Monsters[1].Sprite,
			e.Monsters[1].MonsterSrc,
			rl.NewRectangle(e.Monsters[1].Position.X, e.Monsters[1].Position.Y, 140, 140),
			rl.NewVector2(e.Monsters[1].MonsterDest.Width, e.Monsters[1].MonsterDest.Height),
			0,
			rl.White,
		)
	rl.EndMode2D()
}

func (e *Engine) RenderCrabe() {
	rl.BeginMode2D(e.Camera)
		rl.DrawTexturePro(
			e.Monsters[2].Sprite,
			e.Monsters[2].MonsterSrc,
			rl.NewRectangle(e.Monsters[2].Position.X, e.Monsters[2].Position.Y, 170, 170),
			rl.NewVector2(e.Monsters[2].MonsterDest.Width, e.Monsters[2].MonsterDest.Height),
			0,
			rl.White,
		)
	rl.EndMode2D()
}

func (e *Engine) RenderDragon() {
	rl.BeginMode2D(e.Camera)
		rl.DrawTexturePro(
			e.Monsters[3].Sprite,
			e.Monsters[3].MonsterSrc,
			rl.NewRectangle(e.Monsters[3].Position.X, e.Monsters[3].Position.Y, 150, 150),
			rl.NewVector2(e.Monsters[3].MonsterDest.Width, e.Monsters[3].MonsterDest.Height),
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



func (e *Engine) RenderExplanationShop(m entity.Seller, sentence string) {
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
func (e *Engine) RenderExplanationPnj(m entity.Pnj, sentence string) {
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


	title := "GAME OVER"
	titleWidth := int32(rl.MeasureText(title, 80))
	screenWidth := int32(rl.GetScreenWidth()) 
	xPos := (screenWidth - titleWidth) / 2    


	rl.DrawText(title, xPos, 280, 80, rl.Red)
	rl.DrawText(title, xPos+2, 282, 80, rl.White)

	instructions1 := "Press ENTER to restart"
	instructions2 := "Press ESC to leave"
	instruction1Width := int32(rl.MeasureText(instructions1, 32))
	instruction2Width := int32(rl.MeasureText(instructions2, 32))

	rl.DrawText(instructions1, (screenWidth-instruction1Width)/2, 600, 32, rl.White)
	rl.DrawText(instructions2, (screenWidth-instruction2Width)/2, 640, 32, rl.White)
}


var blinkTimer float32 = 0
var blinkState bool = false

func (e *Engine) RenderHealthBar() {
    screenHeight := int32(rl.GetScreenHeight())
    screenWidth := int32(rl.GetScreenWidth())

    // Dimensions et positions de la barre de vie
    barWidth := int32(250)  // Largeur de la barre de vie
    barHeight := int32(20)  // Hauteur de la barre de vie
    barX := int32(20)       // 20 pixels de marge depuis le bord gauche
    barY := screenHeight - barHeight - 20 // 20 pixels de marge depuis le bord bas


    healthRatio := float32(e.Player.Health) / float32(e.Player.MaxHealth) 
    if healthRatio > 1 {
        healthRatio = 1
    } else if healthRatio < 0 {
        healthRatio = 0
    }
    healthBarWidth := int32(float32(barWidth) * healthRatio)


    rl.DrawRectangle(barX, barY, barWidth, barHeight, rl.Gray)


    rl.DrawRectangle(barX, barY, healthBarWidth, barHeight, rl.Red)


    if healthRatio <= 0.2 {

        blinkTimer += rl.GetFrameTime() 


        if blinkTimer >= 0.5 {
            blinkState = !blinkState
            blinkTimer = 0
        }


        if blinkState {
            thickness := int32(10) 
            rl.DrawRectangle(0, 0, screenWidth, thickness, rl.Red)                       
            rl.DrawRectangle(0, screenHeight-thickness, screenWidth, thickness, rl.Red)   
            rl.DrawRectangle(0, 0, thickness, screenHeight, rl.Red)                       
            rl.DrawRectangle(screenWidth-thickness, 0, thickness, screenHeight, rl.Red)   
        }
    }
}

func (e *Engine) RenderEnduranceBar() {

    screenWidth := int32(rl.GetScreenWidth())
    screenHeight := int32(rl.GetScreenHeight())
    barWidth := int32(150) // Largeur de la barre d'endurance
    barHeight := int32(20) // Hauteur de la barre d'endurance
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


    e.RenderShieldBar()
}

func (e *Engine) ApplyDamageToPlayer(damage int) {
    if e.Player.Shield > 0 {
        if damage <= e.Player.Shield {
            e.Player.Shield -= damage
        } else {
            remainingDamage := damage - e.Player.Shield
            e.Player.Shield = 0
            e.Player.Health -= remainingDamage
        }
    } else {
        e.Player.Health -= damage
    }

    if e.Player.Health < 0 {
        e.Player.Health = 0
    }
}

func (e *Engine) RenderShieldBar() {
    screenHeight := int32(rl.GetScreenHeight())
    barWidth := int32(100) // Largeur de la barre de bouclier
    barHeight := int32(20) // Hauteur de la barre de bouclier
    barX := int32(20)
    barY := screenHeight - barHeight - 60 // Position modifiée pour la barre de bouclier
    shieldRatio := float32(e.Player.Shield) / float32(e.Player.MaxShield)
    if shieldRatio > 1 {
        shieldRatio = 1
    } else if shieldRatio < 0 {
        shieldRatio = 0
    }
    shieldBarWidth := int32(float32(barWidth) * shieldRatio)

    
    rl.DrawRectangle(barX, barY, barWidth, barHeight, rl.Gray)

    // Changer la couleur en fonction du niveau du bouclier
    shieldColor := rl.Blue
    if shieldRatio < 0.5 {
        shieldColor = rl.Yellow 
        if shieldRatio < 0.2 {
            shieldColor = rl.Red
        }
    }

    rl.DrawRectangle(barX, barY, shieldBarWidth, barHeight, shieldColor)
}
