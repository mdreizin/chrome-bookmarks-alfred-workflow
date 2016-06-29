package main

import (
	"fmt"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/service"
	do "gopkg.in/godo.v2"
	"os"
	"path"
	"text/template"
)

const (
	src  = "./workflow"
	dest = "./dist"
)

var (
	config = map[string]string{
		"browser-file": path.Join(src, "browser.yml"),
	}
	bookmarkService = service.NewBookmarkService(config)
)

func ensureDir(c *do.Context, dir string) {
	c.Bash(fmt.Sprintf("mkdir -p %s", dir))
}

func tasks(p *do.Project) {
	browsers, _ := bookmarkService.GetBrowsers()
	workflow := model.NewWorkflow()
	metadata := workflow.MetadataFor(browsers)

	p.Task("default", do.Series{"clean", do.Parallel{"files", "generate"}, "compress"}, nil)

	p.Task("clean", nil, func(c *do.Context) {
		c.Bash(fmt.Sprintf("rm -rf %s", dest))
	})

	p.Task("files", nil, func(c *do.Context) {
		ensureDir(c, dest)

		files := map[string][]string{
			path.Join(src, "bin"):                        {dest},
			path.Join(src, "img"):                        {dest},
			path.Join(src, "browser.yml"):                {dest},
			path.Join(src, "icon.png"):                   {dest},
			path.Join(os.Getenv("GOPATH"), "bin", "cli"): {path.Join(dest, "bin")},
		}

		for _, v := range browsers {
			m := metadata[v.ID]
			ext := path.Ext(v.IconURL)
			s := []string{
				path.Join(dest, m.BookmarkListID+ext),
				path.Join(dest, m.ProfileListID+ext),
			}

			files[path.Join(src, v.IconURL)] = s
		}

		for k, v := range files {
			for _, s := range v {
				if path.Ext(s) == "" {
					ensureDir(c, s)
				}

				c.Bash(fmt.Sprintf("rsync -r %s %s &> /dev/null", k, s))
			}
		}
	})

	p.Task("generate", nil, func(c *do.Context) {
		ensureDir(c, dest)

		t, _ := template.ParseFiles(path.Join(src, "info.plist"))
		f, _ := os.Create(path.Join(dest, "info.plist"))

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
	})

	p.Task("compress", nil, func(c *do.Context) {
		c.Bash(fmt.Sprintf("cd %[2]s && zip -rX %[1]s ./* -x %[1]s &> /dev/null", model.WorkflowAlfredName, dest))
	})
}
