package cut_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/cut"
)

func ExampleCut_characters() {
	// echo "Hello World" | cut -c1-5
	yup.MustRun(
		Cut(Characters("1-5"), strings.NewReader("Hello World")),
	)
	// Output:
	// Hello
}

