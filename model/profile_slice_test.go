package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfileSlice_Match(t *testing.T) {
	test := assert.New(t)
	profiles := ProfileSlice{
		Profile{
			Name:      DefaultProfileName,
			UserName:  "User 1",
			UserEmail: "user1@gmail.com",
		},
		Profile{
			Name:      "User2",
			UserName:  "User 2",
			UserEmail: "user1@gmail.com",
		},
	}

	test.Len(profiles.Match("User 1"), 1)
	test.Len(profiles.Match("User 2"), 1)
	test.Len(profiles.Match("User"), 2)
	test.Len(profiles.Match("gmail.com"), 2)
}

func TestProfileSlice_Sort(t *testing.T) {
	test := assert.New(t)
	profiles := ProfileSlice{
		Profile{
			Name:      "User2",
			UserName:  "User 2",
			UserEmail: "user2@gmail.com",
		},
		Profile{
			Name:      DefaultProfileName,
			UserName:  "User 1",
			UserEmail: "user1@gmail.com",
		},
	}.Sort()

	test.Equal(Profile{
		Name:      DefaultProfileName,
		UserName:  "User 1",
		UserEmail: "user1@gmail.com",
	}, profiles[0])
}

func TestProfileSlice_FirstActive(t *testing.T) {
	test := assert.New(t)
	profiles := ProfileSlice{
		Profile{
			Name:      "User2",
			UserName:  "User 2",
			UserEmail: "user2@gmail.com",
			IsActive:  true,
		},
		Profile{
			Name:      DefaultProfileName,
			UserName:  "User 1",
			UserEmail: "user1@gmail.com",
			IsActive:  false,
		},
	}
	profile, err := profiles.FirstActive()

	test.NoError(err)
	test.Equal(Profile{
		Name:      "User2",
		UserName:  "User 2",
		UserEmail: "user2@gmail.com",
		IsActive:  true,
	}, profile)

	profiles[0].IsActive = false

	profile, err = profiles.FirstActive()

	test.Error(err)
}
