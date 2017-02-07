package main

import (
	"fmt"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	do "gopkg.in/godo.v2"
	"os"
	"path"
)

func copyTask(src string, dest string, browsers model.BrowserSlice, metadata map[string]model.WorkflowMetadata) func(*do.Context) {
	return func(c *do.Context) {
		ensureDir(c, dest)

		files := map[string][]string{
			path.Join(src, "bin"):                                        {dest},
			path.Join(src, "img"):                                        {dest},
			path.Join(src, "browser.yml"):                                {dest},
			path.Join(src, "icon.png"):                                   {dest},
			path.Join(os.Getenv("GOPATH"), "bin", model.WorkflowAppName): {path.Join(dest, "bin")},
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
	}
}
