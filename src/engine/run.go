package engine

import (
	"flag" // Importation du package flag pour gérer les options de ligne de commande
	"fmt"  // Importation du package fmt pour la mise en forme des chaînes
	rl "github.com/gen2brain/raylib-go/raylib" // Importation de Raylib pour les fonctionnalités graphiques
)

// Run : Méthode principale pour exécuter le moteur de jeu
func (engine *Engine) Run() {
	rl.SetTargetFPS(120) // Définit le nombre d'images par seconde ciblé à 120

    rl.ToggleFullscreen() // Bascule entre le mode plein écran et le mode fenêtré

    // Si le jeu est en plein écran, récupérer la largeur et la hauteur de l'écran
    if rl.IsWindowFullscreen() {
        ScreenWidth = rl.GetScreenWidth() // Récupère la largeur de l'écran
        ScreenHeight = rl.GetScreenHeight() // Récupère la hauteur de l'écran
    }

	// Déclaration des options de ligne de commande
	showFPS := flag.Bool("f", false, "Affiche les FPS") // Option pour afficher les FPS
	// fullscreen := flag.Bool("fullscreen", false, "Lance le jeu en plein écran") // Option commentée pour lancer le jeu en plein écran
	showCoord := flag.Bool("coord", false, "Affiche les coordonnées du joueur") // Option pour afficher les coordonnées du joueur
	// fullscreenFlag := flag.Bool("p", false, "Lancer le jeu en mode plein écran") // Option commentée pour le plein écran
	flag.Parse() // Analyse les options de ligne de commande

	// Boucle principale du moteur tant que le jeu est en cours d'exécution
	for engine.IsRunning {
		rl.BeginDrawing() // Commence le processus de dessin

		// Gestion des différents états du menu
		switch engine.StateMenu {
			case HOME: // État d'accueil
				engine.HomeRendering() // Rendu de l'écran d'accueil
				engine.HomeLogic() // Logique de l'écran d'accueil

			case SETTINGS: // État des paramètres
				engine.SettingsLogic() // Logique de l'écran des paramètres

			case PLAY: // État de jeu
				switch engine.StateEngine {
				case INGAME: // État de jeu actif
					engine.InGameRendering() // Rendu de l'écran de jeu
					engine.InGameLogic() // Logique de l'écran de jeu
					if engine.Player.Health <= 0 { // Vérifie si la vie du joueur est inférieure ou égale à 0
						engine.StateEngine = GAMEOVER // Change l'état en GAMEOVER
					}
				
				case INVENTORY: // État d'inventaire
					engine.InventoryRendering() // Rendu de l'écran d'inventaire
					engine.InventoryLogic() // Logique de l'inventaire
					// engine.ComeBackLogic() // Option commentée pour revenir en arrière (non utilisé)
		
				case MENUSELLER: // État de menu du vendeur
					engine.SellerRendering() // Rendu de l'écran du vendeur
					engine.SellerLogic() // Logique de l'écran du vendeur
				
				case PAUSE: // État de pause
					engine.PauseRendering() // Rendu de l'écran de pause
					engine.PauseLogic() // Logique de l'écran de pause
				
				case GAMEOVER: // État de fin de jeu
					engine.GameOverRendering() // Rendu de l'écran de game over
					engine.GameOverLogic() // Logique de l'écran de game over
			}
		}

		// Affichage des FPS si l'option -f est activée
		if *showFPS {
            fps := rl.GetFPS() // Récupère les FPS actuels
            rl.DrawText(fmt.Sprintf("FPS: %d", fps), 10, 10, 20, rl.DarkGray) // Affiche les FPS à l'écran
        }

		// Affichage des coordonnées du joueur si l'option -coord est activée
		if *showCoord {
			// Affiche les coordonnées en format float avec 1 chiffre après la virgule
			coord := fmt.Sprintf("X: %.1f, Y: %.1f", engine.Player.Position.X, engine.Player.Position.Y)
			rl.DrawText(coord, 20, 80, 20, rl.DarkGray) // Affiche les coordonnées à l'écran
		}

		rl.EndDrawing() // Termine le processus de dessin
	}
}
