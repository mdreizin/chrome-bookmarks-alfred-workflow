package repository

import (
	"encoding/json"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/stringutil"
	"io/ioutil"
	"strings"
)

type profileAux struct {
	Root struct {
		Profiles map[string]model.Profile `json:"info_cache"`
		Name     string                   `json:"last_used,omitempty"`
	} `json:"profile"`
}

type JsonProfileRepository struct {
	filename string
}

func (r *JsonProfileRepository) GetProfiles(browser model.Browser) (model.ProfileSlice, error) {
	profiles := model.ProfileSlice{}
	filename := browser.ResolvePath(r.filename)

	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return profiles, err
	}

	aux := profileAux{}
	err = json.Unmarshal(bytes, &aux)

	if err != nil {
		return profiles, err
	}

	profileName := stringutil.DefaultIfEmpty(browser.ProfileName, aux.Root.Name)

	aux.Root.Profiles[model.AutoProfile.Name] = model.AutoProfile

	for k, v := range aux.Root.Profiles {
		name := stringutil.DefaultIfEmpty(v.Name, k)

		profiles = profiles.Add(model.Profile{
			Name:            name,
			IsVirtual:       v.IsVirtual,
			IsActive:        strings.EqualFold(name, profileName),
			AvatarURL:       v.AvatarURL,
			IconURL:         v.AvatarIconURL(browser, k),
			CustomAvatarURL: v.CustomAvatarURL,
			IsDefaultAvatar: v.IsDefaultAvatar,
			DisplayName:     v.DisplayName,
			UserName:        v.UserName,
			UserEmail:       v.UserEmail,
		})
	}

	return profiles, err
}
