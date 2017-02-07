package main

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/stringutil"
	do "gopkg.in/godo.v2"
	"io/ioutil"
	"path"
	"text/template"
)

func generateTask(src string, dest string, workflow model.Workflow, browsers model.BrowserSlice, metadata map[string]model.WorkflowMetadata) func(*do.Context) {
	return func(c *do.Context) {
		ensureDir(c, dest)

		if content, err := ioutil.ReadFile(path.Join(src, "info.plist")); err == nil {
			if t, err := template.New("Workflow").Funcs(sprig.TxtFuncMap()).Parse(string(content)); err == nil {
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

				wr := bytes.NewBufferString("")

				if err := t.Execute(wr, data); err == nil {
					ioutil.WriteFile(path.Join(dest, "info.plist"), wr.Bytes(), 0644)
				}
			}
		}
	}
}
