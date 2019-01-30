package browsers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBrowserSlice_Match(t *testing.T) {
	test := assert.New(t)
	browserSlice := BrowserSlice{
		Browser{
			Name: "canary",
		},
		Browser{
			Name: "chrome",
		},
	}

	test.Len(browserSlice.Match("canary"), 1)
	test.Len(browserSlice.Match("chrome"), 1)
	test.Len(browserSlice.Match("c"), 2)
}

func TestBrowserSlice_FindByName(t *testing.T) {
	test := assert.New(t)
	browserSlice := BrowserSlice{
		Browser{Name: "canary"},
		Browser{Name: "chrome"},
	}
	browser, err := browserSlice.FindByName("chrome")

	test.NoError(err)
	test.Equal(&Browser{Name: "chrome"}, browser)

	browser, err = browserSlice.FindByName("chromium")

	test.Error(err)
}

func TestBrowserSlice_FindIndex(t *testing.T) {
	test := assert.New(t)
	browserSlice := BrowserSlice{
		Browser{
			Name: "canary",
		},
		Browser{
			Name: "chrome",
		},
	}
	i := browserSlice.FindIndex(func(b *Browser) bool { return b.Name == "chrome" })

	test.Equal(browserSlice[i], browserSlice[1])

	i = browserSlice.FindIndex(func(b *Browser) bool { return b.Name == "chromium" })

	test.Equal(i, -1)
}
