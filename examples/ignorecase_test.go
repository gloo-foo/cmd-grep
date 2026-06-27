package grep_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-grep"

	"github.com/gloo-foo/testable"
)

func ExampleGrep_ignoreCase() {
	// echo -e "Apple\nBanana\napricot" | grep -i "APPLE"
	output, _ := testable.Test(command.Grep("APPLE", command.GrepIgnoreCase), "Apple\nBanana\napricot\n")
	fmt.Print(output)
	// Output:
	// Apple
}
