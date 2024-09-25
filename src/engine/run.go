package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (engine *Engine) Run() {
	rl.SetTargetFPS(60)
<<<<<<< Updated upstream
=======
	// si l'option -f est utilisée, affiche les FPS
	showFPS := flag.Bool("f", false, "Affiche les FPS")
	fullscreen := flag.Bool("fullscreen", false, "Lance le jeu en plein ecran")
	flag.Parse()


   
>>>>>>> Stashed changes

	for engine.IsRunning {

		rl.BeginDrawing()

		switch engine.StateMenu {
			case HOME:
				engine.HomeRendering()
				engine.HomeLogic()

			case SETTINGS:
				engine.SettingsLogic()

			case PLAY:
				switch engine.StateEngine {
				case INGAME:
					engine.InGameRendering()
					engine.InGameLogic()
					if engine.Player.Health <= 0 {
						engine.StateEngine = GAMEOVER
					}
				
				case INVENTORY:
					engine.InventoryRendering()
					engine.InventoryLogic()
					//engine.ComeBackLogic()
		
				case MENUSELLER:
					engine.SellerRendering()
					engine.SellerLogic()
				
				case PAUSE:
					engine.PauseRendering()
					engine.PauseLogic()
				
				case GAMEOVER:
					engine.GameOverRendering()
					engine.GameOverLogic()
			}
		}
<<<<<<< Updated upstream
=======

		if *showFPS {
            fps := rl.GetFPS()
            rl.DrawText(fmt.Sprintf("FPS: %d", fps), 10, 10, 20, rl.DarkGray)
        }
		if *fullscreen {
			rl.ToggleFullscreen()
		}
>>>>>>> Stashed changes
			rl.EndDrawing()
	}
}
