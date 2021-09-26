package clean

import (
	"fmt"

	"github.com/spf13/cobra"

	"git.code.oa.com/tencent_abtest/code-generator/cmd/model"
)

//Run ...
func Run(cmd *cobra.Command, args []string) {

	// genMetadata
	metadata, err := model.GenMetadata()
	if err != nil {
		fmt.Printf("GenMetadata fail\n")
		return
	}
	// clean
	err = clean(metadata)
	if err != nil {
		fmt.Printf("clean fail\n")
		return
	}
}
