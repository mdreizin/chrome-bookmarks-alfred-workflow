package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewProfileRepository(t *testing.T) {
	assert := assert.New(t)
	profileRepository := NewProfileRepository("Profiles.json")

	assert.IsType(&JsonProfileRepository{}, profileRepository)
}
