package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestYmlBrowserRepository_GetBrowsers(t *testing.T) {
	assert := assert.New(t)
	repository := YmlBrowserRepository{filename: "testdata/browser.yml"}
	browsers, err := repository.GetBrowsers()

	assert.NoError(err)
	assert.Len(browsers, 3)
}

func TestYmlBrowserRepository_GetBrowsers_ReadFileError(t *testing.T) {
	assert := assert.New(t)
	repository := YmlBrowserRepository{filename: "testdata"}
	browsers, err := repository.GetBrowsers()

	assert.Error(err)
	assert.Len(browsers, 0)
}
