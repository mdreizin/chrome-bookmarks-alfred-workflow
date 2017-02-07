package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBrowserSlice_Match(t *testing.T) {
	test := assert.New(t)
	browsers := BrowserSlice{
		Browser{
			Name: "canary",
		},
		Browser{
			Name: "chrome",
		},
	}

	test.Len(browsers.Match("canary"), 1)
	test.Len(browsers.Match("chrome"), 1)
	test.Len(browsers.Match("c"), 2)
}

func TestBrowserSlice_FindByName(t *testing.T) {
	test := assert.New(t)
	browsers := BrowserSlice{
		Browser{Name: "canary"},
		Browser{Name: "chrome"},
	}
	browser, err := browsers.FindByName("chrome")

	test.NoError(err)
	test.Equal(Browser{Name: "chrome"}, browser)

	browser, err = browsers.FindByName("chromium")

	test.Error(err)
}

func TestBrowserSlice_FindIndex(t *testing.T) {
	test := assert.New(t)
	browsers := BrowserSlice{
		Browser{
			Name: "canary",
		},
		Browser{
			Name: "chrome",
		},
	}
	i := browsers.FindIndex(func(b Browser) bool { return b.Name == "chrome" })

	test.Equal(browsers[i], browsers[1])

	i = browsers.FindIndex(func(b Browser) bool { return b.Name == "chromium" })

	test.Equal(i, -1)
}
