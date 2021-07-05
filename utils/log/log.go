package log

import (
	"fmt"

	"git.code.oa.com/tencent_abtest/code-generator/conf"
)

//Printf ...
func Printf(format string, args ...interface{}) {
	if conf.Verbose {
		fmt.Printf(format+"\n", args)
	}
}
