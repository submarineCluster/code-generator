package log

import (
	"fmt"

	"github.com/submarineCluster/code-generator/conf"
)

//Printf ...
func Printf(format string, args ...interface{}) {
	if conf.Verbose {
		fmt.Printf(format+"\n", args)
	}
}
