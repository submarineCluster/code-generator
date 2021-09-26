package gotools

import (
	"fmt"
	"os/exec"

	"git.code.oa.com/tencent_abtest/code-generator/conf"
)

//GoImport ...
func GoImport(dir string) error {

	// go imports
	importCmd := exec.Command("goimports", "-w", dir)

	if conf.Verbose {
		fmt.Println(importCmd.String())
	}

	// do command
	body, err := importCmd.CombinedOutput()
	if err != nil {
		return err
	}
	if conf.Verbose {
		fmt.Printf("%v\n", string(body))
	}
	return nil
}
