package engine

import (
	"fmt"
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) HomeLogic() {

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	//Menus
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
		rl.StopMusicStream(e.Music)

	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
}

func (e *Engine) SettingsLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyB) {
		e.StateMenu = HOME
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) InGameLogic() {
	// Mouvement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		e.Player.Position.Y -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		e.Player.Position.Y += e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		e.Player.Position.X -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		e.Player.Position.X += e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyV){
		
	}
	

	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}

	e.CheckCollisions()
	
	rl.DrawTextPro()

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-07-Simon_s-In-There-Somewhere.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
}


func (e *Engine) CheckCollisionsWithObjects()  bool{
    playerRect := rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 40, 40)
	// * 2 - 16
    for _, obj := range e.Objects {
		objectRect := rl.NewRectangle(obj.X * 2 - 16, obj.Y * 2 - 16, obj.Width, obj.Height)
        if rl.CheckCollisionRecs(playerRect, objectRect) {
			fmt.Print("coli")
			fmt.Println(objectRect)
			return true
		} 
    }
	return false
}	

func (e *Engine) CheckCollisions() {
	e.MonsterCollisions()
	e.VendorCollisions()
	e.CheckCollisionsWithObjects()
}



func (e *Engine) MonsterCollisions() {

	for _, monster := range e.Monsters {
		if monster.Position.X > e.Player.Position.X-20 &&
			monster.Position.X < e.Player.Position.X+20 &&
			monster.Position.Y > e.Player.Position.Y-20 &&
			monster.Position.Y < e.Player.Position.Y+20 {

			if monster.Name == "claude" {
				e.NormalTalk(monster, "BANGER!!!")
				if rl.IsKeyPressed(rl.KeyE) {
					// lance un combat
				}
			}
		} else {
			//...
		}
	}
}

func (e *Engine) VendorCollisions() {

	for _, vendor := range e.Vendor {
		if vendor.Position.X > e.Player.Position.X-20 &&
			vendor.Position.X < e.Player.Position.X+20 &&
			vendor.Position.Y > e.Player.Position.Y-20 &&
			vendor.Position.Y < e.Player.Position.Y+20 {

			if vendor.Name == "test" {
				e.VendorTalk(vendor, "Bondsdsdsds")
				if rl.IsKeyPressed(rl.KeyE) {
					//lancer un combat ?
				}
			}
		} else {
			//...
		}
	}
}

func (e *Engine) NormalTalk(m entity.Monster,sentence string) {
	e.RenderDialog(m, sentence)
}
func (e *Engine) VendorTalk(v entity.Vendor,sentence string) {
	e.RenderDialogVendor(v, sentence)
}

func (e *Engine) AutoAttack(){
	if rl.IsKeyPressed(rl.KeyE) {
		e.StateEngine = FIGHT
	}
}

func (e *Engine) PauseLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyA) {
		e.StateMenu = HOME
		rl.StopMusicStream(e.Music)
	}

	//Musique
	rl.UpdateMusicStream(e.Music)
}
