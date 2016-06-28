package main

import (
	"fmt"
	"path"
	"os"
	"github.com/urfave/cli"
	"github.com/ruedap/go-alfred"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/service"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/stringutil"
)

var cwd, _ = os.Getwd()

const (
	browserFile = "browser-file"
	profileFile = "profile-file"
	bookmarkFile = "bookmark-file"
)

func extractConfig(c *cli.Context) map[string]string {
	return map[string]string {
		browserFile: c.String(browserFile),
		profileFile: c.String(profileFile),
		bookmarkFile: c.String(bookmarkFile),
	}
}

func newApp() *cli.App {
	flags := []cli.Flag{
		cli.StringFlag{
			Name: browserFile,
			Value: path.Join(cwd, "browser.yml"),
			EnvVar: stringutil.KebabCase(browserFile),
		},
		cli.StringFlag{
			Name: profileFile,
			Value: "Local State",
			EnvVar: stringutil.KebabCase(profileFile),
		},
		cli.StringFlag{
			Name: bookmarkFile,
			Value: "Bookmarks",
			EnvVar: stringutil.KebabCase(bookmarkFile),
		},
		cli.StringFlag{
			Name: "query",
		},
	}

	app := cli.NewApp()
	app.Name = "cli"
	app.Version = model.WorkflowVersion
	app.Usage = model.WorkflowDescription
	app.Author = model.WorkflowAuthor
	app.Email = model.WorkflowEmail
	app.Commands = cli.Commands{cli.Command{
		Name: "bookmarks",
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
					title := ""
					subtitle := ""

					if profile.IsVirtual {
						title = profile.DisplayName
					} else if profile.Name != "" {
						title = profile.Name
					}

					if profile.IsActive {
						title = fmt.Sprintf("%s (Active)", title)
					}

					if !profile.IsVirtual && profile.Name != "" {
						subtitle = browser.PathFor(profile.Name)
					}

					r.AddItem(&alfred.ResponseItem{
						Valid: true,
						UID: profile.Name,
						Title: title,
						Subtitle: subtitle,
						Arg: profile.Name,
						Icon: profile.IconURL,
					})
				}

				xml, _ = r.ToXML()
			}

			fmt.Print(xml)

			return err
		},
	}, cli.Command{
		Name:  "select-profile",
		Flags: flags,
		Action: func(c *cli.Context) error {
			name := c.Args().First()
			query := c.String("query")
			bookmarkService := service.NewBookmarkService(extractConfig(c))
			browsers, _ := bookmarkService.GetBrowsers()
			browser, _ := browsers.FindByName(name)

			browser.ProfileName = query

			err := bookmarkService.UpdateBrowser(browser)

			if err != nil {
				fmt.Print(fmt.Sprintf("Unable to update profile %s", err))
			} else {
				fmt.Print("Profile has been successfully updated")
			}

			return err
		},
	}}

	return app
}
