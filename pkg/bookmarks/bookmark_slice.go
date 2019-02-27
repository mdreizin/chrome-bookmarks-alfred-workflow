package bookmarks

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type BookmarkSlice []Bookmark

func (s BookmarkSlice) Add(v Bookmark) BookmarkSlice {
	return append(s, v)
}

func (s BookmarkSlice) Match(query string) BookmarkSlice {
	bookmarkSlice := BookmarkSlice{}
	fields := strings.Fields(query)

	var regexps []*regexp.Regexp

	for _, field := range fields {
		re := regexp.MustCompile(fmt.Sprintf("(?i)%s", regexp.QuoteMeta(field)))

		regexps = append(regexps, re)
	}

	for _, v := range s {
		every := false

		for _, re := range regexps {
			m := re.MatchString(v.Name) || re.MatchString(v.URL)

			if !m {
				every = false

				break
			} else {
				every = true
			}
		}

		if every == true {
			bookmarkSlice = bookmarkSlice.Add(v)
		}
	}

	return bookmarkSlice
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
