package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultBookmarkService_GetBrowsers(t *testing.T) {
	assert := assert.New(t)
	config := map[string]string{
		"browser-file": "testdata/browser.yml",
	}
	bookmarkService := NewBookmarkService(config)
	browsers, err := bookmarkService.GetBrowsers()

	assert.NoError(err)
	assert.Len(browsers, 1)
}

func TestDefaultBookmarkService_GetProfiles(t *testing.T) {
	assert := assert.New(t)
	config := map[string]string{
		"browser-file": "testdata/browser.yml",
		"profile-file": "Profiles.json",
	}
	bookmarkService := NewBookmarkService(config)
	browsers, _ := bookmarkService.GetBrowsers()
	browser, _ := browsers.FindByName("chrome")
	profiles, err := bookmarkService.GetProfiles(browser, "")

	assert.NoError(err)
	assert.Len(profiles, 2)

	profiles, err = bookmarkService.GetProfiles(browser, "Default")

	assert.NoError(err)
	assert.Len(profiles, 1)
}

func TestDefaultBookmarkService_GetBookmarks(t *testing.T) {
	assert := assert.New(t)
	config := map[string]string{
		"browser-file":  "testdata/browser.yml",
		"profile-file":  "Profiles.json",
		"bookmark-file": "Bookmarks.json",
	}
	bookmarkService := NewBookmarkService(config)
	browsers, _ := bookmarkService.GetBrowsers()
	browser, _ := browsers.FindByName("chrome")
	bookmarks, err := bookmarkService.GetBookmarks(browser, "")

	assert.NoError(err)
	assert.Len(bookmarks, 1)

	bookmarks, err = bookmarkService.GetBookmarks(browser, "alfred")

	assert.NoError(err)
	assert.Len(bookmarks, 1)
}

func TestDefaultBookmarkService_UpdateBrowser(t *testing.T) {
	assert := assert.New(t)
	config := map[string]string{
		"browser-file":  "testdata/browser.yml",
		"profile-file":  "Profiles.json",
		"bookmark-file": "Bookmarks.json",
	}
	bookmarkService := NewBookmarkService(config)
	browsers, _ := bookmarkService.GetBrowsers()
	browser, _ := browsers.FindByName("chrome")

	browser.ProfileName = "Default"

	err := bookmarkService.UpdateBrowser(browser)

	browsers, _ = bookmarkService.GetBrowsers()
	browser, _ = browsers.FindByName("chrome")

	assert.NoError(err)
	assert.Equal("Default", browser.ProfileName)
}
