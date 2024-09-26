package engine

import (
	"flag"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (engine *Engine) Run() {

	rl.SetTargetFPS(120)

    rl.ToggleFullscreen()

    if rl.IsWindowFullscreen() {
        ScreenWidth = rl.GetScreenWidth()
        ScreenHeight = rl.GetScreenHeight()
    }


	showFPS := flag.Bool("f", false, "Affiche les FPS")
	// fullscreen := flag.Bool("fullscreen", false, "Lance le jeu en plein ecran")
	showCoord := flag.Bool("coord", false, "Affiche les coordonnées du joueur")
	// fullscreenFlag := flag.Bool("p", false, "Lancer le jeu en mode plein écran")
	flag.Parse()

        // Si l'option -f est utilisée, afficher les FPS
   
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

	
		/*if *fullscreenFlag {
			rl.ToggleFullscreen()
		}
			if rl.IsWindowFullscreen() {
				ScreenWidth = rl.GetScreenWidth() 
				ScreenHeight = rl.GetScreenHeight() 
			}*/

		if *showFPS {
            fps := rl.GetFPS()
            rl.DrawText(fmt.Sprintf("FPS: %d", fps), 10, 10, 20, rl.DarkGray)
        }

		if *showCoord {
			// obliger de mettre 1 apres la virgule les coord sont en float
			coord := fmt.Sprintf("X: %.1f, Y: %.1f", engine.Player.Position.X, engine.Player.Position.Y)
			rl.DrawText(coord, 20, 80, 20, rl.DarkGray)
		}
		/*if *fullscreen {
			rl.ToggleFullscreen()
		}*/

			rl.EndDrawing()

	}
}
