package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var GlobalConfig *viper.Viper
var FromConfig string

func init() {
	RootCommand.PersistentFlags().StringVarP(&FromConfig, "from-config", "c", "", "only merge and get information from config file")
	GlobalConfig = viper.New()
}

var RootCommand = &cobra.Command{
	Use:   "doc-merger",
	Short: "doc-merger is a tool for merge different type of document",
	Long:  "doc-merger is a tool for merge different type of document",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if FromConfig == "" {
			GlobalConfig.SetConfigFile(".sculptor.json") // optionally look for config in the working directory
			GlobalConfig.SetConfigType("json")
		} else {
			GlobalConfig.SetConfigFile(FromConfig)
			GlobalConfig.SetConfigType("json")
		}
		if err := GlobalConfig.ReadInConfig(); err != nil {
			// 如果没有找到配置文件，尝试创建
			fmt.Println("未找到 config 文件，已创建默认配置")
		}
		GlobalConfig.SetDefault("used", true)
		GlobalConfig.WriteConfig() // 将默认设置写入新的配置文件
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
