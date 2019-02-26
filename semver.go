package semver // import "go.nanasi880.dev/semver"

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Version
type Version struct {
	Raw   string
	Major int
	Minor int
	Patch int
}

// Less is true return if `v` less `rhs`.
func (v Version) Less(rhs Version) bool {

	left := [3]int{
		v.Major, v.Minor, v.Patch,
	}
	right := [3]int{
		rhs.Major, rhs.Minor, rhs.Patch,
	}

	return v.cmp(left, right) < 0
}

// cmp is negative return if `left` less `right`.
// if equal, return 0.
// if greater, return positive.
func (v *Version) cmp(left [3]int, right [3]int) int {

	for i := 0; i < 3; i++ {

		if left[i] == right[i] {
			continue
		}

		if left[i] < right[i] {
			return -1
		}

		return 1
	}

	return 0
}

// Versions is slice of Version.
type Versions []Version

// Len implements sort.Interface.
func (s Versions) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s Versions) Less(i, j int) bool {
	return s[i].Less(s[j])
}

// Swap implements sort.Interface.
func (s Versions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var (
	semverPattern = regexp.MustCompile(`^[0-9]+\.?[0-9]*\.?[0-9]*$`)
)

// Parse is parse of semver string with "v" prefix
func Parse(s string) (Version, error) {
	return ParseWithPrefix(s, "v")
}

// ParseWithPrefix is parse of semver string with custom prefix
func ParseWithPrefix(s string, prefix string) (Version, error) {

	var version Version
	version.Raw = s

	s = strings.TrimLeft(s, prefix)

	if !semverPattern.MatchString(s) {
		return Version{}, fmt.Errorf("parse error: invalid pattern")
	}

	var (
		match  = strings.Split(s, ".")
		values [3]int
	)
	for i := range match {
		v, err := strconv.ParseInt(match[i], 10, 32)
		if err != nil {
			return Version{}, fmt.Errorf("parse error: parse int failed: %v", err)
		}

		values[i] = int(v)
	}

	version.Major = values[0]
	version.Minor = values[1]
	version.Patch = values[2]

	return version, nil
}
