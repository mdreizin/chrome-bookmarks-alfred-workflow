package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBrowserSlice_Match(t *testing.T) {
	assert := assert.New(t)
	browsers := BrowserSlice{
		Browser{
			Name: "canary",
		},
		Browser{
			Name: "chrome",
		},
	}

	assert.Len(browsers.Match("canary"), 1)
	assert.Len(browsers.Match("chrome"), 1)
	assert.Len(browsers.Match("c"), 2)
}

func TestBrowserSlice_FindByName(t *testing.T) {
	assert := assert.New(t)
	browsers := BrowserSlice{
		Browser{Name: "canary"},
		Browser{Name: "chrome"},
	}
	browser, err := browsers.FindByName("chrome")

	assert.NoError(err)
	assert.Equal(Browser{Name: "chrome"}, browser)

	browser, err = browsers.FindByName("chromium")

	assert.Error(err)
}

func TestBrowserSlice_FindIndex(t *testing.T) {
	assert := assert.New(t)
	browsers := BrowserSlice{
		Browser{
			Name: "canary",
		},
		Browser{
			Name: "chrome",
		},
	}
	i := browsers.FindIndex(func(b Browser) bool { return b.Name == "chrome" })

	assert.Equal(browsers[i], browsers[1])

	i = browsers.FindIndex(func(b Browser) bool { return b.Name == "chromium" })

	assert.Equal(i, -1)
}
