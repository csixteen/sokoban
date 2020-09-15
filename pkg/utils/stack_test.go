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
