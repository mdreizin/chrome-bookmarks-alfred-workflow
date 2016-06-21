package repository

import (
	"io/ioutil"
	"encoding/json"
	"strings"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
)

type profileAux struct {
	Root struct {
		Profiles	map[string]model.Profile		`json:"info_cache"`
	 	Name		string							`json:"last_used,omitempty"`
	} `json:"profile"`
}

type JsonProfileRepository struct {
	filename	string
}

func (r *JsonProfileRepository) GetProfiles(browser model.Browser) (model.ProfileSlice, error) {
	profiles := model.ProfileSlice{}
	filename := browser.FullPathFor(r.filename)

	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return profiles, err
	}

	aux := profileAux{}
	err = json.Unmarshal(bytes, &aux)

	if err != nil {
		return profiles, err
	}

	profileName := aux.Root.Name

	if browser.ProfileName != "" {
		profileName = browser.ProfileName
	}

	for k, v := range aux.Root.Profiles {
		profiles = profiles.Add(model.Profile{
			Name: k,
			IsActive: strings.EqualFold(k, profileName),
			AvatarURL: v.AvatarURL,
			IconURL: v.AvatarIconURL(browser, k),
			CustomAvatarURL: v.CustomAvatarURL,
			IsDefaultAvatar: v.IsDefaultAvatar,
			DisplayName: v.DisplayName,
			UserName: v.UserName,
			UserEmail: v.UserEmail,
		})
	}

	return profiles, err
}
