package resource

import (
	"fmt"
	"os"

	"github.com/submarineCluster/code-generator/conf"
	"github.com/submarineCluster/code-generator/utils/gotools"

	"github.com/markbates/pkger"

	"github.com/spf13/cobra"

	"github.com/submarineCluster/code-generator/cmd/model"
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
	if conf.Verbose {
		fmt.Printf("%+v\n", metadata)
	}

	// gen code file by metadata
	err = GenFileByMetadata(metadata)
	if err != nil {
		fmt.Printf("genCodeFileByResourceData fail: %v", err)
		return
	}

	// do fmt
	if !conf.ProtoOnly {
		err = gotools.GoFmt(metadata.GenDir)
		if err != nil {
			fmt.Printf("gofmt fail:%v", err)
		}
		err = gotools.GoImport(metadata.GenDir)
		if err != nil {
			fmt.Printf("gofmt fail:%v", err)
		}
	}

	// verify
	err = verify(metadata)
	if err != nil {
		fmt.Printf("verify fail: %v", err)
	}

	// done
}

// verify 。。。
func verify(metadata *model.Metadata) error {
	return nil
}

//verifyIsExist ...
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
