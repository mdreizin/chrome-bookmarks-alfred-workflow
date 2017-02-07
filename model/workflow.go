package model

import "github.com/satori/go.uuid"

const (
	WorkflowAppName     string = "chrome-bookmarks"
	WorkflowBundleID    string = "com.mdreizin.chrome.bookmarks"
	WorkflowName        string = "Chrome Bookmarks"
	WorkflowDescription string = "Chrome/Canary/Chromium bookmarks search workflow for Alfred"
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
	AppName     string
}

func NewWorkflow() Workflow {
	return Workflow{
		BundleID:    WorkflowBundleID,
		Version:     "dev",
		Name:        WorkflowName,
		Description: WorkflowDescription,
		Category:    WorkflowCategory,
		Author:      WorkflowAuthor,
		URL:         WorkflowURL,
		Readme:      "",
		AppName:     WorkflowAppName,
	}
}

func (w Workflow) MetadataFor(browsers BrowserSlice) map[string]WorkflowMetadata {
	metadata := map[string]WorkflowMetadata{}

	for _, v := range browsers {
		metadata[v.ID] = WorkflowMetadata{
			BookmarkListID: v.ID,
			BookmarkOpenID: uuid.NewV4().String(),
			ProfileListID:  uuid.NewV4().String(),
			ProfileSetID:   uuid.NewV4().String(),
		}
	}

	return metadata
}
