package workflows

type WorkflowRepository interface {
	GetWorkflowMetadata() (WorkflowMetadata, error)
}
