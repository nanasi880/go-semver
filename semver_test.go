package semver

import (
	"sort"
	"testing"
)

func TestParse(t *testing.T) {

	testCases := Versions {
		{
			Raw:   "v1.0.0",
			Major: 1,
			Minor: 0,
			Patch: 0,
		},
		{
			Raw:   "v1.2.3",
			Major: 1,
			Minor: 2,
			Patch: 3,
		},
		{
			Raw:   "v1.1",
			Major: 1,
			Minor: 1,
			Patch: 0,
		},
		{
			Raw:   "v1",
			Major: 1,
			Minor: 0,
			Patch: 0,
		},
	}

	for _, test := range testCases {

		v, err := Parse(test.Raw)
		if err != nil {
			t.Fatal(test.Raw, " ", err)
		}

		if test.Major != v.Major || test.Minor != v.Minor || test.Patch != v.Patch {
			t.Fatal(test.Raw, " ", test)
		}
	}
}

func TestParseError(t *testing.T) {

	invalidStrings := []string {
		"a.b.c",
		"abc",
		"v1.0.0-devel",
		"v1.0.0.0",
	}

	for _, s := range invalidStrings {

		v, err := Parse(s)
		if err == nil {
			t.Fatal(s, " ", v)
		}
	}
}

func TestSort(t *testing.T) {

	versionStrings := []string {
		"v1.0.0",
		"v2",
		"v1.2.3",
		"v1.1",
	}

	versions := make(Versions, len(versionStrings))
	for i, s := range versionStrings {
		v, err := Parse(s)
		if err != nil {
			t.Fatal(err)
		}

		versions[i] = v
	}

	sort.Sort(versions)

	wants := []string {
		"v1.0.0",
		"v1.1",
		"v1.2.3",
		"v2",
	}

	for i := range versions {
		if versions[i].Raw != wants[i] {
			t.Fatal(versions)
		}
	}
}
