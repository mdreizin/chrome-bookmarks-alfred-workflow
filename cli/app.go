package main

import (
	"fmt"
	"path"
	"os"
	"github.com/urfave/cli"
	"github.com/ruedap/go-alfred"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/workflow"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/service"
)

var cwd, _ = os.Getwd()

func extractConfig(c *cli.Context) map[string]string {
	return map[string]string {
		"browser-file": c.String("browser-file"),
		"profile-file": c.String("profile-file"),
		"bookmark-file": c.String("bookmark-file"),
	}
}

func newApp() *cli.App {
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "browser-file, brf",
			Value: path.Join(cwd, "browser.yml"),
			EnvVar: "BROWSER_FILENAME",
		},
		cli.StringFlag{
			Name: "profile-file, prf",
			Value: "Local State",
			EnvVar: "PROFILE_FILENAME",
		},
		cli.StringFlag{
			Name: "bookmark-file, bof",
			Value: "Bookmarks",
			EnvVar: "BOOKMARK_FILENAME",
		},
		cli.StringFlag{
			Name: "query, q",
		},
	}

	app := cli.NewApp()
	app.Name = "cli"
	app.Version = workflow.Version
	app.Usage = workflow.Description
	app.Author = workflow.Author
	app.Email = workflow.Email
	app.Commands = cli.Commands{cli.Command{
		Name: "bookmarks",
		Aliases: []string{"b"},
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
						Valid: true,
						UID: bookmark.URL,
						Title: bookmark.Name,
						Subtitle: bookmark.URL,
						Arg: bookmark.URL,
						Icon: bookmark.IconURL,
					})
				}

				xml, _ = r.ToXML()
			}

			fmt.Print(xml)

			return err
		},
	}, cli.Command{
		Name: "profiles",
		Aliases: []string{"p"},
		Flags: flags,
		Action: func(c *cli.Context) error {
			name := c.Args().First()
			query := c.String("query")
			bookmarkService := service.NewBookmarkService(extractConfig(c))
			browsers, _ := bookmarkService.GetBrowsers()
			browser, _ := browsers.FindByName(name)
			profiles, err := bookmarkService.GetProfiles(browser, query)

			var xml string

			if err != nil {
				xml = alfred.ErrorXML(c.App.Name, fmt.Sprintf("Error: %s", err.Error()), err.Error())
			} else {
				r := alfred.NewResponse()

				for _, profile := range profiles {
					title := profile.Name

					if profile.IsActive {
						title = fmt.Sprintf("%s (Active)", title)
					}

					r.AddItem(&alfred.ResponseItem{
						Valid: true,
						UID: profile.Name,
						Title: title,
						Subtitle: browser.PathFor(browser.Name, profile.Name),
						Arg: profile.Name,
						Icon: profile.IconURL,
					})
				}

				xml, _ = r.ToXML()
			}

			fmt.Print(xml)

			return err
		},
	}}

	return app
}
