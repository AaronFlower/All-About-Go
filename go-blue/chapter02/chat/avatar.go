package main

import (
	"errors"
	"io/ioutil"
	"path"
)

// ErrNoAvatarURL is the error that is returned when the
// avatar instance is unable to provide an avatar URL.
var ErrNoAvatarURL = errors.New("chat: Unable to get an Avatar URL")

// Avatar represents types capable of representing user profile pictures.
type Avatar interface {
	GetAvatarURL(ChatUser) (string, error)
}

// TryAvatars type is simply a slice of Avatar objects that we are free
// to add methods to.
type TryAvatars []Avatar

// GetAvatarURL returns an avatar's URL
func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

// AuthAvatar implements the Avatar interface.
type AuthAvatar struct{}

// UseAuthAvatar uses
var UseAuthAvatar AuthAvatar

// GetAvatarURL returns the client URL
func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if len(url) == 0 {
		return "", ErrNoAvatarURL
	}
	return url, nil
}

// GravatarAvatar defines a Gr avatar
type GravatarAvatar struct{}

// UseGravatar initiates a placeholder, a helpful variable
var UseGravatar GravatarAvatar

// GetAvatarURL returns the avatar URL
func (GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

// FileSystemAvatar defines a avatar service using file system.
type FileSystemAvatar struct{}

// UseFileSystemAvatar defines a placeholder
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL returns a client avatar
func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	userID := u.UniqueID()
	files, err := ioutil.ReadDir("avatars")
	if err != nil {
		return "", ErrNoAvatarURL
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if match, _ := path.Match(userID+"*", file.Name()); match {
			return "/avatars/" + file.Name(), nil
		}
	}
	return "", ErrNoAvatarURL
}
