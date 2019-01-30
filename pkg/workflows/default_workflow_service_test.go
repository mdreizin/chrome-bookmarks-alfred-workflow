package workflows

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/bookmarks"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/profiles"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultBookmarkService_GetBrowsers(t *testing.T) {
	test := assert.New(t)
	workflowService := WorkflowService(&DefaultWorkflowService{
		BrowserService: browsers.BrowserService(&browsers.DefaultBrowserService{
			BrowserRepository: browsers.BrowserRepository(&browsers.YmlBrowserRepository{
				Filename: "test/browser.yml",
			}),
		}),
	})
	browserSlice, err := workflowService.GetBrowsers()

	test.NoError(err)
	test.Len(browserSlice, 1)
}

func TestDefaultBookmarkService_GetProfiles(t *testing.T) {
	test := assert.New(t)
	workflowService := WorkflowService(&DefaultWorkflowService{
		BrowserService: browsers.BrowserService(&browsers.DefaultBrowserService{
			BrowserRepository: browsers.BrowserRepository(&browsers.YmlBrowserRepository{
				Filename: "test/browser.yml",
			}),
		}),
		ProfileService: profiles.ProfileService(&profiles.DefaultProfileService{
			ProfileRepository: profiles.ProfileRepository(&profiles.JsonProfileRepository{
				Filename: "Profiles.json",
			}),
		}),
	})
	browserSlice, _ := workflowService.GetBrowsers()
	browser, _ := browserSlice.FindByName("chrome")
	profileSlice, err := workflowService.GetProfiles(browser, "")

	test.NoError(err)
	test.Len(profileSlice, 2)

	profileSlice, err = workflowService.GetProfiles(browser, profiles.DefaultProfileName)

	test.NoError(err)
	test.Len(profileSlice, 1)
}

func TestDefaultBookmarkService_GetBookmarks(t *testing.T) {
	test := assert.New(t)
	workflowService := WorkflowService(&DefaultWorkflowService{
		BrowserService: browsers.BrowserService(&browsers.DefaultBrowserService{
			BrowserRepository: browsers.BrowserRepository(&browsers.YmlBrowserRepository{
				Filename: "test/browser.yml",
			}),
		}),
		ProfileService: profiles.ProfileService(&profiles.DefaultProfileService{
			ProfileRepository: profiles.ProfileRepository(&profiles.JsonProfileRepository{
				Filename: "Profiles.json",
			}),
		}),
		BookmarkService: bookmarks.BookmarkService(&bookmarks.DefaultBookmarkService{
			BookmarkRepository: bookmarks.BookmarkRepository(&bookmarks.JsonBookmarkRepository{
				Filename: "Bookmarks.json",
			}),
		}),
	})
	browserSlice, _ := workflowService.GetBrowsers()
	browser, _ := browserSlice.FindByName("chrome")
	bookmarkSlice, err := workflowService.GetBookmarks(browser, "")

	test.NoError(err)
	test.Len(bookmarkSlice, 1)

	bookmarkSlice, err = workflowService.GetBookmarks(browser, "alfred")

	test.NoError(err)
	test.Len(bookmarkSlice, 1)
}

func TestDefaultBookmarkService_UpdateBrowser(t *testing.T) {
	test := assert.New(t)
	workflowService := WorkflowService(&DefaultWorkflowService{
		BrowserService: browsers.BrowserService(&browsers.DefaultBrowserService{
			BrowserRepository: browsers.BrowserRepository(&browsers.YmlBrowserRepository{
				Filename: "test/browser.yml",
			}),
		}),
		ProfileService: profiles.ProfileService(&profiles.DefaultProfileService{
			ProfileRepository: profiles.ProfileRepository(&profiles.JsonProfileRepository{
				Filename: "Profiles.json",
			}),
		}),
		BookmarkService: bookmarks.BookmarkService(&bookmarks.DefaultBookmarkService{
			BookmarkRepository: bookmarks.BookmarkRepository(&bookmarks.JsonBookmarkRepository{
				Filename: "Bookmarks.json",
			}),
		}),
	})
	browserSlice, _ := workflowService.GetBrowsers()
	browser, _ := browserSlice.FindByName("chrome")

	browser.ProfileName = profiles.DefaultProfileName

	err := workflowService.UpdateBrowser(browser)

	browserSlice, _ = workflowService.GetBrowsers()
	browser, _ = browserSlice.FindByName("chrome")

	test.NoError(err)
	test.Equal(profiles.DefaultProfileName, browser.ProfileName)
}

func TestDefaultBookmarkService_GetWorkflowMetadata(t *testing.T) {
	test := assert.New(t)
	workflowService := WorkflowService(&DefaultWorkflowService{
		WorkflowRepository: WorkflowRepository(&YmlWorkflowRepository{
			Filename: "test/workflow.yml",
		}),
	})
	workflowMeta, err := workflowService.GetWorkflowMetadata()

	test.NoError(err)
	test.Len(workflowMeta, 1)
}
