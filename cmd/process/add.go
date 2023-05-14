package process

import (
	"github.com/spf13/cobra"
	"merger/cmd"
	"os"
)

var document_type string
var csv_delimiter string
var filename string

func init() {
	AddCommand.PersistentFlags().StringVarP(&document_type, "type", "t", "csv", "document type")
	AddCommand.Flags().StringVarP(&csv_delimiter, "delimiter", "d", ",", "csv delimiter")
	AddCommand.Flags().StringVarP(&filename, "filename", "f", "input.csv", "filename")
	cmd.RootCommand.AddCommand(AddCommand)
}

var AddCommand = &cobra.Command{
	Use:   "add",
	Short: "add is a tool for add one document with its processing",
	Long: "add is a tool for add one document with its processing. " +
		"The args are ([tagValue for select] [query for this field]){.*}",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args)%2 == 1 {
			panic("args length must be even")
		}
		if len(csv_delimiter) != 1 {
			panic("csv delimiter must be one character")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		FileExist(filename)
		queryMap := make(map[string]string, 10)
		// map Key (string)tagName  => (string)query
		for i, v := range args {
			if i%2 == 1 {
				queryMap[args[i-1]] = v
			}
		}
		GlobalConfig.Set("file."+filename+".data", queryMap)
		GlobalConfig.Set("file."+filename+".type", document_type)
		GlobalConfig.Set("file."+filename+".csv_delimiter", csv_delimiter)
		defer GlobalConfig.WriteConfig()
	},
}

func FileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
