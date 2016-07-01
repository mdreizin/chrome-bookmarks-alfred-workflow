package model

import (
	"regexp"
	"sort"
	"strings"
)

type BookmarkSlice []Bookmark

func (s BookmarkSlice) Add(v Bookmark) BookmarkSlice {
	return append(s, v)
}

func (s BookmarkSlice) Match(query string) BookmarkSlice {
	f := s[:0]

	fields := strings.Fields(query)
	regexps := []*regexp.Regexp{}

	for _, field := range fields {
		re := regexp.MustCompile(regexp.QuoteMeta(field))

		regexps = append(regexps, re)
	}

	for _, v := range s {
		for _, re := range regexps {
			if re.MatchString(v.Name) || re.MatchString(v.URL) {
				f = f.Add(v)
				break
			}
		}
	}

	return f
}

func (s BookmarkSlice) Len() int {
	return len(s)
}

func (s BookmarkSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s BookmarkSlice) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

func (s BookmarkSlice) Sort() BookmarkSlice {
	sort.Sort(s)

	return s
}
