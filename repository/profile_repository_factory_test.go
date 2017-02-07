package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProfileRepository(t *testing.T) {
	test := assert.New(t)
	profileRepository := NewProfileRepository("Profiles.json")

	test.IsType(&JsonProfileRepository{}, profileRepository)
}
