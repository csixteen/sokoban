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

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushToStack(t *testing.T) {
	s := NewStack()

	s.Push('w')

	val, _ := s.Top()
	assert.Equal(t, 'w', val, "Top of the stack should contain `w`")
	assert.Equal(t, 1, s.Size(), "Size of stack should be 1")
}

func TestPopFromStack(t *testing.T) {
	s := NewStack()

	s.Push('w')

	val, _ := s.Pop()
	assert.Equal(t, 'w', val, "Val should be `w`")
	assert.Equal(t, 0, s.Size(), "Size of stack should be 0")
}

func TestPopFromEmptyStack(t *testing.T) {
	s := NewStack()

	_, err := s.Pop()
	assert.Error(t, err)
}

func TestTopFromEmptyStack(t *testing.T) {
	s := NewStack()

	_, err := s.Top()
	assert.Error(t, err)
}
