package service

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultBookmarkService_GetBrowsers(t *testing.T) {
	test := assert.New(t)
	config := map[string]string{
		"browser-file": "testdata/browser.yml",
	}
	bookmarkService := NewBookmarkService(config)
	browsers, err := bookmarkService.GetBrowsers()

	test.NoError(err)
	test.Len(browsers, 1)
}

func TestDefaultBookmarkService_GetProfiles(t *testing.T) {
	test := assert.New(t)
	config := map[string]string{
		"browser-file": "testdata/browser.yml",
		"profile-file": "Profiles.json",
	}
	bookmarkService := NewBookmarkService(config)
	browsers, _ := bookmarkService.GetBrowsers()
	browser, _ := browsers.FindByName("chrome")
	profiles, err := bookmarkService.GetProfiles(browser, "")

	test.NoError(err)
	test.Len(profiles, 2)

	profiles, err = bookmarkService.GetProfiles(browser, model.DefaultProfileName)

	test.NoError(err)
	test.Len(profiles, 1)
}

func TestDefaultBookmarkService_GetBookmarks(t *testing.T) {
	test := assert.New(t)
	config := map[string]string{
		"browser-file":  "testdata/browser.yml",
		"profile-file":  "Profiles.json",
		"bookmark-file": "Bookmarks.json",
	}
	bookmarkService := NewBookmarkService(config)
	browsers, _ := bookmarkService.GetBrowsers()
	browser, _ := browsers.FindByName("chrome")
	bookmarks, err := bookmarkService.GetBookmarks(browser, "")

	test.NoError(err)
	test.Len(bookmarks, 1)

	bookmarks, err = bookmarkService.GetBookmarks(browser, "alfred")

	test.NoError(err)
	test.Len(bookmarks, 1)
}

func TestDefaultBookmarkService_UpdateBrowser(t *testing.T) {
	test := assert.New(t)
	config := map[string]string{
		"browser-file":  "testdata/browser.yml",
		"profile-file":  "Profiles.json",
		"bookmark-file": "Bookmarks.json",
	}
	bookmarkService := NewBookmarkService(config)
	browsers, _ := bookmarkService.GetBrowsers()
	browser, _ := browsers.FindByName("chrome")

	browser.ProfileName = model.DefaultProfileName

	err := bookmarkService.UpdateBrowser(browser)

	browsers, _ = bookmarkService.GetBrowsers()
	browser, _ = browsers.FindByName("chrome")

	test.NoError(err)
	test.Equal(model.DefaultProfileName, browser.ProfileName)
}
