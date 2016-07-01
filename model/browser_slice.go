package model

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type BrowserSlice []Browser

func (s BrowserSlice) Add(v Browser) BrowserSlice {
	return append(s, v)
}

func (s BrowserSlice) Match(query string) BrowserSlice {
	f := BrowserSlice{}

	re := regexp.MustCompile(fmt.Sprintf("(?i)%s", regexp.QuoteMeta(query)))

	for _, v := range s {
		if re.MatchString(v.Name) {
			f = f.Add(v)
		}
	}

	return f
}

func (s BrowserSlice) Find(f func(Browser) bool) (Browser, error) {
	for _, v := range s {
		if f(v) {
			return v, nil
		}
	}

	return Browser{}, errors.New("`Browser` is not found")
}

func (s BrowserSlice) FindIndex(f func(Browser) bool) int {
	for i, v := range s {
		if f(v) {
			return i
		}
	}

	return -1
}

func (s BrowserSlice) FindByName(name string) (Browser, error) {
	return s.Find(func(b Browser) bool { return strings.EqualFold(b.Name, name) })
}
