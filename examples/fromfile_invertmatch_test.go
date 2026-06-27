package grep_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-grep"

	"github.com/gloo-foo/testable"
)

func ExampleGrep_fromFile_invertMatch() {
	// grep -v "ERROR" testdata/log_entries.txt
	output, _ := testable.Test(command.Grep("ERROR", command.GrepInvert), "2024-01-15 10:23:45 INFO Server started\n2024-01-15 10:24:12 ERROR Connection failed\n2024-01-15 10:24:30 INFO Request processed\n2024-01-15 10:25:01 WARN Low memory\n2024-01-15 10:25:45 ERROR Database timeout\n2024-01-15 10:26:00 INFO Cache cleared\n")
	fmt.Print(output)
	// Output:
	// 2024-01-15 10:23:45 INFO Server started
	// 2024-01-15 10:24:30 INFO Request processed
	// 2024-01-15 10:25:01 WARN Low memory
	// 2024-01-15 10:26:00 INFO Cache cleared
}
