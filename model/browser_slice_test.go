package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
