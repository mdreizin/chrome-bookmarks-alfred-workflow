package main

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/stringutil"
	"github.com/urfave/cli"
	"os"
	"path"
)

var cwd, _ = os.Getwd()

const (
	browserFile  = "browser-file"
	profileFile  = "profile-file"
	bookmarkFile = "bookmark-file"
)

func extractConfig(c *cli.Context) map[string]string {
	return map[string]string{
		browserFile:  c.String(browserFile),
		profileFile:  c.String(profileFile),
		bookmarkFile: c.String(bookmarkFile),
	}
}

func newApp() *cli.App {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:   browserFile,
			Value:  path.Join(cwd, "browser.yml"),
			EnvVar: stringutil.KebabCase(browserFile),
		},
		cli.StringFlag{
			Name:   profileFile,
			Value:  "Local State",
			EnvVar: stringutil.KebabCase(profileFile),
		},
		cli.StringFlag{
			Name:   bookmarkFile,
			Value:  "Bookmarks",
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
	app.Commands = cli.Commands{
		bookmarkCommand(flags),
		profileCommand(flags),
		selectProfileCommand(flags),
	}

	return app
}
