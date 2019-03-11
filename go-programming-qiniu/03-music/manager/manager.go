package mlib

import "errors"

// MusicEntry represents one song's info.
type MusicEntry struct {
	ID     string
	Name   string
	Artist string
	Type   string
}

// MusicManager manages the musics.
type MusicManager struct {
	musics []MusicEntry
}

// NewMusicManager creates a Music Manger
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

// Add addes one music into the manager
func (m *MusicManager) Add(entry *MusicEntry) {
	m.musics = append(m.musics, *entry)
}

// RemoveByIndex removes one music into the manager
func (m *MusicManager) RemoveByIndex(index int) (entry *MusicEntry) {
	if index < 0 || index >= len(m.musics) {
		return
	}

	entry = &m.musics[index]

	if index < len(m.musics)-1 {
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 {
		m.musics = make([]MusicEntry, 0)
	} else {
		m.musics = m.musics[:index-1]
	}
	return
}

// GetByIndex returns MusicEntry by given index
func (m MusicManager) GetByIndex(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("index out of range")
	}
	return &m.musics[index], nil
}

// GetByName returns MusicEntry byt given name
func (m MusicManager) GetByName(name string) (music *MusicEntry) {
	if len(m.musics) == 0 {
		return
	}

	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return
}

// Len returns the length of all musics
func (m MusicManager) Len() int {
	return len(m.musics)
}
