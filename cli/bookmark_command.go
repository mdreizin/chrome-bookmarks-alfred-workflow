package main

import (
	"fmt"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/service"
	"github.com/ruedap/go-alfred"
	"github.com/urfave/cli"
)

func bookmarkCommand(flags []cli.Flag) cli.Command {
	return cli.Command{
		Name:  "bookmarks",
		Flags: flags,
		Action: func(c *cli.Context) error {
			name := c.Args().First()
			query := c.String("query")
			bookmarkService := service.NewBookmarkService(extractConfig(c))
			browsers, _ := bookmarkService.GetBrowsers()
			browser, _ := browsers.FindByName(name)
			bookmarks, err := bookmarkService.GetBookmarks(browser, query)

			var xml string

			if err != nil {
				xml = alfred.ErrorXML(c.App.Name, fmt.Sprintf("Error: %s", err.Error()), err.Error())
			} else {
				r := alfred.NewResponse()

				for _, bookmark := range bookmarks {
					r.AddItem(&alfred.ResponseItem{
						Valid:    true,
						UID:      bookmark.URL,
						Title:    bookmark.Name,
						Subtitle: bookmark.URL,
						Arg:      bookmark.URL,
						Icon:     bookmark.IconURL,
					})
				}

				xml, _ = r.ToXML()
			}

			fmt.Print(xml)

			return err
		},
	}
}
