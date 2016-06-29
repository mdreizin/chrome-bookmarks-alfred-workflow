package repository

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"regexp"
	"strings"
)

type YmlBrowserRepository struct {
	filename string
}

func (r *YmlBrowserRepository) GetBrowsers() (model.BrowserSlice, error) {
	browsers := model.BrowserSlice{}
	bytes, err := ioutil.ReadFile(r.filename)

	if err != nil {
		return browsers, err
	}

	err = yaml.Unmarshal(bytes, &browsers)

	if err == nil {
		re := regexp.MustCompile(`(\\ )`)

		for i, v := range browsers {
			browsers[i].ProfileName = re.ReplaceAllString(v.ProfileName, " ")
		}
	}

	return browsers, err
}

func (r *YmlBrowserRepository) UpdateBrowser(b model.Browser) error {
	browsers, err := r.GetBrowsers()

	if err != nil {
		return err
	}

	i := browsers.FindIndex(func(v model.Browser) bool { return strings.EqualFold(b.Name, v.Name) })

	if i >= 0 {
		browsers[i] = b

		bytes, err := yaml.Marshal(browsers)

		if err != nil {
			return err
		}

		err = ioutil.WriteFile(r.filename, bytes, 0644)
	}

	return err
}
