package main

import (
	"flag"
	"fmt"
	"github.com/deanishe/awgo"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/bookmarks"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/profiles"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/workflows"
	"os"
	"path"
	"strings"
)

var (
	cwd, _       = os.Getwd()
	wf           *aw.Workflow
	browserFile  string
	profileFile  string
	bookmarkFile string
	query        string
	browser      string
	command      string
)

func init() {
	flag.StringVar(&browserFile, "browser-file", path.Join(cwd, "browser.yml"), "")
	flag.StringVar(&profileFile, "profile-file", "Local State", "")
	flag.StringVar(&bookmarkFile, "bookmark-file", "Bookmarks", "")
	flag.StringVar(&query, "query", "", "")
	flag.StringVar(&browser, "browser", "", "")
	flag.StringVar(&command, "command", "", "")

	wf = aw.New()
}

func run() error {
	workflowService := workflows.WorkflowService(&workflows.DefaultWorkflowService{
		BrowserService: browsers.BrowserService(&browsers.DefaultBrowserService{
			BrowserRepository: browsers.BrowserRepository(&browsers.YmlBrowserRepository{
				Filename: browserFile,
			}),
		}),
		ProfileService: profiles.ProfileService(&profiles.DefaultProfileService{
			ProfileRepository: profiles.ProfileRepository(&profiles.JsonProfileRepository{
				Filename: profileFile,
			}),
		}),
		BookmarkService: bookmarks.BookmarkService(&bookmarks.DefaultBookmarkService{
			BookmarkRepository: bookmarks.BookmarkRepository(&bookmarks.JsonBookmarkRepository{
				Filename: bookmarkFile,
			}),
		}),
	})

	switch command {
	case "bookmark-list":
		browserSlice, err := workflowService.GetBrowsers()

		if err != nil {
			return err
		}

		b, err := browserSlice.FindByName(browser)

		if err != nil {
			return err
		}

		bookmarkSlice, err := workflowService.GetBookmarks(b, query)

		if err != nil {
			return err
		}

		if len(bookmarkSlice) > 0 {
			for _, bookmark := range bookmarkSlice {
				wf.NewItem(bookmark.Name).
					Subtitle(bookmark.URL).
					Arg(bookmark.URL).
					NewModifier(aw.ModCmd).
					Subtitle(strings.Join(bookmark.Path, " → "))
			}
		} else {
			wf.NewItem("No matching bookmarks").
				Subtitle("Try search bookmarks using a different query")
		}
	case "profile-list":
		browserSlice, err := workflowService.GetBrowsers()

		if err != nil {
			return err
		}

		browser, err := browserSlice.FindByName(browser)

		if err != nil {
			return err
		}

		profileSlice, err := workflowService.GetProfiles(browser, query)

		if err != nil {
			return err
		}

		if len(profileSlice) > 0 {
			for _, profile := range profileSlice {
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
					subtitle = browser.JoinPath(profile.Name)
				}

				wf.NewItem(title).
					Subtitle(subtitle).
					Valid(true).
					Icon(&aw.Icon{Value: profile.IconURL, Type: aw.IconTypeImage}).
					Arg(profile.Name)
			}
		} else {
			wf.NewItem("No matching profiles").
				Subtitle("Try search profiles using a different query")
		}
	case "profile-set":
		browserSlice, err := workflowService.GetBrowsers()

		if err != nil {
			return err
		}

		b, err := browserSlice.FindByName(browser)

		if err != nil {
			return err
		}

		b.ProfileName = query

		err = workflowService.UpdateBrowser(b)

		if err == nil {
			fmt.Printf("%s profile has been successfully updated", b.Description)
		} else {
			fmt.Printf("Unable to update %s profile", b.Description)
		}

		return nil
	}

	return nil
}

func main() {
	wf.Args()
	flag.Parse()

	err := run()

	if err != nil {
		wf.FatalError(err)
	} else {
		if len(wf.Feedback.Items) > 0 {
			wf.SendFeedback()
		}
	}
}
