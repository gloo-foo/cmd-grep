package grep_test

import (
	"fmt"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-grep"
)

func ExampleGrep_ignoreCase() {
	// echo -e "Apple\nBanana\napricot" | grep -i "APPLE"
	output, _ := testable.Test(command.Grep("APPLE", command.GrepIgnoreCase), "Apple\nBanana\napricot\n")
	fmt.Print(output)
	// Output:
	// Apple
}
