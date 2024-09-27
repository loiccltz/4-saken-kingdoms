package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib" // Import du module raylib pour les fonctionnalités graphiques
)

// Définition de la structure Object représentant un objet avec des coordonnées, des dimensions et une rotation
type Object struct {
	X        float32 `json:"x"`        // Coordonnée X de l'objet
	Y        float32 `json:"y"`        // Coordonnée Y de l'objet
	Width    float32 `json:"width"`    // Largeur de l'objet
	Height   float32 `json:"height"`   // Hauteur de l'objet
	Rotation float32 `json:"rotation"` // Rotation de l'objet
}

// Structure pour représenter un "Ranger", similaire à un Object
type Ranger struct {
	X      float32 `json:"x"`      // Coordonnée X du ranger
	Y      float32 `json:"y"`      // Coordonnée Y du ranger
	Width  float32 `json:"width"`  // Largeur du ranger
	Height float32 `json:"height"` // Hauteur du ranger
}

// Représente un calque (layer) de la carte, contenant des données sur les tuiles, objets, et boss
type Layer struct {
	Data      []int    `json:"data"`    // Données des tuiles du calque
	Height    int      `json:"height"`  // Hauteur du calque (en nombre de tuiles)
	ID        int      `json:"id"`      // Identifiant unique du calque
	Name      string   `json:"name"`    // Nom du calque
	Opacity   float32  `json:"opacity"` // Opacité du calque
	Type      string   `json:"type"`    // Type de calque (par exemple, "tilelayer" ou "objectgroup")
	Visible   bool     `json:"visible"` // Indicateur de visibilité du calque
	Width     int      `json:"width"`   // Largeur du calque (en nombre de tuiles)
	X         int      `json:"x"`       // Position X de départ du calque
	Y         int      `json:"y"`       // Position Y de départ du calque
	Objects   []Object `json:"objects"` // Liste des objets sur le calque (s'il s'agit d'un calque d'objets)
}

// Définition d'un TileSet (jeu de tuiles), avec ses caractéristiques
type TileSet struct {
	Columns     int    `json:"columns"`     // Nombre de colonnes dans l'image du jeu de tuiles
	FirstGid    int    `json:"firstgid"`    // Premier identifiant global utilisé pour ces tuiles
	Image       string `json:"image"`       // Chemin de l'image du jeu de tuiles
	ImageHeight int    `json:"imageheight"` // Hauteur de l'image du jeu de tuiles
	ImageWidth  int    `json:"imagewidth"`  // Largeur de l'image du jeu de tuiles
	Margin      int    `json:"margin"`      // Marge entre les tuiles dans l'image
	Name        string `json:"name"`        // Nom du jeu de tuiles
	Spacing     int    `json:"spacing"`     // Espacement entre les tuiles
	TileCount   int    `json:"tilecount"`   // Nombre total de tuiles dans l'image
	TileHeight  int    `json:"tileheight"`  // Hauteur d'une tuile
	TileWidth   int    `json:"tilewidth"`   // Largeur d'une tuile
}

// Représente la structure complète d'une carte en JSON (exportée depuis un éditeur de carte)
type MapJSON struct {
	CompressionLevel int       `json:"compressionLevel"` // Niveau de compression utilisé
	Height           int       `json:"height"`           // Hauteur de la carte en nombre de tuiles
	Infinite         bool      `json:"infinite"`         // Si la carte est infinie ou non
	Layers           []Layer   `json:"layers"`           // Liste des calques présents sur la carte
	NextLayerID      int       `json:"nextlayerid"`      // ID du prochain calque
	NextObjectID     int       `json:"nextobjectid"`     // ID du prochain objet
	Orientation      string    `json:"orientation"`      // Orientation de la carte (orthogonale, isométrique, etc.)
	RenderOrder      string    `json:"renderorder"`      // Ordre de rendu des tuiles
	TiledVersion     string    `json:"tiledversion"`     // Version de l'éditeur de carte utilisé
	TileHeight       int       `json:"tileheight"`       // Hauteur d'une tuile
	TileSets         []TileSet `json:"tilesets"`         // Liste des jeux de tuiles utilisés
	TileWidth        int       `json:"tilewidth"`        // Largeur d'une tuile
	Type             string    `json:"type"`             // Type de la carte
	Version          string    `json:"version"`          // Version de la carte
	Width            int       `json:"width"`            // Largeur de la carte en nombre de tuiles
}

// Initialise la carte à partir d'un fichier
func (e *Engine) InitMap(mapFile string) {
	// Lecture du fichier JSON de la carte
	file, err := os.ReadFile(mapFile)
	if err != nil {
		fmt.Println(err) // Affiche une erreur si la lecture échoue
		os.Exit(1)       // Termine le programme en cas d'échec
	}

	// Décode le fichier JSON dans la structure e.MapJSON
	json.Unmarshal(file, &e.MapJSON)

	// Charge toutes les textures nécessaires à partir des TileSets définis dans la carte
	for _, TileSet := range e.MapJSON.TileSets {
		path := path.Dir(mapFile) + "/"                                // Récupère le chemin du dossier contenant le fichier
		e.Sprites[TileSet.Name] = rl.LoadTexture(path + TileSet.Image) // Charge l'image du jeu de tuiles
	}


	// je récupère tous les objets dans le map.json
	for _, Layer := range e.MapJSON.Layers {
		if Layer.Type == "objectgroup" {
			e.Objects = append(e.Objects, Layer.Objects...) // Ajoute les objets du calque à e.Objects
		}
	}
}

// Fonction pour rendre (afficher) la carte à l'écran
func (e *Engine) RenderMap() {
	// Définition de rectangles source et destination pour le rendu des tuiles
	srcRectangle := rl.Rectangle{X: 0, Y: 0, Width: float32(e.MapJSON.TileHeight), Height: float32(e.MapJSON.TileHeight)}
	destRectangle := rl.Rectangle{X: 0, Y: 0, Width: float32(e.MapJSON.TileWidth), Height: float32(e.MapJSON.TileWidth)}
	column_counter := 0 // Compteur pour savoir quand passer à la ligne suivante lors du rendu des tuiles

	// Boucle à travers les calques de la carte
	for _, Layer := range e.MapJSON.Layers {
		// Boucle à travers chaque tuile du calque (définie dans Layer.Data)
		for _, tile := range Layer.Data {
			if tile != 0 { // Si la tuile n'est pas vide
				wantedTileSet := e.MapJSON.TileSets[0] // Sélectionne par défaut le premier jeu de tuiles
				// Cherche le bon TileSet correspondant à l'identifiant de la tuile
				for _, TileSet := range e.MapJSON.TileSets {
					if TileSet.FirstGid <= tile {
						wantedTileSet = TileSet
					}
				}

				// Calcule l'index de la tuile dans le jeu de tuiles
				index := tile - wantedTileSet.FirstGid

				// Définit la position de la tuile dans l'image du jeu de tuiles
				srcRectangle.X = float32(index)
				srcRectangle.Y = 0

				// Si l'index dépasse le nombre de colonnes, ajuste X et Y
				if index >= wantedTileSet.Columns {
					srcRectangle.X = float32(index % wantedTileSet.Columns)
					srcRectangle.Y = float32(index / wantedTileSet.Columns)
				}

				// Ajuste les coordonnées source pour le rendu (taille des tuiles de 16x16 pixels)
				srcRectangle.X *= 16
				srcRectangle.Y *= 16

				// Utilise Raylib pour dessiner la texture à l'écran
				rl.DrawTexturePro(
					e.Sprites[wantedTileSet.Name], // Texture du jeu de tuiles
					srcRectangle,                  // Rectangle source (partie de l'image à afficher)
					destRectangle,                 // Rectangle destination (où l'afficher)
					rl.Vector2{X: 0, Y: 0},        // Point d'origine (ici sans décalage)
					0,                             // Rotation (ici sans rotation)
					rl.White,                      // Couleur (ici blanc, sans teinte)
				)
			}

			// Après chaque tuile, passe à la tuile suivante (vers la droite)
			destRectangle.X += 32
			column_counter += 1

			// Si on atteint la fin de la ligne (largeur de la carte), passe à la ligne suivante
			if column_counter >= e.MapJSON.Width {
				destRectangle.X = 0
				destRectangle.Y += 32
				column_counter = 0
			}
		}

		// Réinitialise les coordonnées après avoir rendu un calque
		destRectangle.X, destRectangle.Y, column_counter = 0, 0, 0
	}
}
