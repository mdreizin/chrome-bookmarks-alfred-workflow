package main

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/service"
	do "gopkg.in/godo.v2"
	"path"
)

func tasks(p *do.Project) {
	src := "./workflow"
	dest := "./build"
	config := map[string]string{
		"browser-file": path.Join(src, "browser.yml"),
	}
	bookmarkService := service.NewBookmarkService(config)
	browsers, _ := bookmarkService.GetBrowsers()
	workflow := model.NewWorkflow()
	metadata := workflow.GenerateMetadata(browsers)

	p.Task("default", do.Series{"clean", do.Parallel{"copy", "generate"}, "compress"}, nil)
	p.Task("clean", nil, cleanTask(dest))
	p.Task("copy", nil, copyTask(src, dest, browsers, metadata))
	p.Task("generate", nil, generateTask(src, dest, workflow, browsers, metadata))
	p.Task("compress", nil, compressTask(dest))
}
