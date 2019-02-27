package browsers

import (
	"github.com/stretchr/testify/assert"
	"os/user"
	"path"
	"strings"
	"testing"
)

const (
	fullPath           = "~/Library/Application Support/Google/Chrome"
	defaultProfileName = "Default"
)

func pathFor(elem ...string) string {
	paths := append([]string{fullPath}, elem...)

	return path.Join(paths...)
}

func resolvePath(elem ...string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir

	return strings.Replace(pathFor(elem...), tilde, dir+sep, 1)
}

func TestBrowser_JoinPath(t *testing.T) {
	test := assert.New(t)
	browser := &Browser{Path: fullPath}

	test.Equal(path.Join(fullPath, defaultProfileName), browser.JoinPath(defaultProfileName))
}

func TestBrowser_ResolvePath(t *testing.T) {
	test := assert.New(t)
	browser := &Browser{Path: fullPath}

	test.Equal(resolvePath(), browser.ResolvePath())
	test.Equal(resolvePath(defaultProfileName), browser.ResolvePath(defaultProfileName))
}
