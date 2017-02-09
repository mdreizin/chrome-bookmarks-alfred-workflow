package repository

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestYmlBrowserRepository_GetBrowsers(t *testing.T) {
	test := assert.New(t)
	browserRepository := YmlBrowserRepository{filename: "testdata/browser.yml"}
	browsers, err := browserRepository.GetBrowsers()

	test.NoError(err)
	test.Len(browsers, 3)
	test.Equal("Profile 1", browsers[0].ProfileName)
}

func TestYmlBrowserRepository_GetBrowsers_ReadFileError(t *testing.T) {
	test := assert.New(t)
	browserRepository := YmlBrowserRepository{filename: "testdata"}
	browsers, err := browserRepository.GetBrowsers()

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

	browser.ProfileName = model.DefaultProfileName

	err := repository.UpdateBrowser(browser)

	test.NoError(err)

	browsers, _ = repository.GetBrowsers()
	browser, _ = browsers.FindByName("chrome")

	test.Equal(model.DefaultProfileName, browser.ProfileName)

	os.RemoveAll(dirname)
}
