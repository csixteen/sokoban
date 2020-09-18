// MIT License
//
// Copyright (c) 2020 Pedro Rodrigues
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"fmt"
	"image"
	"time"

	_ "image/png"

	t "github.com/csixteen/sokoban/pkg/types"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/markbates/pkger"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const (
	SpritesPath = "/assets/sprites/sokoban_tilesheet.png"
	LevelsPath  = "/assets/levels/levels.dat"
	TileSize    = 64
)

var (
	frames       = 0
	second       = time.Tick(time.Second)
	showingText  = false
	textDuration = 2
)

var (
	allLevels    [][]string
	currentLevel = 0
	board        *t.Board
	boardWidth   int
	boardHeight  int
)

var CharToTile = map[rune]int{
	'w': 49,
	'b': 15,
	'g': 98,
	'o': 14,
	// Char directions (Vim keys)
	'h': 2,
	'j': 0,
	'k': 24,
	'l': 26,
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := pkger.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

func detectKeyPress(w *pixelgl.Window, board *t.Board) {
	if w.JustPressed(pixelgl.KeyLeft) {
		board.MoveUp()
	}
	if w.JustPressed(pixelgl.KeyRight) {
		board.MoveDown()
	}
	if w.JustPressed(pixelgl.KeyDown) {
		board.MoveLeft()
	}
	if w.JustPressed(pixelgl.KeyUp) {
		board.MoveRight()
	}
	if w.JustPressed(pixelgl.KeyQ) {
		w.SetClosed(true)
	}
	if w.JustPressed(pixelgl.KeyR) {
		board.Reset()
	}
}

func drawBoard(
	win *pixelgl.Window,
	batch *pixel.Batch,
	sprites pixel.Picture,
	frames []pixel.Rect,
	board *t.Board,
) {
	batch.Clear()

	for row := 0; row < boardHeight; row += 1 {
		for col := 0; col < boardWidth; col += 1 {
			val, _ := board.Get(row, col)
			if tile, ok := CharToTile[val]; ok {
				elem := pixel.NewSprite(sprites, frames[tile])
				r := float64(row*TileSize) + TileSize/2
				c := float64(col*TileSize) + TileSize/2
				elem.Draw(batch, pixel.IM.Moved(pixel.V(r, c)))
			}
		}
	}

	batch.Draw(win)
}

func displayText(win *pixelgl.Window, duration int, p string, args ...interface{}) {
	dt := float64(len(p)) * 13 / 2 // 13 because Face7x13
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(
		pixel.V(
			float64((boardHeight*TileSize)/2)-dt,
			float64((boardWidth*TileSize)/2),
		),
		basicAtlas,
	)
	fmt.Fprintf(basicTxt, p, args...)

	win.Clear(colornames.Black)
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2))
	win.Update()

	showingText = true
	textDuration = duration
}

func run() {
	//--------------------------------------------
	//    Initialize sprites and tile frames

	sprites, err := loadPicture(SpritesPath)
	if err != nil {
		panic(err)
	}

	var tileFrames []pixel.Rect
	for x := sprites.Bounds().Min.X; x < sprites.Bounds().Max.X; x += TileSize {
		for y := sprites.Bounds().Min.Y; y < sprites.Bounds().Max.Y; y += TileSize {
			tileFrames = append(tileFrames, pixel.R(x, y, x+TileSize, y+TileSize))
		}
	}

	batch := pixel.NewBatch(&pixel.TrianglesData{}, sprites)

	//----------------------------------------------
	//     Load levels and create a new board

	allLevels = loadLevels(LevelsPath)
	board = t.NewBoard(allLevels[currentLevel])
	boardWidth, boardHeight = board.Bounds()

	cfg := pixelgl.WindowConfig{
		Title: "Sokoban",
		Bounds: pixel.R(
			0,
			0,
			float64(boardHeight*TileSize),
			float64(boardWidth*TileSize),
		),
		VSync: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	displayText(win, 2, "Level %d", currentLevel+1)

	// main loop
	for !win.Closed() {
		if board.IsVictory() {
			currentLevel++
			if currentLevel == len(allLevels) {
				win.SetClosed(true)
			} else {
				displayText(win, 2, "Level %d", currentLevel+1)
				board = t.NewBoard(allLevels[currentLevel])
				boardWidth, boardHeight = board.Bounds()
				win.SetBounds(pixel.R(
					0,
					0,
					float64(boardHeight*TileSize),
					float64(boardWidth*TileSize),
				))
			}
		}

		if !showingText {
			detectKeyPress(win, board)

			win.Clear(colornames.Darkslategray)

			drawBoard(win, batch, sprites, tileFrames, board)

			win.Update()
		}

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0

			if showingText {
				textDuration--
				if textDuration == 0 {
					showingText = false
				}
			}
		default:
		}
	}
}

func main() {
	pkger.Include(SpritesPath)
	pkger.Include(LevelsPath)

	pixelgl.Run(run)
}
