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
	rl.ClearBackground(rl.Blue)

	rl.DrawText("Home Menu", int32(rl.GetScreenWidth())/2-rl.MeasureText("Home Menu", 25)/2, int32(rl.GetScreenHeight())/4-150, 40, rl.RayWhite)
    rl.DrawText("[Enter] to Play", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Enter] to Play", 0)/2, int32(rl.GetScreenHeight())/4, 20, rl.RayWhite)
    rl.DrawText("[Esc] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 0)/2, int32(rl.GetScreenHeight())/4+100, 20, rl.RayWhite)
    rl.DrawText("Monster lvl 1", int32(rl.GetScreenWidth())/2-rl.MeasureText("Monster lvl 1 ", 199)/2, int32(rl.GetScreenHeight())/4, 20, rl.Yellow)
    rl.DrawText("Monster lvl 2", int32(rl.GetScreenWidth())/2-rl.MeasureText("Monster lvl 2", 200)/2, int32(rl.GetScreenHeight())/4+100, 20, rl.Orange)
    rl.DrawText("Monster lvl 3", int32(rl.GetScreenWidth())/2-rl.MeasureText("Monster lvl 3", 200)/2, int32(rl.GetScreenHeight())/4+200, 20, rl.Red)
    rl.DrawText("Monster lvl 4", int32(rl.GetScreenWidth())/2-rl.MeasureText("Monster lvl 4", 200)/2, int32(rl.GetScreenHeight())/4+300, 20, rl.Black)
    rl.DrawText("HISTOIRE", int32(rl.GetScreenWidth())/1-rl.MeasureText("HISTOIRE", 100)/2, int32(rl.GetScreenHeight())/4, 20, rl.RayWhite)
    rl.DrawText("Z move forward", int32(rl.GetScreenWidth())/2-rl.MeasureText("Z move forward", -50)/2, int32(rl.GetScreenHeight())/4+250, 20, rl.RayWhite)
    rl.DrawText("Q move forward", int32(rl.GetScreenWidth())/2-rl.MeasureText("Q move forward", -50)/2, int32(rl.GetScreenHeight())/4+270, 20, rl.RayWhite)
    rl.DrawText("S move forward", int32(rl.GetScreenWidth())/2-rl.MeasureText("S move forward", -50)/2, int32(rl.GetScreenHeight())/4+290, 20, rl.RayWhite)
    rl.DrawText("D move forward", int32(rl.GetScreenWidth())/2-rl.MeasureText("D move forward", -50)/2, int32(rl.GetScreenHeight())/4+310, 20, rl.RayWhite)
    rl.DrawText("4 SAKEN KINGDOM", int32(rl.GetScreenWidth())/2-rl.MeasureText("4 SAKEN KINGDOM", 0)/2, int32(rl.GetScreenHeight())/4-200, 20, rl.Black)
}


// AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
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
}

// AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
func (e *Engine) PauseRendering() {
	rl.ClearBackground(rl.Red)
	rl.DrawText("Paused", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paused", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[Esc] to Resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[f] to Exit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[f] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)
	rl.EndDrawing()
}
func (e *Engine) InventoryRendering(){
	rl.ClearBackground(rl.Gray)
	rl.DrawText("Inventaire", int32(rl.GetScreenWidth())/2-rl.MeasureText("Inventaire", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
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
