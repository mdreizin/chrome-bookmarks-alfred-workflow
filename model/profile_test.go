package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestProfile_AvatarIconURL(t *testing.T) {
	assert := assert.New(t)
	browser := Browser{Path: "testdata"}
	profile := Profile{
		AvatarURL: "chrome://theme/IDR_PROFILE_AVATAR_12",
		CustomAvatarURL: "avatar.png",
		IsDefaultAvatar: false,
	}

	assert.Equal("testdata/Default/avatar.png", profile.AvatarIconURL(browser, "Default"))

	profile.IsDefaultAvatar = true

	assert.Equal("testdata/Avatars/avatar_ninja.png", profile.AvatarIconURL(browser, "Default"))
}
