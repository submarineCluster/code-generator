package goenv

import (
	"os/exec"
	"path/filepath"
	"strings"

	"git.code.oa.com/tencent_abtest/code-generator/conf"

	"git.code.oa.com/tencent_abtest/code-generator/utils/callstack"

	"github.com/pkg/errors"
)

/*
get go mod name
*/

//GetModuleName ...
func GetModuleName() (string, error) {

	// get module name
	moduleCMD := exec.Command("go", "list", "-m")

	// do command
	body, err := moduleCMD.CombinedOutput()
	if err != nil {
		return "", err
	}

	result := strings.Trim(string(body), "\n")
	if strings.EqualFold(result, "command-line-arguments") {
		return "", errors.Errorf("Must be under the go mod project")
	}

	return result, nil
}

//GetModulePath ...
func GetModulePath() (string, error) {
	moduleCMD := exec.Command("go", "env", "GOMOD")
	body, err := moduleCMD.CombinedOutput()
	if err != nil {
		return "", err
	}

	result := strings.Trim(string(body), "\n")

	// not go mod project
	if strings.EqualFold(result, "/dev/null") || len(result) == 0 {
		if conf.ProtoOnly {
			moduleCMD = exec.Command("pwd")
			body, err = moduleCMD.CombinedOutput()
			if err != nil {
				return "", err
			}
			return strings.Trim(string(body), "\n"), nil
		}
		return "", errors.Errorf("Must be under the go mod project")
	}

	return filepath.Dir(string(body)), nil
}

//GetGoPkg ...
func GetGoPkg() string {
	frame, err := callstack.CallFrame(0)
	if err != nil {
		return ""
	}
	return filepath.Join(filepath.Dir(frame.File()), "../..")
}
