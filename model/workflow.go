package model

import "github.com/satori/go.uuid"

const (
	WorkflowBundleID    string = "com.mdreizin.chrome.bookmarks"
	WorkflowName        string = "Chrome Bookmarks"
	WorkflowDescription string = "Chrome/Canary/Chromium bookmarks search workflow for Alfred"
	WorkflowVersion     string = "0.1.2"
	WorkflowAuthor      string = "Marat Dreizin"
	WorkflowCategory    string = "Productivity"
	WorkflowEmail       string = "marat.dreizin@gmail.com"
	WorkflowURL         string = "http://mdreizin.github.io/chrome-bookmarks-alfred-workflow"
	WorkflowAlfredName  string = "chrome-bookmarks.alfredworkflow"
)

type Workflow struct {
	BundleID    string
	Version     string
	Name        string
	Description string
	Category    string
	Author      string
	URL         string
	Readme      string
}

func NewWorkflow() Workflow {
	return Workflow{
		BundleID:    WorkflowBundleID,
		Version:     WorkflowVersion,
		Name:        WorkflowName,
		Description: WorkflowDescription,
		Category:    WorkflowCategory,
		Author:      WorkflowAuthor,
		URL:         WorkflowURL,
		Readme:      "",
	}
}

func (w Workflow) MetadataFor(browsers BrowserSlice) map[string]WorkflowMetadata {
	metadata := map[string]WorkflowMetadata{}

	for _, v := range browsers {
		metadata[v.ID] = WorkflowMetadata{
			BookmarkListID: uuid.NewV4().String(),
			BookmarkOpenID: uuid.NewV4().String(),
			ProfileListID:  uuid.NewV4().String(),
			ProfileSetID:   uuid.NewV4().String(),
		}
	}

	return metadata
}
