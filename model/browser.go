package model

import (
	"path"
	"strings"
	"os/user"
)

const (
	tilde = "~/"
	sep = "/"
)

type Browser struct {
	ID				string	`yaml:"id,omitempty"`
	Name			string	`yaml:"name"`
	FullName		string	`yaml:"fullName,omitempty"`
	Description		string	`yaml:"description"`
	Path			string	`yaml:"path"`
	IconURL			string	`yaml:"iconUrl"`
	ProfileName		string	`yaml:"profileName,omitempty"`
}

func (b Browser) PathFor(elem ...string) string {
	paths := append([]string{b.Path}, elem...)

	return path.Join(paths...)
}

func (b Browser) FullPathFor(elem ...string) string {
	fullPath := b.PathFor(elem...)

	if fullPath[:2] == tilde {
		usr, _ := user.Current()
		dir := usr.HomeDir

		fullPath = strings.Replace(fullPath, tilde, dir + sep, 1)
	}

	return fullPath
}
