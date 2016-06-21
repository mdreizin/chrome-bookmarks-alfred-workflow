package main

import (
	"testing"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

func runApp(args []string) (string, error) {
	xml, err := captureStdout(func() error {
		app := newApp()
		app.Writer = ioutil.Discard

		return app.Run(args)
	})

	return xml, err
}

func TestApp_GetBookmarks(t *testing.T) {
	assert := assert.New(t)

	args := []string{"cli", "bookmarks", "chrome"}

	xml, err := runApp(args)

	assert.NoError(err)
	assert.Equal(xml, `<?xml version="1.0" encoding="UTF-8"?>
<items><item valid="true" arg="http://mdreizin.github.io/chrome-bookmarks-alfred-workflow" uid="http://mdreizin.github.io/chrome-bookmarks-alfred-workflow"><title>chrome-bookmarks-alfred-workflow</title><subtitle>http://mdreizin.github.io/chrome-bookmarks-alfred-workflow</subtitle><icon>img/chrome.png</icon></item></items>`)
}

func TestApp_GetBookmarks_WithQuery(t *testing.T) {
	assert := assert.New(t)

	args := []string{"cli", "bookmarks", "chrome", "-query", "chrome-bookmarks"}

	xml, err := runApp(args)

	assert.NoError(err)
	assert.Equal(xml, `<?xml version="1.0" encoding="UTF-8"?>
<items><item valid="true" arg="http://mdreizin.github.io/chrome-bookmarks-alfred-workflow" uid="http://mdreizin.github.io/chrome-bookmarks-alfred-workflow"><title>chrome-bookmarks-alfred-workflow</title><subtitle>http://mdreizin.github.io/chrome-bookmarks-alfred-workflow</subtitle><icon>img/chrome.png</icon></item></items>`)
}

func TestApp_GetBookmarks_Error(t *testing.T) {
	assert := assert.New(t)

	args := []string{"cli", "bookmarks", "chromium"}

	_, err := runApp(args)

	assert.Error(err)
}

func TestApp_GetProfiles(t *testing.T) {
	assert := assert.New(t)

	args := []string{"cli", "profiles", "chrome"}

	xml, err := runApp(args)

	assert.NoError(err)
	assert.Equal(xml, `<?xml version="1.0" encoding="UTF-8"?>
<items><item valid="true" arg="Default" uid="Default"><title>Default (Active)</title><subtitle>testdata/chrome/Default</subtitle><icon>testdata/Avatars/avatar_alien.png</icon></item></items>`)
}

func TestApp_Profiles_WithQuery(t *testing.T) {
	assert := assert.New(t)

	args := []string{"cli", "profiles", "chrome", "-query", "Default"}

	xml, err := runApp(args)

	assert.NoError(err)
	assert.Equal(xml, `<?xml version="1.0" encoding="UTF-8"?>
<items><item valid="true" arg="Default" uid="Default"><title>Default (Active)</title><subtitle>testdata/chrome/Default</subtitle><icon>testdata/Avatars/avatar_alien.png</icon></item></items>`)
}

func TestApp_GetProfiles_Error(t *testing.T) {
	assert := assert.New(t)

	args := []string{"cli", "profiles", "chromium"}

	_, err := runApp(args)

	assert.Error(err)
}
