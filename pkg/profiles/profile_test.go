package profiles

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfile_AvatarIconURL(t *testing.T) {
	test := assert.New(t)
	browser := browsers.Browser{Path: "test", IconURL: "img/chrome.png"}
	profile := Profile{
		AvatarURL:       "chrome://theme/IDR_PROFILE_AVATAR_12",
		CustomAvatarURL: "avatar.png",
		IsDefaultAvatar: false,
	}

	test.Equal("img/chrome.png", profile.AvatarIconURL(&browser, DefaultProfileName))

	profile.IsDefaultAvatar = true

	test.Equal("test/Avatars/avatar_ninja.png", profile.AvatarIconURL(&browser, DefaultProfileName))
}
