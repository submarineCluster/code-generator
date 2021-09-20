/*Package cmd ...
Copyright © 2020 submarineCluster <1145480206@qqcom>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/spf13/cobra"

	"git.code.oa.com/tencent_abtest/code-generator/cmd/resource"
	"git.code.oa.com/tencent_abtest/code-generator/conf"
)

// resourceCmd represents the resource command
var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "generate crud code for resource",
	Long: `generate crud code for resource:

code-generator resource will generator crud code for resource`,
	Run:     resource.Run,
	Example: "code-generator resource -n codeGenerator",
	PreRun: func(cmd *cobra.Command, args []string) {

		//fmt.Printf("%v\n%v\n%v\n", ldflags.Version, ldflags.GOVersion, ldflags.BuildTime)
		//
		//upgradeCmd := exec.Command("go", "get", "-u", "git.code.oa.com/tencent_abtest/code-generator")
		//log.Printf("%v", upgradeCmd.String())
		//body, err := upgradeCmd.CombinedOutput()
		//if err != nil {
		//	log.Printf("%v, err=%v", upgradeCmd.String(), err)
		//}
		//log.Printf("%v", string(body))
	},
}

func init() {
	rootCmd.AddCommand(resourceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resourceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	resourceCmd.Flags().StringVarP(&conf.ResourceName, "name", "n", "", "资源对象名称，必填，如demo")
	resourceCmd.Flags().StringVarP(&conf.TemplateDir, "templateDir", "t", "",
		"path of template-code，推荐使用ddd，备选trpc, 例如 code-generator resource -n demo -t ddd")
	resourceCmd.Flags().BoolVarP(&conf.Verbose, "verbose", "v", false, "verbose")
	resourceCmd.Flags().BoolVar(&conf.DaoMetrics, "daoMetrics", true, "指标上报代码生成，TODO")
	resourceCmd.Flags().BoolVar(&conf.APIServer, "apiServer", false, "API server代码生成，TODO 暂不需要")
	resourceCmd.Flags().BoolVarP(&conf.CacheEnable, "cacheEnable", "c", false, "开启对象数据redis缓存，一般不需要，除非像用户等经常需要Get获取的")
	resourceCmd.Flags().StringVarP(&conf.StorageT, "storageType", "s", string(conf.StorageTypeMongo),
		"storage-type, eg: mysql/mongo, mongo default， TODO")
	resourceCmd.Flags().BoolVarP(&conf.ProtoOnly, "protoOnly", "p", false, "单纯生成proto文件")
	resourceCmd.Flags().StringVarP(&conf.AppName, "app", "a", "TAB", "appName default TAB, 生成协议文件时需要传入")
	resourceCmd.Flags().StringVar(&conf.ServerName, "server", "", "serverName, 一般用123上的服务名，请使用小蛇式如 mab_scheduler")
	resourceCmd.MarkFlagRequired("name")
	resourceCmd.MarkFlagRequired("server")
}
