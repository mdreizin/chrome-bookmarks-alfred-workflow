package repository

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonProfileRepository_GetProfiles(t *testing.T) {
	test := assert.New(t)
	repository := JsonProfileRepository{filename: "Profiles.json"}
	browser := model.Browser{Path: "testdata"}
	profiles, err := repository.GetProfiles(browser)

	test.NoError(err)
	test.Len(profiles, 3)

	browser.ProfileName = "Default"

	profiles, err = repository.GetProfiles(browser)

	test.NoError(err)
	test.Len(profiles, 3)
}

func TestJsonProfileRepository_GetProfiles_ReadFile(t *testing.T) {
	test := assert.New(t)
	repository := JsonProfileRepository{filename: "Profiles.json"}
	browser := model.Browser{Path: "testdata/Default"}
	profiles, err := repository.GetProfiles(browser)

	test.Error(err)
	test.Len(profiles, 0)
}

func TestJsonProfileRepository_GetProfiles_Unmarshal(t *testing.T) {
	test := assert.New(t)
	repository := JsonProfileRepository{filename: "Profiles"}
	browser := model.Browser{Path: "testdata"}
	profiles, err := repository.GetProfiles(browser)

	test.Error(err)
	test.Len(profiles, 0)
}
