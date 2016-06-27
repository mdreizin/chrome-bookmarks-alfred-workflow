package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"io/ioutil"
)

func TestYmlBrowserRepository_GetBrowsers(t *testing.T) {
	assert := assert.New(t)
	repository := YmlBrowserRepository{filename: "testdata/browser.yml"}
	browsers, err := repository.GetBrowsers()

	assert.NoError(err)
	assert.Len(browsers, 3)
	assert.Equal("Profile 1", browsers[0].ProfileName)
}

func TestYmlBrowserRepository_GetBrowsers_ReadFileError(t *testing.T) {
	assert := assert.New(t)
	repository := YmlBrowserRepository{filename: "testdata"}
	browsers, err := repository.GetBrowsers()

	assert.Error(err)
	assert.Len(browsers, 0)
}

func TestYmlBrowserRepository_UpdateBrowser(t *testing.T) {
	assert := assert.New(t)
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

	assert.NoError(err)

	browsers, _ = repository.GetBrowsers()
	browser, _ = browsers.FindByName("chrome")

	assert.Equal("Default", browser.ProfileName)

	os.RemoveAll(dirname)
}
