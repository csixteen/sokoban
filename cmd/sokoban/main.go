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
	PuzzleSize  = 8
	Dim         = TileSize * PuzzleSize
)

var (
	frames       = 0
	second       = time.Tick(time.Second)
	level        = 0
	showingText  = false
	textDuration = 2
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

	for row := 0; row < PuzzleSize; row += 1 {
		for col := 0; col < PuzzleSize; col += 1 {
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
	dt := float64(len(p))
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(Dim/2-2*dt, Dim/2), basicAtlas)
	fmt.Fprintf(basicTxt, p, args...)

	win.Clear(colornames.Black)
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2))
	win.Update()

	showingText = true
	textDuration = duration
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Sokoban",
		Bounds: pixel.R(0, 0, Dim, Dim),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

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

	levels := LoadLevels(LevelsPath)
	board := t.NewBoard(levels[level])
	batch := pixel.NewBatch(&pixel.TrianglesData{}, sprites)

	displayText(win, 2, "Level %d", level+1)

	for !win.Closed() {
		if !showingText {
			detectKeyPress(win, board)

			win.Clear(colornames.Darkslategray)

			drawBoard(win, batch, sprites, tileFrames, board)

			if board.IsVictory() {
				level++
				if level == len(levels) {
					win.SetClosed(true)
				} else {
					displayText(win, 2, "Level %d", level+1)
					board = t.NewBoard(levels[level])
				}
			}

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
