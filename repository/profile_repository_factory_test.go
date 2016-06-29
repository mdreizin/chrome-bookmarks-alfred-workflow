package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProfileRepository(t *testing.T) {
	assert := assert.New(t)
	profileRepository := NewProfileRepository("Profiles.json")

	assert.IsType(&JsonProfileRepository{}, profileRepository)
}
