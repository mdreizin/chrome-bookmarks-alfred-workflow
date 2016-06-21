package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewBrowserRepository(t *testing.T) {
	assert := assert.New(t)
	browserRepository := NewBrowserRepository("browser.yml")

	assert.IsType(&YmlBrowserRepository{}, browserRepository)
}
