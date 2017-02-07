package main

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/stringutil"
	do "gopkg.in/godo.v2"
	"os"
	"path"
	"text/template"
)

func generateTask(src string, dest string, workflow model.Workflow, browsers model.BrowserSlice, metadata map[string]model.WorkflowMetadata) func(*do.Context) {
	return func(c *do.Context) {
		ensureDir(c, dest)

		t, _ := template.ParseFiles(path.Join(src, "info.plist"))
		f, _ := os.Create(path.Join(dest, "info.plist"))
		version := c.Args.AsString("version")

		workflow.Version = stringutil.VersionWithoutPrefix(version)

		data := struct {
			Workflow model.Workflow
			Browsers model.BrowserSlice
			Metadata map[string]model.WorkflowMetadata
		}{
			workflow,
			browsers,
			metadata,
		}

		t.Execute(f, data)
		f.Close()
	}
}
