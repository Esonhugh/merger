package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var DebugLevel string
var GlobalConfig *viper.Viper

func init() {
	GlobalConfig = viper.New()
	GlobalConfig.SetConfigFile("./.sculptor.yaml")
	GlobalConfig.SetConfigType("yaml")
	if err := GlobalConfig.ReadInConfig(); err != nil {
		// 如果没有找到配置文件，尝试创建
		fmt.Println("未找到 config 文件，已创建默认配置")
	}
	GlobalConfig.SetDefault("used", true)
	GlobalConfig.WriteConfig() // 将默认设置写入新的配置文件
	RootCommand.PersistentFlags().StringVar(&DebugLevel, "debug", "info", "debug level")
}

var RootCommand = &cobra.Command{
	Use:   "doc-merger",
	Short: "doc-merger is a tool for merge different type of document",
	Long:  "doc-merger is a tool for merge different type of document",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
