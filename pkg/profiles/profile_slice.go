package profiles

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
)

type ProfileSlice []*Profile

func (s ProfileSlice) Add(v *Profile) ProfileSlice {
	return append(s, v)
}

func (s ProfileSlice) Match(query string) ProfileSlice {
	f := ProfileSlice{}

	re := regexp.MustCompile(fmt.Sprintf("(?i)%s", regexp.QuoteMeta(query)))

	for _, v := range s {
		if re.MatchString(v.Name) || re.MatchString(v.DisplayName) || re.MatchString(v.UserName) || re.MatchString(v.UserEmail) {
			f = f.Add(v)
		}
	}

	return f
}

func (s ProfileSlice) Find(f func(*Profile) bool) (*Profile, error) {
	for _, v := range s {
		if f(v) {
			return v, nil
		}
	}

	return nil, errors.New("profile is not found")
}

func (s ProfileSlice) FirstActive() (*Profile, error) {
	return s.Find(func(p *Profile) bool { return p.IsActive == true })
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
