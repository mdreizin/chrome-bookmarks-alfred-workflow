package workflows

type Workflow struct {
	BookmarkListID string `yaml:"bookmarkListId"`
	BookmarkOpenID string `yaml:"bookmarkOpenId"`
	ProfileListID  string `yaml:"profileListId"`
	ProfileSetID   string `yaml:"profileSetId"`
}

type WorkflowMetadata map[string]Workflow
