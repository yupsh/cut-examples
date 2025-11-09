package cut_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/cut"
)

func ExampleCut_characters() {
	// echo "Hello World" | cut -c1-5
	gloo.MustRun(
		Cut(Characters("1-5"), strings.NewReader("Hello World")),
	)
	// Output:
	// Hello
}
