package model

import (
	"errors"
	"github.com/renstrom/fuzzysearch/fuzzy"
	"sort"
)

type ProfileSlice []Profile

func (s ProfileSlice) Add(v Profile) ProfileSlice {
	return append(s, v)
}

func (s ProfileSlice) Match(query string) ProfileSlice {
	f := s[:0]

	for _, v := range s {
		targets := []string{}

		for _, x := range []string{v.Name, v.DisplayName, v.UserName, v.UserEmail} {
			if x != "" {
				targets = append(targets, x)
			}
		}

		matches := fuzzy.Find(query, targets)

		if len(matches) > 0 {
			f = f.Add(v)
		}
	}

	return f
}

func (s ProfileSlice) Find(f func(Profile) bool) (Profile, error) {
	for _, v := range s {
		if f(v) {
			return v, nil
		}
	}

	return Profile{}, errors.New("`Profile` is not found")
}

func (s ProfileSlice) FirstActive() (Profile, error) {
	return s.Find(func(p Profile) bool { return p.IsActive == true })
}

func (s ProfileSlice) Len() int {
	return len(s)
}

func (s ProfileSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ProfileSlice) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

func (s ProfileSlice) Sort() ProfileSlice {
	sort.Sort(s)

	return s
}
