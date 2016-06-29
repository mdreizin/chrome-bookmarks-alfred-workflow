package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfileSlice_Match(t *testing.T) {
	assert := assert.New(t)
	profiles := ProfileSlice{
		Profile{
			Name:      "Default",
			UserName:  "User 1",
			UserEmail: "user1@gmail.com",
		},
		Profile{
			Name:      "User2",
			UserName:  "User 2",
			UserEmail: "user1@gmail.com",
		},
	}

	assert.Len(profiles.Match("User 1"), 1)
	assert.Len(profiles.Match("User 2"), 1)
	assert.Len(profiles.Match("User"), 2)
	assert.Len(profiles.Match("gmail.com"), 2)
}

func TestProfileSlice_Sort(t *testing.T) {
	assert := assert.New(t)
	profiles := ProfileSlice{
		Profile{
			Name:      "User2",
			UserName:  "User 2",
			UserEmail: "user2@gmail.com",
		},
		Profile{
			Name:      "Default",
			UserName:  "User 1",
			UserEmail: "user1@gmail.com",
		},
	}.Sort()

	assert.Equal(Profile{
		Name:      "Default",
		UserName:  "User 1",
		UserEmail: "user1@gmail.com",
	}, profiles[0])
}

func TestProfileSlice_FirstActive(t *testing.T) {
	assert := assert.New(t)
	profiles := ProfileSlice{
		Profile{
			Name:      "User2",
			UserName:  "User 2",
			UserEmail: "user2@gmail.com",
			IsActive:  true,
		},
		Profile{
			Name:      "Default",
			UserName:  "User 1",
			UserEmail: "user1@gmail.com",
			IsActive:  false,
		},
	}
	profile, err := profiles.FirstActive()

	assert.NoError(err)
	assert.Equal(Profile{
		Name:      "User2",
		UserName:  "User 2",
		UserEmail: "user2@gmail.com",
		IsActive:  true,
	}, profile)

	profiles[0].IsActive = false

	profile, err = profiles.FirstActive()

	assert.Error(err)
}
