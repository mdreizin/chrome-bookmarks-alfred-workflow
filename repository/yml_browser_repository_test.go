package repository

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestYmlBrowserRepository_GetBrowsers(t *testing.T) {
	test := assert.New(t)
	repository := YmlBrowserRepository{filename: "testdata/browser.yml"}
	browsers, err := repository.GetBrowsers()

	test.NoError(err)
	test.Len(browsers, 3)
	test.Equal("Profile 1", browsers[0].ProfileName)
}

func TestYmlBrowserRepository_GetBrowsers_ReadFileError(t *testing.T) {
	test := assert.New(t)
	repository := YmlBrowserRepository{filename: "testdata"}
	browsers, err := repository.GetBrowsers()

	test.Error(err)
	test.Len(browsers, 0)
}

func TestYmlBrowserRepository_UpdateBrowser(t *testing.T) {
	test := assert.New(t)
	filename := path.Join("testdata", "marat", "browser.yml") // uuid.NewV4().String()
	dirname := path.Dir(filename)
	bytes, _ := ioutil.ReadFile("testdata/browser.yml")

	os.MkdirAll(dirname, 0777)
	ioutil.WriteFile(filename, bytes, 0644)

	repository := YmlBrowserRepository{filename: filename}
	browsers, _ := repository.GetBrowsers()
	browser, _ := browsers.FindByName("chrome")

	browser.ProfileName = "Default"

	err := repository.UpdateBrowser(browser)

	test.NoError(err)

	browsers, _ = repository.GetBrowsers()
	browser, _ = browsers.FindByName("chrome")

	test.Equal("Default", browser.ProfileName)

	os.RemoveAll(dirname)
}
