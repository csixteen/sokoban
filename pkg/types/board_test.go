package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	data := []string{
		"wwwwwwww",
		"wffwfffw",
		"wjbggbfw",
		"wfbgbffw",
		"wfwggbfw",
		"wffwfffw",
		"wwfffwww",
		"wwwwwwww",
	}

	board := NewBoard(data)
	w, h := board.Bounds()

	assert.Equal(t, 8, w, "Width should be 8")
	assert.Equal(t, 8, h, "Height should be 8")

	v, _ := board.Get(2, 1)
	assert.Equal(t, 'j', v, "There should be a `j` in cell (2, 1)")
	assert.True(t, isPlayer(v))

	board.Put(1, 1, 'o')
	v, _ = board.Get(1, 1)
	assert.Equal(t, 'o', v, "There should be a `c` in cell (1, 1)")
	assert.True(t, isUnmovable(v))

	v, err := board.Remove(1, 1)
	assert.Error(t, err)
	v, _ = board.Get(1, 1)
	assert.Equal(t, 'o', v, "There should be an `o` in cell (1, 1)")

	r, c := board.FindPlayer()
	assert.Equal(t, 2, r)
	assert.Equal(t, 1, c)

	board.SetPlayerChar('k')
	r, c = board.FindPlayer()
	v, _ = board.Get(r, c)
	assert.Equal(t, 'k', v)
	board.Remove(r, c)
	v, _ = board.Get(r, c)
	assert.Equal(t, v, 'f')
}

func TestMovePlayer(t *testing.T) {
	data := []string{
		"wwwwwwww",
		"wffwfffw",
		"wjbggbfw",
		"wfbgbffw",
		"wfwggbfw",
		"wffwfffw",
		"wwfffwww",
		"wwwwwwww",
	}

	board := NewBoard(data)
	board.MoveDown()
	v, _ := board.Get(2, 1)
	assert.True(t, isFloor(v))
	v, _ = board.Get(3, 1)
	assert.True(t, isPlayer(v))
	r, c := board.FindPlayer()
	assert.Equal(t, 3, r)
	assert.Equal(t, 1, c)

	board.MoveRight()
	r, c = board.FindPlayer()
	assert.Equal(t, 3, r)
	assert.Equal(t, 2, c)
	v, _ = board.Get(3, 2)
	assert.True(t, isPlayer(v))
	v, _ = board.Get(3, 3)
	assert.True(t, isUnmovable(v))
}

func TestMovePlayerTwoBlocks(t *testing.T) {
	data := []string{"jbbgw"}

	board := NewBoard(data)

	board.MoveRight()
	r, c := board.FindPlayer()
	assert.Equal(t, 0, r)
	assert.Equal(t, 1, c)

	v, _ := board.Get(0, 0)
	assert.True(t, isFloor(v))

	v, _ = board.Get(0, 2)
	assert.Equal(t, 'b', v)
	v, _ = board.Get(0, 3)
	assert.Equal(t, 'o', v)
}
