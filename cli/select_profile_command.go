package main

import (
	"fmt"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/service"
	"github.com/urfave/cli"
)

func selectProfileCommand(flags []cli.Flag) cli.Command {
	return cli.Command{
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
	}
}
