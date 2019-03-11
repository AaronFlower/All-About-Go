package mlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	manager := NewMusicManager()
	assert.Equal(manager.Len(), 0, "The length of an initial manager's musics should be zero.")
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	manager := NewMusicManager()

	manager.Add(&MusicEntry{ID: "1", Name: "Black Humor"})
	assert.Equal(manager.Len(), 1, "1 Length after Add one music.")
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)
	manager := NewMusicManager()

	manager.Add(&MusicEntry{ID: "1", Name: "Black Humor"})
	assert.Equal(manager.Len(), 1, "1 Length after Add one music.")
	manager.RemoveByIndex(0)
	assert.Equal(manager.Len(), 0, "0 Length after Remove one music.")
}
