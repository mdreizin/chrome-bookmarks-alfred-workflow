package repository

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonProfileRepository_GetProfiles(t *testing.T) {
	test := assert.New(t)
	profileRepository := JsonProfileRepository{filename: "Profiles.json"}
	browser := model.Browser{Path: "testdata"}
	profiles, err := profileRepository.GetProfiles(browser)

	test.NoError(err)
	test.Len(profiles, 3)

	browser.ProfileName = model.DefaultProfileName

	profiles, err = profileRepository.GetProfiles(browser)

	test.NoError(err)
	test.Len(profiles, 3)
}

func TestJsonProfileRepository_GetProfiles_ReadFile(t *testing.T) {
	test := assert.New(t)
	profileRepository := JsonProfileRepository{filename: "Profiles.json"}
	browser := model.Browser{Path: "testdata/Default"}
	profiles, err := profileRepository.GetProfiles(browser)

	test.Error(err)
	test.Len(profiles, 0)
}

func TestJsonProfileRepository_GetProfiles_Unmarshal(t *testing.T) {
	test := assert.New(t)
	profileRepository := JsonProfileRepository{filename: "Profiles"}
	browser := model.Browser{Path: "testdata"}
	profiles, err := profileRepository.GetProfiles(browser)

	test.Error(err)
	test.Len(profiles, 0)
}
