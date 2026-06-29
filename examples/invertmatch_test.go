package grep_test

import (
	"fmt"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-grep"
)

func ExampleGrep_invertMatch() {
	// echo -e "apple\nbanana\napricot" | grep -v "ap"
	output, _ := testable.Test(command.Grep("ap", command.GrepInvert), "apple\nbanana\napricot\n")
	fmt.Print(output)
	// Output:
	// banana
}
