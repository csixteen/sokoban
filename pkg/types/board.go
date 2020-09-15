package types

import (
	"errors"

	u "github.com/csixteen/sokoban/pkg/utils"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Board struct {
	matrix        [][]*u.Stack
	width, height int
	pRow, pCol    int // Player coordinates on the board
	goals         int
}

// NewBoard generates a new Board from an array of strings, where
// each character represents an element on the board.
//
// Types of characters:
//  w - wall
//  b - movable block
//  h - player facing left
//  j - player facing down
//  k - player facing up
//  l - player facing right
//  f - floor
//  g - goal (where you must place a block onto)
//  o - block on top of goal (becomes unmovable)
func NewBoard(data []string) *Board {
	rows := len(data)
	cols := len(data[0])

	var goals, pRow, pCol int

	m := make([][]*u.Stack, rows)
	for row := 0; row < rows; row++ {
		m[row] = make([]*u.Stack, cols)
		for col, c := range data[row] {
			m[row][col] = u.NewStack()
			m[row][col].Push('f')

			if isFloor(c) {
				continue
			}

			if isGoal(c) {
				goals++
			}

			m[row][col].Push(c)

			if isPlayer(c) {
				pCol = col
				pRow = row
			}
		}
	}

	return &Board{
		matrix: m,
		width:  cols,
		height: rows,
		pRow:   pRow,
		pCol:   pCol,
		goals:  goals,
	}
}

///-------------------------------
///        Game action

func (b *Board) IsVictory() bool {
	return b.goals == 0
}

///----------------------------------------------------------
///                 Board manipulation

// Bounds returns a pair (width, height) representing the
// bounds of the board.
func (b *Board) Bounds() (int, int) {
	return b.width, b.height
}

// Get returns the rune that's on position (row, col) of the board
func (b *Board) Get(row, col int) (rune, error) {
	return b.matrix[row][col].Top()
}

// Put puts the rune `elem` on the position (row, col) in the board
func (b *Board) Put(row, col int, elem rune) {
	b.matrix[row][col].Push(elem)
}

// Remove removes the rune that's on the position (row, col) of the board
func (b *Board) Remove(row, col int) (rune, error) {
	val, err := b.Get(row, col)
	if err != nil {
		return -1, err
	}

	if isFloor(val) || isUnmovable(val) {
		return -1, errors.New("Can't remove such element from the board")
	}

	return b.matrix[row][col].Pop()
}

// Moves an element from (sRow, sCol) to the adjacent cell according to
// the direction. If the element is unmovable, it fails. If the element
// is movable, it will check whether the adjacent cell is walkable (either
// floor or goal) or contains a movable object. If it contains a movable
// object, it will try to move it as well, in a recursive fashion.
func (b *Board) MoveFrom(sRow, sCol int, d Direction) (int, int, error) {
	elem, _ := b.Get(sRow, sCol)
	if !isMovable(elem) {
		return -1, -1, errors.New("Cannot move unmovable element")
	}

	nextRow := sRow
	nextCol := sCol

	switch d {
	case Up:
		nextRow--
	case Down:
		nextRow++
	case Left:
		nextCol--
	case Right:
		nextCol++
	}

	nextElem, _ := b.Get(nextRow, nextCol)
	if isWalkable(nextElem) {
		elem, _ = b.Remove(sRow, sCol)
		if isGoal(nextElem) && !isPlayer(elem) {
			elem = 'o'
			b.goals--
		}
		b.Put(nextRow, nextCol, elem)
		return nextRow, nextCol, nil
	}

	_, _, err := b.MoveFrom(nextRow, nextCol, d)
	if err != nil {
		return -1, -1, err
	}

	return b.MoveFrom(sRow, sCol, d)
}

///-------------------------------------------------------------
///         Board elements assertions and predicates

func isPlayer(c rune) bool {
	return c == 'h' || c == 'j' || c == 'k' || c == 'l'
}

func isGoal(c rune) bool {
	return c == 'g'
}

func isMovable(c rune) bool {
	return c == 'b' || isPlayer(c)
}

func isUnmovable(c rune) bool {
	return c == 'w' || c == 'o'
}

func isFloor(c rune) bool {
	return c == 'f'
}

func isWalkable(c rune) bool {
	return isGoal(c) || isFloor(c)
}

///-----------------------------------------
///          Player manipulation

// FindPlayer returns a pair (row, col) representing the location
// of the player on the board.
func (b *Board) FindPlayer() (int, int) {
	return b.pRow, b.pCol
}

func (b *Board) SetPlayerPos(row, col int) {
	b.pRow = row
	b.pCol = col
}

func (b *Board) SetPlayerChar(c rune) {
	row, col := b.FindPlayer()
	b.matrix[row][col].Pop()
	b.matrix[row][col].Push(c)
}

func (b *Board) movePlayer(d Direction) {
	switch d {
	case Up:
		b.SetPlayerChar('k')
	case Down:
		b.SetPlayerChar('j')
	case Left:
		b.SetPlayerChar('h')
	case Right:
		b.SetPlayerChar('l')
	}

	r, c := b.FindPlayer()
	row, col, err := b.MoveFrom(r, c, d)
	if err == nil {
		b.SetPlayerPos(row, col)
	}
}

func (b *Board) MoveRight() {
	b.movePlayer(Right)
}

func (b *Board) MoveLeft() {
	b.movePlayer(Left)
}

func (b *Board) MoveUp() {
	b.movePlayer(Up)
}

func (b *Board) MoveDown() {
	b.movePlayer(Down)
}
