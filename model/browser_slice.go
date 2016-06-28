package model

import (
	"strings"
	"errors"
	"github.com/renstrom/fuzzysearch/fuzzy"
)

type BrowserSlice []Browser

func (s BrowserSlice) Add(v Browser) BrowserSlice {
	return append(s, v)
}

func (s BrowserSlice) Match(query string) BrowserSlice {
	f := s[:0]

	for _, v := range s {
		matches := fuzzy.Find(query, []string{v.Name})

		if len(matches) > 0 {
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
