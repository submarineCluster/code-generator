package log

import (
	"fmt"

	"git.code.oa.com/leoshli/code-generator/conf"
)

//Printf ...
func Printf(format string, args ...interface{}) {
	if conf.Verbose {
		fmt.Printf(format+"\n", args)
	}
}
