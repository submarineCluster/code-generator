package resource

import (
	"fmt"
	"os"

	"git.code.oa.com/leoshli/code-generator/utils/gotools"

	"github.com/markbates/pkger"

	"github.com/spf13/cobra"

	"git.code.oa.com/leoshli/code-generator/cmd/model"
)

//Run ...
func Run(cmd *cobra.Command, args []string) {

	//fmt.Printf("start run resource")
	//WalkPkgerFile()
	// gen resource-data
	metadata, err := model.GenMetadata()
	if err != nil {
		fmt.Printf("GenMetadata fail: %v", err)
		return
	}

	// gen code file by metadata
	err = GenFileByMetadata(metadata)
	if err != nil {
		fmt.Printf("genCodeFileByResourceData fail: %v", err)
		return
	}

	// do fmt
	err = gotools.GoFmt(metadata.GenDir)
	if err != nil {
		fmt.Printf("gofmt fail:%v", err)
	}
	err = gotools.GoImport(metadata.GenDir)
	if err != nil {
		fmt.Printf("gofmt fail:%v", err)
	}

	// verify
	err = verify(metadata)
	if err != nil {
		fmt.Printf("verify fail: %v", err)
	}

	// done
}

func verify(metadata *model.Metadata) error {
	return nil
}

func verifyIsExist(metadata *model.Metadata) bool {
	_, err := os.Stat(metadata.GenDir)
	return os.IsExist(err)
}

//WalkPkgerFile ...
func WalkPkgerFile() {

	fmt.Printf("start walkPkger\n")

	pkger.Walk("/cmd/resource", func(path string, info os.FileInfo, err error) error {
		//fmt.Printf(info.Name() + "\n")
		if info.IsDir() {
			return nil
		}

		fmt.Printf(info.Name() + "\n")
		return nil
	})
}
