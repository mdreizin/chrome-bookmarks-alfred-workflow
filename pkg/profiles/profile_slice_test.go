package profiles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfileSlice_Match(t *testing.T) {
	test := assert.New(t)
	profileSlice := ProfileSlice{
		&Profile{
			Name:      DefaultProfileName,
			UserName:  "User 1",
			UserEmail: "user1@gmail.com",
		},
		&Profile{
			Name:      "User2",
			UserName:  "User 2",
			UserEmail: "user1@gmail.com",
		},
	}

	test.Len(profileSlice.Match("User 1"), 1)
	test.Len(profileSlice.Match("User 2"), 1)
	test.Len(profileSlice.Match("User"), 2)
	test.Len(profileSlice.Match("gmail.com"), 2)
}

func TestProfileSlice_Sort(t *testing.T) {
	test := assert.New(t)
	profileSlice := ProfileSlice{
		&Profile{
			Name:      "User2",
			UserName:  "User 2",
			UserEmail: "user2@gmail.com",
		},
		&Profile{
			Name:      DefaultProfileName,
			UserName:  "User 1",
			UserEmail: "user1@gmail.com",
		},
	}.Sort()

	test.Equal(&Profile{
		Name:      DefaultProfileName,
		UserName:  "User 1",
		UserEmail: "user1@gmail.com",
	}, profileSlice[0])
}

func TestProfileSlice_FirstActive(t *testing.T) {
	test := assert.New(t)
	profileSlice := ProfileSlice{
		&Profile{
			Name:      "User2",
			UserName:  "User 2",
			UserEmail: "user2@gmail.com",
			IsActive:  true,
		},
		&Profile{
			Name:      DefaultProfileName,
			UserName:  "User 1",
			UserEmail: "user1@gmail.com",
			IsActive:  false,
		},
	}
	profile, err := profileSlice.FirstActive()

	test.NoError(err)
	test.Equal(&Profile{
		Name:      "User2",
		UserName:  "User 2",
		UserEmail: "user2@gmail.com",
		IsActive:  true,
	}, profile)

	profileSlice[0].IsActive = false

	profile, err = profileSlice.FirstActive()

	test.Error(err)
}
