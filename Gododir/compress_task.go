package main

import (
	"fmt"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	do "gopkg.in/godo.v2"
)

func compressTask(dest string) func(*do.Context) {
	return func(c *do.Context) {
		c.Bash(fmt.Sprintf("cd %[2]s && zip -rX %[1]s ./* -x %[1]s &> /dev/null", model.WorkflowAlfredName, dest))
	}
}
