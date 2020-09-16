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
