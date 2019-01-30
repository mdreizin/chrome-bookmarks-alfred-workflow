package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/workflows"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	templateExt   = ".gohtml"
	versionPrefix = "v"
)

var (
	version              = "dev"
	workflowFile         string
	workflowTemplateFile string
	browserFile          string
	outDir               string
	assetDir             string
)

func init() {
	flag.StringVar(&workflowFile, "workflow-file", "", "")
	flag.StringVar(&workflowTemplateFile, "workflow-tmpl-file", "", "")
	flag.StringVar(&browserFile, "browser-file", "", "")
	flag.StringVar(&assetDir, "asset-dir", "", "")
	flag.StringVar(&outDir, "out-dir", "", "")
}

func run() error {
	workflowService := workflows.WorkflowService(&workflows.DefaultWorkflowService{
		BrowserService: browsers.BrowserService(&browsers.DefaultBrowserService{
			BrowserRepository: &browsers.YmlBrowserRepository{
				Filename: browserFile,
			},
		}),
		WorkflowRepository: workflows.WorkflowRepository(&workflows.YmlWorkflowRepository{
			Filename: workflowFile,
		}),
	})

	browserSlice, err := workflowService.GetBrowsers()

	if err != nil {
		return err
	}

	workflowMetadata, err := workflowService.GetWorkflowMetadata()

	var cmd []string

	for k, v := range workflowMetadata {
		if b, err := browserSlice.FindByName(k); err == nil {
			cmd = append(cmd, fmt.Sprintf("cp %s %s", path.Join(assetDir, b.IconURL), path.Join(outDir, v.BookmarkListID+filepath.Ext(b.IconURL))))
			cmd = append(cmd, fmt.Sprintf("cp %s %s", path.Join(assetDir, b.IconURL), path.Join(outDir, v.ProfileListID+filepath.Ext(b.IconURL))))
		}
	}

	err = exec.Command("/bin/sh", "-c", strings.Join(cmd, "; ")).Run()

	if err != nil {
		return err
	}

	content, err := ioutil.ReadFile(workflowTemplateFile)

	if err != nil {
		return err
	}

	tmpl, err := template.New("Workflow").Funcs(sprig.TxtFuncMap()).Parse(string(content))

	if err != nil {
		return err
	}

	data := struct {
		Version  string
		Browsers browsers.BrowserSlice
		Metadata workflows.WorkflowMetadata
	}{
		strings.TrimPrefix(version, versionPrefix),
		browserSlice,
		workflowMetadata,
	}

	wr := bytes.NewBufferString("")
	err = tmpl.Execute(wr, data)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(path.Join(outDir, strings.TrimSuffix(filepath.Base(workflowTemplateFile), templateExt)), wr.Bytes(), 0644)
}

func main() {
	flag.Parse()

	err := run()

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
