package workflows

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYmlWorkflowRepository_GetWorkflowMetadata(t *testing.T) {
	test := assert.New(t)
	workflowRepository := WorkflowRepository(&YmlWorkflowRepository{
		Filename: "test/workflow.yml",
	})
	workflowMetadata, err := workflowRepository.GetWorkflowMetadata()

	test.NoError(err)
	test.Len(workflowMetadata, 1)
}
