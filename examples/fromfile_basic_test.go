package grep_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-grep"

	"github.com/gloo-foo/testable"
)

func ExampleGrep_fromFile_basic() {
	// grep "ap" testdata/fruits.txt
	output, _ := testable.Test(command.Grep("ap"), "apple\nbanana\napricot\ncherry\norange\n")
	fmt.Print(output)
	// Output:
	// apple
	// apricot
}
