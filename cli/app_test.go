package main

import (
	"fmt"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
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
	test := assert.New(t)

	args := []string{"cli", "bookmarks", "chrome"}

	xml, err := runApp(args)

	test.NoError(err)
	test.Equal(`<?xml version="1.0" encoding="UTF-8"?>
<items><item valid="true" arg="http://mdreizin.github.io/chrome-bookmarks-alfred-workflow" uid="http://mdreizin.github.io/chrome-bookmarks-alfred-workflow"><title>chrome-bookmarks-alfred-workflow</title><subtitle>http://mdreizin.github.io/chrome-bookmarks-alfred-workflow</subtitle><icon>img/chrome.png</icon></item></items>`, xml)
}

func TestApp_GetBookmarks_WithQuery(t *testing.T) {
	test := assert.New(t)

	args := []string{"cli", "bookmarks", "chrome", "-query", "chrome-bookmarks"}

	xml, err := runApp(args)

	test.NoError(err)
	test.Equal(`<?xml version="1.0" encoding="UTF-8"?>
<items><item valid="true" arg="http://mdreizin.github.io/chrome-bookmarks-alfred-workflow" uid="http://mdreizin.github.io/chrome-bookmarks-alfred-workflow"><title>chrome-bookmarks-alfred-workflow</title><subtitle>http://mdreizin.github.io/chrome-bookmarks-alfred-workflow</subtitle><icon>img/chrome.png</icon></item></items>`, xml)
}

func TestApp_GetBookmarks_Error(t *testing.T) {
	test := assert.New(t)

	args := []string{"cli", "bookmarks", "chromium"}

	_, err := runApp(args)

	test.Error(err)
}

func TestApp_GetProfiles(t *testing.T) {
	test := assert.New(t)

	args := []string{"cli", "profiles", "chrome"}

	xml, err := runApp(args)

	test.NoError(err)
	test.Equal(fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<items><item valid="true" arg="" uid=""><title>%s</title><subtitle></subtitle><icon>img/chrome.png</icon></item><item valid="true" arg="Default" uid="Default"><title>Default (Active)</title><subtitle>testdata/Default</subtitle><icon>testdata/Avatars/avatar_alien.png</icon></item></items>`, model.AutoProfile.DisplayName), xml)
}

func TestApp_GetProfiles_WithQuery(t *testing.T) {
	test := assert.New(t)

	args := []string{"cli", "profiles", "chrome", "-query", "Default"}

	xml, err := runApp(args)

	test.NoError(err)
	test.Equal(`<?xml version="1.0" encoding="UTF-8"?>
<items><item valid="true" arg="Default" uid="Default"><title>Default (Active)</title><subtitle>testdata/Default</subtitle><icon>testdata/Avatars/avatar_alien.png</icon></item></items>`, xml)
}

func TestApp_GetProfiles_Error(t *testing.T) {
	test := assert.New(t)

	args := []string{"cli", "profiles", "chromium"}

	_, err := runApp(args)

	test.Error(err)
}

func TestApp_SetProfiles(t *testing.T) {
	test := assert.New(t)

	args := []string{"cli", "select-profile", "chrome", "-query", "Default"}

	str, err := runApp(args)

	test.NoError(err)
	test.Equal("Profile has been successfully updated", str)
}
