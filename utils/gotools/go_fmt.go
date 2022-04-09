package gotools

import (
	"fmt"
	"os/exec"

	"github.com/submarineCluster/code-generator/conf"
)

//GoFmt ...
func GoFmt(dir string) error {

	// go fmt
	fmtCmd := exec.Command("go", "fmt", dir+"/...")
	if conf.Verbose {
		fmt.Println(fmtCmd.String())
	}

	// do command
	body, err := fmtCmd.CombinedOutput()
	if err != nil {
		return err
	}
	if conf.Verbose {
		fmt.Printf("%v\n", string(body))
	}
	return nil
}
