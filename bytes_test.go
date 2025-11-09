package cut_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/cut"
)

func ExampleCut_bytes() {
	// echo "abcdefgh" | cut -b1,3,5
	gloo.MustRun(
		Cut(Bytes("1,3,5"), strings.NewReader("abcdefgh")),
	)
	// Output:
	// ace
}
