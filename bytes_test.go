package cut_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/cut"
)

func ExampleCut_bytes() {
	// echo "abcdefgh" | cut -b1,3,5
	yup.MustRun(
		Cut(Bytes("1,3,5"), strings.NewReader("abcdefgh")),
	)
	// Output:
	// ace
}

