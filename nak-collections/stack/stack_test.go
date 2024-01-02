package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	assert := assert.New(t)
	stack := NewStack(5)
	stack.push(6)
	stack.push(7)
	stack.push(8)
	fullstack := stack.to_string()
	assert.Equal(uint(4), stack.length)
	assert.Equal("HEAD -> 8 -> 7 -> 6 -> 5 -> END", fullstack)
}

func TestPush_strings(t *testing.T) {
	assert := assert.New(t)
	stack := NewStack("t")
	stack.push("e")
	stack.push("s")
	stack.push("t")
	fullstack := stack.to_string()
	assert.Equal(uint(4), stack.length)
	assert.Equal("HEAD -> t -> s -> e -> t -> END", fullstack)
}

func TestPop(t *testing.T) {
	assert := assert.New(t)
	stack := NewStack(5)
	stack.push(6)
	stack.push(7)
	stack.push(8)
	popped := stack.pop()
	assert.Equal(8, *popped)
	assert.Equal(uint(3), stack.length)
	popped = stack.pop()
	assert.Equal(7, *popped)
	fullstack := stack.to_string()
	assert.Equal(uint(2), stack.length)
	assert.Equal("HEAD -> 6 -> 5 -> END", fullstack)
}

func TestIsZero(t *testing.T) {
	assert := assert.New(t)
	assert.True(IsZero(0))
	zero := *new(int)
	assert.True(IsZero(zero))
}

func TestIsNotZero(t *testing.T) {
	assert := assert.New(t)
	assert.False(IsZero(1))
	assert.True(!IsZero(1))
}
