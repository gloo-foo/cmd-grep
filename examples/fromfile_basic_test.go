package grep_test

import (
	"fmt"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-grep"
)

func ExampleGrep_fromFile_basic() {
	// grep "ap" testdata/fruits.txt
	output, _ := testable.Test(command.Grep("ap"), "apple\nbanana\napricot\ncherry\norange\n")
	fmt.Print(output)
	// Output:
	// apple
	// apricot
}
