package model

import (
	"os"
	"path"
	"strings"
)

const (
	tilde = "~/"
	sep   = "/"
)

type Browser struct {
	ID          string `yaml:"id,omitempty"`
	Name        string `yaml:"name"`
	FullName    string `yaml:"fullName,omitempty"`
	Description string `yaml:"description"`
	Path        string `yaml:"path"`
	IconURL     string `yaml:"iconUrl"`
	ProfileName string `yaml:"profileName,omitempty"`
}

func (b Browser) JoinPath(elem ...string) string {
	paths := append([]string{b.Path}, elem...)

	return path.Join(paths...)
}

func (b Browser) ResolvePath(elem ...string) string {
	fullPath := b.JoinPath(elem...)

	if fullPath[:2] == tilde {
		homeDir := os.Getenv("HOME")

		fullPath = strings.Replace(fullPath, tilde, homeDir+sep, 1)
	}

	return fullPath
}
