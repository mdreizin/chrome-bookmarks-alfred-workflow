package main

import (
	"os"
	"testing"
	"bytes"
	"io"
	"github.com/stretchr/testify/assert"
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
	os.Setenv("BROWSER_FILENAME", "testdata/browser.yml")
	os.Setenv("PROFILE_FILENAME", "Profiles.json")
	os.Setenv("BOOKMARK_FILENAME", "Bookmarks.json")

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
