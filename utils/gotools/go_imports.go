package gotools

import (
	"fmt"
	"os/exec"

	"git.code.oa.com/leoshli/code-generator/conf"
)

//GoImport ...
func GoImport(dir string) error {

	importCmd := exec.Command("goimports", "-w", dir)

	if conf.Verbose {
		fmt.Println(importCmd.String())
	}

	body, err := importCmd.CombinedOutput()
	if err != nil {
		return err
	}
	if conf.Verbose {
		fmt.Printf("%v\n", string(body))
	}
	return nil
}
