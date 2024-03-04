package render

import (
	"sandbox/lib/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

// Tilemap is used to read in a tilemap and returns a 2D array of tiles
func Tilemap(path string, tileSize, width, height int) [][]*Tile {
	img := assets.LoadImage(path)
	tiles := make([][]*Tile, width)

	for i := 0; i < width; i++ {
		tiles[i] = make([]*Tile, height)
		for j := 0; j < height; j++ {
			tiles[i][j] = &Tile{
				X:       i * tileSize,
				Y:       j * tileSize,
				Width:   tileSize,
				Height:  tileSize,
				Texture: img,
			}
		}
	}

	return tiles
}

// Tile is a struct that represents a tile
type Tile struct {
	X       int
	Y       int
	Width   int
	Height  int
	Texture *ebiten.Image
}
