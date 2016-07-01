package main

import (
	"fmt"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/service"
	"github.com/ruedap/go-alfred"
	"github.com/urfave/cli"
)

func profileCommand(flags []cli.Flag) cli.Command {
	return cli.Command{
		Name:  "profiles",
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
						Valid:    true,
						UID:      profile.Name,
						Title:    title,
						Subtitle: subtitle,
						Arg:      profile.Name,
						Icon:     profile.IconURL,
					})
				}

				xml, _ = r.ToXML()
			}

			fmt.Print(xml)

			return err
		},
	}
}
