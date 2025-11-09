package cut_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/cut"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleCut_fromFile_fields() {
	// cat testdata/fields.txt | cut -d: -f2
	gloo.MustRun(
		Cut(Delimiter(":"), Fields("2"), gloo.File("testdata/fields.txt")),
	)
	// Output:
	// two
	// beta
	// second
}
