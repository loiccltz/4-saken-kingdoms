package engine

import (
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

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderPlayer()
	e.RenderVendor()

	rl.EndMode2D() // On finit le rendu camera

	// Ecriture fixe (car pas affectée par le mode camera)
	rl.DrawText("Playing", int32(rl.GetScreenWidth())/2-rl.MeasureText("Playing", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)
	rl.DrawRectangle(int32(e.Player.Position.X), int32(e.Player.Position.Y),150,150, rl.Beige);
}

func (e *Engine) PauseRendering() {
	rl.ClearBackground(rl.Red)

	rl.DrawText("Paused", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paused", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Q]/[A] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)

	rl.EndDrawing()
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

func (e *Engine) RnderAnimation(){
	
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

func (e *Engine) RenderVendor() {
	for _, vendor := range e.Vendor {
		rl.DrawTexturePro(
			vendor.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(vendor.Position.X, vendor.Position.Y, 150, 150),
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

func (e *Engine) RenderDialogVendor(v entity.Vendor, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(v.Position.X),
		int32(v.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}
