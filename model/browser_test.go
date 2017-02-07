package model

import (
	"github.com/stretchr/testify/assert"
	"os/user"
	"path"
	"strings"
	"testing"
)

const fullPath = "~/Library/Application Support/Google/Chrome"

func pathFor(elem ...string) string {
	paths := append([]string{fullPath}, elem...)

	return path.Join(paths...)
}

func fullPathFor(elem ...string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir

	return strings.Replace(pathFor(elem...), "~/", dir+"/", 1)
}

func TestBrowser_PathFor(t *testing.T) {
	test := assert.New(t)
	browser := &Browser{Path: fullPath}

	test.Equal(path.Join(fullPath, "Default"), browser.PathFor("Default"))
}

func TestBrowser_FullPathFor(t *testing.T) {
	test := assert.New(t)
	browser := &Browser{Path: fullPath}

	test.Equal(fullPathFor(), browser.FullPathFor())
	test.Equal(fullPathFor("Default"), browser.FullPathFor("Default"))
}
