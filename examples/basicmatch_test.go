package grep_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-grep"

	"github.com/gloo-foo/testable"
)

func ExampleGrep_basicMatch() {
	// echo -e "apple\nbanana\napricot\ncherry" | grep "ap"
	output, _ := testable.Test(command.Grep("ap"), "apple\nbanana\napricot\ncherry\n")
	fmt.Print(output)
	// Output:
	// apple
	// apricot
}
