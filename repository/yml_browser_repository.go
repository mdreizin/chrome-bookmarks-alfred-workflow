package repository

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
)

type YmlBrowserRepository struct {
	filename	string
}

func (r *YmlBrowserRepository) GetBrowsers() (model.BrowserSlice, error) {
	browsers := model.BrowserSlice{}
	bytes, err := ioutil.ReadFile(r.filename)

	if err != nil {
		return browsers, err
	}

	err = yaml.Unmarshal(bytes, &browsers)

	return browsers, err
}
