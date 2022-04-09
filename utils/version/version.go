package version

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/submarineCluster/code-generator/ldflags"

	"github.com/submarineCluster/code-generator/utils/callstack"
)

//GetVersion ...
func GetVersion() string {
	// 获取frame
	frame, err := callstack.CallFrame(0)
	if err != nil {
		return ""
	}

	// 项目路径
	goPkgPath := filepath.Join(filepath.Dir(frame.File()), "../..")

	// 获取版本
	version := strings.TrimPrefix(goPkgPath, os.Getenv("GOPATH")+"/pkg/mod/")
	return fmt.Sprintf("%s\ngoversion=%v\nbuild-time=%v\n", version, ldflags.GOVersion, ldflags.BuildTime)
}
