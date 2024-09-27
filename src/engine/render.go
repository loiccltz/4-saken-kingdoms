package engine

import (
	"fmt"
	"main/src/building"
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"  // Importation de Raylib pour les fonctions de rendu graphique
)

// Rendering : Fonction de base pour le rendu, nettoie l'écran avec une couleur d'arrière-plan
func (e *Engine) Rendering() {
	rl.ClearBackground(rl.Blue) // Efface l'arrière-plan avec une couleur bleue
}

// HomeRendering : Rendu du menu principal
func (e *Engine) HomeRendering() {
	image := rl.LoadImage("4SKMENUENTRE-Photoroom.png") // Charge l'image du menu
	texture := rl.LoadTextureFromImage(image) // Crée une texture à partir de l'image
	rl.DrawTexture(texture, 0, 0, rl.White) // Dessine la texture en haut à gauche
}

// InGameRendering : Rendu du jeu en cours
func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)
	rl.BeginMode2D(e.Camera)
	e.RenderMap()
	e.RenderMobs()
	
	e.RenderPlayer()
	e.RenderShoot()
	e.RenderWolf()
	
	rl.EndMode2D()
	
	
	e.RenderMobsHealthBar()
	e.RenderHealthBar()
	e.RenderEnduranceBar()
	e.RenderShieldBar()

	e.RenderTower()
	e.UpdateAnimation()
	e.UpdateAndRenderShield()
	e.RenderSeller()


	e.RenderSeller()

	e.RenderWolf()
	e.RenderCrabe()
	e.RenderDragon()
	e.RenderGriffon()
	e.RenderPlayer()
	rl.EndMode2D()
	e.RenderPlayer()

	// Affichage des commandes à l'écran
	rl.DrawText("[Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Pause", 340)/4, int32(rl.GetScreenHeight())/2-320, 20, rl.Black)

	rl.DrawText("[A] AND [E] to navigate in the inventory", int32(rl.GetScreenWidth())/2-rl.MeasureText("[A] AND [E] to navigate in the inventory", 470)/4, int32(rl.GetScreenHeight())/2-280, 20, rl.Black)
	rl.DrawText("[I] to Inventory", int32(rl.GetScreenWidth())/2-rl.MeasureText("[I] to Inventory", 320)/4, int32(rl.GetScreenHeight())/2-280, 20, rl.Black)
	rl.DrawText("[Enter] to Attack", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Enter] to Attack", 280)/4, int32(rl.GetScreenHeight())/2-240, 20, rl.Black)
}

// InventoryRendering : Rendu de l'inventaire
func (e *Engine) InventoryRendering() {
	screenWidth := int32(rl.GetScreenWidth()) // Largeur de l'écran

	title := "Inventaire" // Titre de l'inventaire
	titleFontSize := int32(80) // Taille de la police du titre
	titleWidth := int32(rl.MeasureText(title, int32(titleFontSize))) // Largeur du titre
	titleXPos := (screenWidth - titleWidth) / 2 // Position horizontale centrée
	titleYPos := int32(50) // Position verticale du titre

	// Dessin du titre de l'inventaire avec un effet d'ombre
	rl.DrawText(title, titleXPos-2, titleYPos-2, titleFontSize, rl.Blue)
	rl.DrawText(title, titleXPos+2, titleYPos+2, titleFontSize, rl.Blue)
	rl.DrawText(title, titleXPos, titleYPos, titleFontSize, rl.White)

	itemSize := int32(80) // Taille des éléments de l'inventaire
	itemSpacing := int32(10) // Espacement entre les éléments

	// Dimensions et positions de la zone d'inventaire
	inventoryWidth := int32(7)*(itemSize+itemSpacing) + itemSpacing
	inventoryHeight := int32(120)
	inventoryXPos := (screenWidth - inventoryWidth) / 2
	inventoryYPos := titleYPos + titleFontSize + 50

	rl.DrawRectangle(inventoryXPos, inventoryYPos, inventoryWidth, inventoryHeight, rl.Gray) // Dessin de la zone d'inventaire

	// Gestion de la sélection des éléments de l'inventaire
	if rl.IsKeyPressed(rl.KeyQ) {
		e.selectedIndex-- // Décrémentation de l'index sélectionné
		if e.selectedIndex < 0 {
			e.selectedIndex = 6 // Réinitialisation si l'index est inférieur à 0
		}
	} else if rl.IsKeyPressed(rl.KeyE) {
		e.selectedIndex++ // Incrémentation de l'index sélectionné
		if e.selectedIndex > 6 {
			e.selectedIndex = 0 // Réinitialisation si l'index est supérieur à 6
		}
	}

	// Utilisation de l'élément sélectionné
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.UseSelectedItem()
	}

	// Rendu de l'élément sélectionné
	if len(e.Player.Inventory) > 0 {
		item := e.Player.Inventory[0] // Prend le premier élément de l'inventaire
		rl.DrawTexture(item.Sprite, 50, 50, rl.White) // Dessin de l'élément
		rl.DrawText(fmt.Sprintf("x%d", item.Quantity), 50, 110, 20, rl.White) // Affichage de la quantité
	}

	// Rendu des éléments de l'inventaire
	for i := 0; i < 7; i++ {
		itemXPos := inventoryXPos + int32(i)*(itemSize + itemSpacing) + itemSpacing
		itemYPos := inventoryYPos + (inventoryHeight-itemSize)/2

		if i == e.selectedIndex {
			rl.DrawRectangle(itemXPos-5, itemYPos-5, itemSize+10, itemSize+10, rl.White) // Mise en surbrillance de l'élément sélectionné
		}

		rl.DrawRectangle(itemXPos, itemYPos, itemSize, itemSize, rl.Black) // Dessin de la case de l'élément

		if i < len(e.Player.Inventory) {
			item := e.Player.Inventory[i]
			rl.DrawText(item.Name, itemXPos+10, itemYPos+40, 20, rl.White) // Affichage du nom de l'élément
		}
	}

	// Retour au jeu si échap ou 'I' est pressé
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyI) {
		e.StateEngine = INGAME
	}
}

// PauseRendering : Rendu du menu de pause
func (e *Engine) PauseRendering() {
	image := rl.LoadImage("4SKPAUSEMENU.png") // Charge l'image du menu de pause
	texture := rl.LoadTextureFromImage(image) // Crée une texture à partir de l'image
	rl.DrawTexture(texture, 0, 0, rl.White) // Dessine la texture
}

// SellerRendering : Rendu du menu du vendeur
func (e *Engine) SellerRendering() {
	rl.ClearBackground(rl.Beige) // Efface l'arrière-plan avec une couleur beige

	// Affichage du titre du menu du vendeur
	rl.DrawText("MenuSeller", int32(rl.GetScreenWidth())/2-rl.MeasureText("MesnuSeller", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[R] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[R] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.White)
	
	// Affichage des éléments à vendre
	rl.DrawText("ITEM 1", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 1", 20)/2, int32(rl.GetScreenHeight())/2+200, 20, rl.Black)
	rl.DrawText("ITEM 2", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 2", 20)/2, int32(rl.GetScreenHeight())/2+250, 20, rl.Black)
	rl.DrawText("ITEM 3", int32(rl.GetScreenWidth())/2-rl.MeasureText("ITEM 3", 20)/2, int32(rl.GetScreenHeight())/2+300, 20, rl.Black)
}

// RenderPlayer : Rendu du joueur
func (e *Engine) RenderPlayer(){
	rl.BeginMode2D(e.Camera) // Début du mode 2D
	rl.DrawTexturePro(
		e.Player.Sprite, // Sprite du joueur
		e.Player.PlayerSrc, // Source du sprite
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 48, 48),
		rl.NewVector2(e.Player.PlayerDest.Width, e.Player.PlayerDest.Height),
		0,
		rl.White,
	)
	rl.EndMode2D()
}
// RenderWolf : Rendu des loups
func (e *Engine) RenderWolf() {
	rl.BeginMode2D(e.Camera)
		rl.DrawTexturePro(
			e.Monsters[0].Sprite,
			e.Monsters[0].MonsterSrc,
			rl.NewRectangle(e.Monsters[0].Position.X, e.Monsters[0].Position.Y, 150, 150), // Position et dimensions
			rl.NewVector2(e.Monsters[0].MonsterDest.Width, e.Monsters[0].MonsterDest.Height), // Point d'origine pour la rotation
			0,
			rl.White,// Couleur
		)
	rl.EndMode2D() // Fin du mode 2D
}

// RenderGriffon : Rendu des griffons
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
// RenderCrabe : Rendu des crabes
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
// RenderDragon : Rendu des dragons
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

// RenderShoot : Rendu des tirs
func (e *Engine) RenderShoot() {
	for _, Shoot := range e.Shoot {
		rl.DrawTexturePro(
			Shoot.Sprite,
			rl.NewRectangle(0, 0, 10, 10),
			rl.NewRectangle(Shoot.Position.X, Shoot.Position.Y, 20, 20),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
}

// RenderMobs : Rendu des mobs
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

// RenderSeller : Rendu du vendeur
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

// RenderTower : Rendu de la tour
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
// RenderDialog : Rendu d'un dialogue d'un monstre
func (e *Engine) RenderDialog(m entity.Monster, sentence string) {
    rl.BeginMode2D(e.Camera) // Commence le mode 2D avec la caméra actuelle
    rl.DrawText(
        sentence, // Le texte à afficher
        int32(m.Position.X), // Position horizontale du monstre
        int32(m.Position.Y)+50, // Position verticale (50 pixels au-dessus du monstre)
        10, // Taille de la police
        rl.RayWhite, // Couleur du texte
    )
    rl.EndMode2D() // Termine le mode 2D
}

// RenderDialogMobs : Rendu d'un dialogue pour les mobs
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

// RenderExplanation : Rendu d'une explication pour une tour
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

// RenderExplanationShop : Rendu d'une explication pour le vendeur
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
func (e *Engine) RenderExplanationPnj(pnj entity.Pnj, sentence string) {
	rl.BeginMode2D(e.Camera)
	rl.DrawText(
		sentence,
		int32(pnj.Position.X),
		int32(pnj.Position.Y)+50,
		10,
		rl.Black,


	)
	rl.EndMode2D()
}

// RenderExplanationPnjCypher : Rendu d'une explication pour un PNJ spécifique
func (e *Engine) RenderExplanationPnjCypher(m entity.Pnj, sentence string) {

	rl.BeginMode2D(e.Camera)
	rl.DrawText(
		sentence,
		int32(m.Position.X)-50,
		int32(m.Position.Y)+50,
		10,
		rl.Black,
	)
	rl.EndMode2D()
}

// GameOverRendering : Rendu de l'écran de game over
func (e *Engine) GameOverRendering() {
    rl.ClearBackground(rl.Black) // Efface l'écran avec une couleur noire

    title := "GAME OVER" // Titre à afficher
    titleWidth := int32(rl.MeasureText(title, 80)) // Largeur du titre
    screenWidth := int32(rl.GetScreenWidth()) // Largeur de l'écran
    xPos := (screenWidth - titleWidth) / 2 // Position centrée

    // Dessin du titre avec un effet d'ombre
    rl.DrawText(title, xPos, 280, 80, rl.Red)
    rl.DrawText(title, xPos+2, 282, 80, rl.White)

    // Instructions pour redémarrer ou quitter
    instructions1 := "Press ENTER to restart"
    instructions2 := "Press ESC to leave"
    instruction1Width := int32(rl.MeasureText(instructions1, 32))
    instruction2Width := int32(rl.MeasureText(instructions2, 32))

    // Dessin des instructions
    rl.DrawText(instructions1, (screenWidth-instruction1Width)/2, 600, 32, rl.White)
    rl.DrawText(instructions2, (screenWidth-instruction2Width)/2, 640, 32, rl.White)
}

var blinkTimer float32 = 0 // Timer pour le clignotement
var blinkState bool = false // État du clignotement

// RenderHealthBar : Rendu de la barre de vie du joueur
func (e *Engine) RenderHealthBar() {
    screenHeight := int32(rl.GetScreenHeight())
    screenWidth := int32(rl.GetScreenWidth())

    // Dimensions et position de la barre de vie
    barWidth := int32(250)  // Largeur de la barre
    barHeight := int32(20)  // Hauteur de la barre
    barX := int32(20)       // Position X
    barY := screenHeight - barHeight - 20 // Position Y

    // Ratio de vie
    healthRatio := float32(e.Player.Health) / float32(e.Player.MaxHealth) 
    if healthRatio > 1 {
        healthRatio = 1
    } else if healthRatio < 0 {
        healthRatio = 0
    }
    healthBarWidth := int32(float32(barWidth) * healthRatio) // Largeur de la barre de vie

    // Dessin de la barre de fond et de la barre de vie
    rl.DrawRectangle(barX, barY, barWidth, barHeight, rl.Gray)
    rl.DrawRectangle(barX, barY, healthBarWidth, barHeight, rl.Red)

    // Clignotement si la vie est critique
    if healthRatio <= 0.2 {
        blinkTimer += rl.GetFrameTime() // Incrémente le timer

        if blinkTimer >= 0.5 { // Changement d'état toutes les 0.5 secondes
            blinkState = !blinkState
            blinkTimer = 0
        }

        // Dessin du clignotement aux bords de l'écran
        if blinkState {
            thickness := int32(10) 
            rl.DrawRectangle(0, 0, screenWidth, thickness, rl.Red)                       
            rl.DrawRectangle(0, screenHeight-thickness, screenWidth, thickness, rl.Red)   
            rl.DrawRectangle(0, 0, thickness, screenHeight, rl.Red)                       
            rl.DrawRectangle(screenWidth-thickness, 0, thickness, screenHeight, rl.Red)   
        }
    }
}

// RenderEnduranceBar : Rendu de la barre d'endurance
func (e *Engine) RenderEnduranceBar() {
    screenWidth := int32(rl.GetScreenWidth())
    screenHeight := int32(rl.GetScreenHeight())
    barWidth := int32(150) // Largeur de la barre d'endurance
    barHeight := int32(20) // Hauteur de la barre
    barX := screenWidth - barWidth - 20 // Position X
    barY := screenHeight - barHeight - 20 // Position Y

    // Ratio d'endurance
    enduranceRatio := float32(e.Player.Endurance) / float32(e.Player.MaxEndurance)
    if enduranceRatio > 1 {
        enduranceRatio = 1
    } else if enduranceRatio < 0 {
        enduranceRatio = 0
    }
    enduranceBarWidth := int32(float32(barWidth) * enduranceRatio) // Largeur de la barre d'endurance

    // Dessin de la barre de fond et de la barre d'endurance
    rl.DrawRectangle(barX, barY, barWidth, barHeight, rl.Gray)
    rl.DrawRectangle(barX, barY, enduranceBarWidth, barHeight, rl.Yellow)
}

// UpdateAndRenderShield : Mise à jour et rendu de la barre de bouclier
func (e *Engine) UpdateAndRenderShield() {
    e.Player.UpdateShield() // Mise à jour de la barre de bouclier
    e.RenderShieldBar() // Rendu de la barre
}

// ApplyDamageToPlayer : Applique les dégâts au joueur en tenant compte du bouclier
func (e *Engine) ApplyDamageToPlayer(damage int) {
    if e.Player.Shield > 0 { // Si le joueur a un bouclier
        if damage <= e.Player.Shield {
            e.Player.Shield -= damage // Réduit les dégâts du bouclier
        } else {
            remainingDamage := damage - e.Player.Shield
            e.Player.Shield = 0 // Le bouclier est épuisé
            e.Player.Health -= remainingDamage // Applique les dégâts restants à la vie
        }
    } else {
        e.Player.Health -= damage // Applique les dégâts directement à la vie
    }

    // S'assure que la vie ne devient pas négative
    if e.Player.Health < 0 {
        e.Player.Health = 0
    }
}

// RenderShieldBar : Rendu de la barre de bouclier
func (e *Engine) RenderShieldBar() {
    screenHeight := int32(rl.GetScreenHeight())
    barWidth := int32(100) // Largeur de la barre de bouclier
    barHeight := int32(20) // Hauteur de la barre
    barX := int32(20) // Position X
    barY := screenHeight - barHeight - 60 // Position Y

    // Ratio du bouclier
    shieldRatio := float32(e.Player.Shield) / float32(e.Player.MaxShield)
    if shieldRatio > 1 {
        shieldRatio = 1
    } else if shieldRatio < 0 {
        shieldRatio = 0
    }
    shieldBarWidth := int32(float32(barWidth) * shieldRatio) // Largeur de la barre de bouclier

    // Dessin de la barre de fond
    rl.DrawRectangle(barX, barY, barWidth, barHeight, rl.Gray)

    // Changer la couleur en fonction du niveau du bouclier
    shieldColor := rl.Blue
    if shieldRatio < 0.5 {
        shieldColor = rl.Yellow // Jaune si le bouclier est faible
        if shieldRatio < 0.2 {
            shieldColor = rl.Red // Rouge si le bouclier est critique
        }
    }

    // Dessin de la barre de bouclier
    rl.DrawRectangle(barX, barY, shieldBarWidth, barHeight, shieldColor)

}

func (e *Engine) RenderMobsHealthBar() {
    barWidth := int32(60)  // Largeur de la barre de vie
    barHeight := int32(9)  // Hauteur de la barre de vie

    // Boucle à travers tous les mobs
    for _, mob := range e.Mobs {
        if mob.IsAlive {  //dessiner la barre que si le mob est vivant
            // Utiliser la fonction GetWorldToScreen pour obtenir la position de la barre de vie sur l'écran
            screenPosition := rl.GetWorldToScreen2D(mob.Position, e.Camera)

            barX := int32(screenPosition.X) + 125  // Ajustement horizontal (décalage à droite)
            barY := int32(screenPosition.Y) + int32(mob.Sprite.Height)/2 + 50 // Ajustement vertical (décalage vers le bas)

            // Calculer le ratio de santé du mob
            healthRatio := float32(mob.Health) / float32(mob.MaxHealth)
            if healthRatio > 1 {
                healthRatio = 1
            } else if healthRatio < 0 {
                healthRatio = 0
            }
            healthBarWidth := int32(float32(barWidth) * healthRatio)

            // Dessiner la barre de vie en arrière-plan (barre grise)
            rl.DrawRectangle(barX, barY, barWidth, barHeight, rl.Gray)

            // Dessiner la barre de vie en fonction de la santé restante (barre blanche)
            rl.DrawRectangle(barX, barY, healthBarWidth, barHeight, rl.White)
        }
    }
}