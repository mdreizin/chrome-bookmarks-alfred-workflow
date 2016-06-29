package model

import (
	"github.com/renstrom/fuzzysearch/fuzzy"
	"sort"
)

type BookmarkSlice []Bookmark

func (s BookmarkSlice) Add(v Bookmark) BookmarkSlice {
	return append(s, v)
}

func (s BookmarkSlice) Match(query string) BookmarkSlice {
	f := s[:0]

	for _, v := range s {
		matches := fuzzy.Find(query, []string{v.Name, v.URL})

		if len(matches) > 0 {
			f = f.Add(v)
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
