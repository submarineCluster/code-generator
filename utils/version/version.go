package version

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"git.code.oa.com/leoshli/code-generator/ldflags"

	"git.code.oa.com/leoshli/code-generator/utils/callstack"
)

//GetVersion ...
func GetVersion() string {
	frame, err := callstack.CallFrame(0)
	if err != nil {
		return ""
	}
	goPkgPath := filepath.Join(filepath.Dir(frame.File()), "../..")

	version := strings.TrimPrefix(goPkgPath, os.Getenv("GOPATH")+"/pkg/mod/")
	return fmt.Sprintf("%s\ngoversion=%v\nbuild-time=%v\n", version, ldflags.GOVersion, ldflags.BuildTime)
}
