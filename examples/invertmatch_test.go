package grep_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-grep"

	"github.com/gloo-foo/testable"
)

func ExampleGrep_invertMatch() {
	// echo -e "apple\nbanana\napricot" | grep -v "ap"
	output, _ := testable.Test(command.Grep("ap", command.GrepInvert), "apple\nbanana\napricot\n")
	fmt.Print(output)
	// Output:
	// banana
}
