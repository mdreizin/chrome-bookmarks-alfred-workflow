package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBrowserRepository(t *testing.T) {
	assert := assert.New(t)
	browserRepository := NewBrowserRepository("browser.yml")

	assert.IsType(&YmlBrowserRepository{}, browserRepository)
}
