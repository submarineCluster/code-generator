package gotools

import (
	"fmt"
	"os/exec"

	"git.code.oa.com/leoshli/code-generator/conf"
)

//GoFmt ...
func GoFmt(dir string) error {

	fmtCmd := exec.Command("go", "fmt", dir+"/...")
	if conf.Verbose {
		fmt.Println(fmtCmd.String())
	}

	body, err := fmtCmd.CombinedOutput()
	if err != nil {
		return err
	}
	if conf.Verbose {
		fmt.Printf("%v\n", string(body))
	}
	return nil
}
