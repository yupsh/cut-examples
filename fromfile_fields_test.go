package cut_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/cut"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleCut_fromFile_fields() {
	// cat testdata/fields.txt | cut -d: -f2
	yup.MustRun(
		Cut(Delimiter(":"), Fields("2"), yup.File("testdata/fields.txt")),
	)
	// Output:
	// two
	// beta
	// second
}

