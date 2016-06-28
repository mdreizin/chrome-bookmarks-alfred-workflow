package model

import (
	"strings"
	"os"
)

const avatarResourcePrefix = "chrome://theme/"

// https://chromium.googlesource.com/chromium/src/+/master/chrome/app/theme/theme_resources.grd
var avatarResources = map[string]string {
	"IDR_PROFILE_AVATAR_0": "avatar_generic.png",
	"IDR_PROFILE_AVATAR_1": "avatar_generic_aqua.png",
	"IDR_PROFILE_AVATAR_2": "avatar_generic_blue.png",
	"IDR_PROFILE_AVATAR_3": "avatar_generic_green.png",
	"IDR_PROFILE_AVATAR_4": "avatar_generic_orange.png",
	"IDR_PROFILE_AVATAR_5": "avatar_generic_purple.png",
	"IDR_PROFILE_AVATAR_6": "avatar_generic_red.png",
	"IDR_PROFILE_AVATAR_7": "avatar_generic_yellow.png",
	"IDR_PROFILE_AVATAR_8": "avatar_secret_agent.png",
	"IDR_PROFILE_AVATAR_9": "avatar_superhero.png",
	"IDR_PROFILE_AVATAR_10": "avatar_volley_ball.png",
	"IDR_PROFILE_AVATAR_11": "avatar_businessman.png",
	"IDR_PROFILE_AVATAR_12": "avatar_ninja.png",
	"IDR_PROFILE_AVATAR_13": "avatar_alien.png",
	"IDR_PROFILE_AVATAR_14": "avatar_awesome.png",
	"IDR_PROFILE_AVATAR_15": "avatar_flower.png",
	"IDR_PROFILE_AVATAR_16": "avatar_pizza.png",
	"IDR_PROFILE_AVATAR_17": "avatar_soccer.png",
	"IDR_PROFILE_AVATAR_18": "avatar_burger.png",
	"IDR_PROFILE_AVATAR_19": "avatar_cat.png",
	"IDR_PROFILE_AVATAR_20": "avatar_cupcake.png",
	"IDR_PROFILE_AVATAR_21": "avatar_dog.png",
	"IDR_PROFILE_AVATAR_22": "avatar_horse.png",
	"IDR_PROFILE_AVATAR_23": "avatar_margarita.png",
	"IDR_PROFILE_AVATAR_24": "avatar_note.png",
	"IDR_PROFILE_AVATAR_25": "avatar_sun_cloud.png",
	"IDR_PROFILE_AVATAR_26": "avatar_placeholder.png",
}

type Profile struct {
	Name				string
	IsActive			bool
	IsVirtual			bool
	IconURL				string
	AvatarURL			string		`json:"avatar_icon"`
	CustomAvatarURL		string		`json:"gaia_picture_file_name"`
	IsDefaultAvatar 	bool		`json:"is_using_default_avatar"`
	DisplayName			string		`json:"gaia_name"`
	UserName 			string		`json:"name"`
	UserEmail 			string		`json:"user_name"`
}

var AutoProfile = Profile{
	Name: "",
	DisplayName: "Auto",
	IsActive: true,
	IsVirtual: true,
}

func (p Profile) AvatarIconURL(browser Browser, profileName string) string {
	var iconURL string

	if p.IsDefaultAvatar {
		name := avatarResources[strings.Replace(p.AvatarURL, avatarResourcePrefix, "", -1)]

		iconURL = browser.FullPathFor("Avatars", name)
	} else {
		iconURL = browser.FullPathFor(profileName, p.CustomAvatarURL)
	}

	if stat, err := os.Stat(iconURL); os.IsNotExist(err) || stat.IsDir() {
		iconURL = browser.IconURL
	}

	return iconURL
}
