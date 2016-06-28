package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWorkflow_MetadataFor(t *testing.T) {
	assert := assert.New(t)
	browser := Browser{ID: "1"}
	browsers := BrowserSlice{
		browser,
	}
	workflow := NewWorkflow()
	metadata := workflow.MetadataFor(browsers)

	assert.NotEmpty(metadata[browser.ID])
}
