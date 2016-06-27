package main

import (
	"os"
	"testing"
	"bytes"
	"io"
	"github.com/stretchr/testify/assert"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/stringutil"
)

func captureStdout(f func() error) (string, error) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer

	io.Copy(&buf, r)

	return buf.String(), err
}

func TestMain(m *testing.M) {
	os.Setenv(stringutil.KebabCase(browserFile), "testdata/browser.yml")
	os.Setenv(stringutil.KebabCase(profileFile), "Profiles.json")
	os.Setenv(stringutil.KebabCase(bookmarkFile), "Bookmarks.json")

	os.Exit(m.Run())
}

func TestRun(t *testing.T) {
	assert := assert.New(t)

	os.Args = []string{"cli", "bookmarks", "chrome"}

	_, err := captureStdout(func() error {
		main()

		return nil
	})

	assert.NoError(err)
}
