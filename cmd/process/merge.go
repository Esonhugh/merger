package process

import (
	"github.com/spf13/cobra"
	"merger/cmd"
	template2 "merger/cmd/template"
	"os"
	"text/template"
)

var output string
var outputFormatIsText bool
var outputField string
var StructDefine string

func init() {
	MergeCommand.Flags().StringVarP(&output, "output", "o", "output.data", "output filename")
	MergeCommand.Flags().BoolVarP(&outputFormatIsText, "formatIsText", "t", false,
		"if true,output is oneline one string, else output as  (default) json")
	MergeCommand.Flags().StringVarP(&outputField, "field", "f", "", "the specific name of field you need export."+
		"if == \"\" will make all exported ")
	MergeCommand.Flags().StringVarP(&StructDefine, "struct", "s", os.Getenv("STRUCT_DEFINE"),
		"struct define like: \n"+
			"(type output) struct { Name   string `select:\"domain\"`;Source string `select:\"source\"`}\n")
	cmd.RootCommand.AddCommand(MergeCommand)
}

var InputFiles []InputFile

var MergeCommand = &cobra.Command{
	Use:   "merge",
	Short: "merge is process your pre-add files",
	PreRun: func(cmd *cobra.Command, args []string) {
		FileData := GlobalConfig.Sub("file")
		if FileData == nil {
			panic("Not defined file data")
		}
		for i, _ := range FileData.AllSettings() {
			var File InputFile
			FileData.Sub(i).Unmarshal(&File)
			File.FileName = i
			InputFiles = append(InputFiles, File)
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		tx := template.Must(template.New("main").Parse(template2.MainTemplate))
		s := GoCodeTemplate{
			StructDefine:  StructDefine,
			InputFileList: InputFiles,
			Output: OutputOption{
				FileName: output,
				Field:    outputField,
				IsText:   outputFormatIsText,
			},
		}
		file, err := os.OpenFile("out.go", os.O_CREATE|os.O_WRONLY, 0644)
		defer file.Close()
		if err != nil {
			panic(err)
		}
		err = tx.Execute(file, s)
		if err != nil {
			panic(err)
		}
	},

	PostRun: func(cmd *cobra.Command, args []string) {
		defer GlobalConfig.WriteConfig()
		GlobalConfig.Set("output.filename", output)
		GlobalConfig.Set("output.formatIsText", outputFormatIsText)
		GlobalConfig.Set("output.structDefine", StructDefine)
		GlobalConfig.Set("output.outputField", outputField)
	},
}
