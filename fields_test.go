package cut_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/cut"
)

func ExampleCut_fields() {
	// echo "one:two:three:four" | cut -d: -f2
	gloo.MustRun(
		Cut(Delimiter(":"), Fields("2"), strings.NewReader("one:two:three:four")),
	)
	// Output:
	// two
}
