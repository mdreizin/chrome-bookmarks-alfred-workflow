package model

import "github.com/satori/go.uuid"

const WorkflowBundleID string = "com.mdreizin.chrome.bookmarks"
const WorkflowName string = "Chrome Bookmarks"
const WorkflowDescription string = "Chrome/Canary/Chromium bookmarks search workflow for Alfred"
const WorkflowVersion string = "0.1.0"
const WorkflowAuthor string = "Marat Dreizin"
const WorkflowCategory string = "Productivity"
const WorkflowEmail string = "marat.dreizin@gmail.com"
const WorkflowURL string = "http://mdreizin.github.io/chrome-bookmarks-alfred-workflow"
const WorkflowAlfredName string = "chrome-bookmarks.alfredworkflow"

type Workflow struct {
	BundleID		string
	Version			string
	Name			string
	Description		string
	Category		string
	Author			string
	URL				string
	Readme			string
}

func NewWorkflow() Workflow {
	return Workflow{
		BundleID: WorkflowBundleID,
		Version: WorkflowVersion,
		Name: WorkflowName,
		Description: WorkflowDescription,
		Category: WorkflowCategory,
		Author: WorkflowAuthor,
		URL: WorkflowURL,
		Readme: "",
	}
}

func (w Workflow) MetadataFor(browsers BrowserSlice) map[string]WorkflowMetadata {
	metadata := map[string]WorkflowMetadata{}

	for _, v := range browsers {
		metadata[v.ID] = WorkflowMetadata{
			BookmarkListID: uuid.NewV4().String(),
			BookmarkOpenID: uuid.NewV4().String(),
			ProfileListID: uuid.NewV4().String(),
			ProfileSetID: uuid.NewV4().String(),
		}
	}

	return metadata
}
