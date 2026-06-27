package grep_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-grep"

	"github.com/gloo-foo/testable"
)

func ExampleGrep_fromFile_caseInsensitive() {
	// grep -i "error" testdata/mixed_case.txt
	output, _ := testable.Test(command.Grep("error", command.GrepIgnoreCase), "Error: failed to connect\nWARNING: check logs\nerror: retry attempt\nInfo: process started\nERROR: timeout occurred\n")
	fmt.Print(output)
	// Output:
	// Error: failed to connect
	// error: retry attempt
	// ERROR: timeout occurred
}
