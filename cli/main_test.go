package main

import (
	"bytes"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/stringutil"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
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
	test := assert.New(t)

	os.Args = []string{"cli", "bookmarks", "chrome"}

	_, err := captureStdout(func() error {
		main()

		return nil
	})

	test.NoError(err)
}
