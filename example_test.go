package semver_test

import (
	"fmt"

	"go.nanasi880.dev/semver"
)

func ExampleParse() {

	version, err := semver.Parse("v1.2.3")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d %d %d", version.Major, version.Minor, version.Patch)
	// Output:
	// 1 2 3
}
