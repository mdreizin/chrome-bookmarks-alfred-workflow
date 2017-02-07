package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBrowserRepository(t *testing.T) {
	test := assert.New(t)
	browserRepository := NewBrowserRepository("browser.yml")

	test.IsType(&YmlBrowserRepository{}, browserRepository)
}
