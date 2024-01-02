package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendFirst(t *testing.T) {
	assert := assert.New(t)
	list := NewList[int]()
	list.append(5)
	fullstring := list.to_string()
	assert.Equal(uint(1), list.length)
	assert.Equal("HEAD -> 5 -> END", fullstring)
}

func TestAppendMore(t *testing.T) {
	assert := assert.New(t)
	list := NewList[int]()
	list.append(5)
	list.append(6)
	list.append(7)
	list.append(8)
	fullstring := list.to_string()
	assert.Equal(uint(4), list.length)
	assert.Equal("HEAD -> 5 -> 6 -> 7 -> 8 -> END", fullstring)
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)
	list := NewList[int]()
	list.append(5)
	list.append(6)
	list.append(7)
	list.append(8)
	list.remove(2)
	assert.Equal(uint(3), list.length)
	fullstring := list.to_string()
	assert.Equal("HEAD -> 5 -> 6 -> 8 -> END", fullstring)
}

func TestRemoveEnd(t *testing.T) {
	assert := assert.New(t)
	list := NewList[int]()
	list.append(5)
	list.append(6)
	list.append(7)
	list.remove(2)
	assert.Equal(uint(2), list.length)
	fullstring := list.to_string()
	assert.Equal("HEAD -> 5 -> 6 -> END", fullstring)
}

func TestRemoveStart(t *testing.T) {
	assert := assert.New(t)
	list := NewList[int]()
	list.append(5)
	list.append(6)
	list.append(7)
	list.remove(0)
	assert.Equal(uint(2), list.length)
	fullstring := list.to_string()
	assert.Equal("HEAD -> 6 -> 7 -> END", fullstring)
}

func TestRemoveEmptyList(t *testing.T) {
	assert := assert.New(t)
	list := NewList[int]()
	err := list.remove(4)
	assert.True(err != nil)
	assert.Equal(err.Error(), "cannot remove from empty list")
}

func TestRemoveInvalidIndex(t *testing.T) {
	assert := assert.New(t)
	list := NewList[int]()
	list.append(56)
	err := list.remove(4)
	assert.True(err != nil)
	assert.Equal(err.Error(), "cannot remove from index 4. List size is 1")
}

func TestPutFirst(t *testing.T) {
	assert := assert.New(t)
	list := NewList[int]()
	list.append(5)
	list.put(9, 0)
	fullstring := list.to_string()
	assert.Equal(uint(2), list.length)
	assert.Equal("HEAD -> 5 -> 9 -> END", fullstring)
}

func TestPutMore(t *testing.T) {
	assert := assert.New(t)
	list := NewList[int]()
	list.append(5)
	list.append(6)
	list.append(7)
	list.append(8)
	list.put(9, 2)
	fullstring := list.to_string()
	assert.Equal(uint(5), list.length)
	assert.Equal("HEAD -> 5 -> 6 -> 7 -> 9 -> 8 -> END", fullstring)
}

func TestGetFirst(t *testing.T) {
	assert := assert.New(t)
	list := NewList[int]()
	list.append(5)
	list.append(6)
	list.append(7)

	test, err := list.get(0)
	assert.True(err == nil)
	assert.Equal(5, test)
}

func TestGetOther(t *testing.T) {
	assert := assert.New(t)
	list := NewList[int]()
	list.append(5)
	list.append(6)
	list.append(7)

	test, err := list.get(2)
	assert.True(err == nil)
	assert.Equal(7, test)
}

func TestGetInvalid(t *testing.T) {
	assert := assert.New(t)
	list := NewList[int]()
	list.append(5)
	list.append(6)
	list.append(7)

	_, err := list.get(5)
	assert.True(err != nil)
	assert.Equal(err.Error(), "cannot get from index 5. List size is 3")
}
