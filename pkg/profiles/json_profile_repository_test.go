package profiles

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonProfileRepository_GetProfiles(t *testing.T) {
	test := assert.New(t)
	profileRepository := JsonProfileRepository{Filename: "Profiles.json"}
	browser := browsers.Browser{Path: "test"}
	profileSlice, err := profileRepository.GetProfiles(&browser)

	test.NoError(err)
	test.Len(profileSlice, 3)

	browser.ProfileName = DefaultProfileName

	profileSlice, err = profileRepository.GetProfiles(&browser)

	test.NoError(err)
	test.Len(profileSlice, 3)
}

func TestJsonProfileRepository_GetProfiles_ReadFile(t *testing.T) {
	test := assert.New(t)
	profileRepository := JsonProfileRepository{Filename: "Profiles.json"}
	browser := browsers.Browser{Path: "test/Default"}
	profileSlice, err := profileRepository.GetProfiles(&browser)

	test.Error(err)
	test.Len(profileSlice, 0)
}

func TestJsonProfileRepository_GetProfiles_Unmarshal(t *testing.T) {
	test := assert.New(t)
	profileRepository := JsonProfileRepository{Filename: "Profiles"}
	browser := browsers.Browser{Path: "test"}
	profileSlice, err := profileRepository.GetProfiles(&browser)

	test.Error(err)
	test.Len(profileSlice, 0)
}
