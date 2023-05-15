package process

import (
	"github.com/spf13/cobra"
	"log"
	"merger/cmd"
	template2 "merger/cmd/template"
	"os"
	"os/exec"
	"text/template"
)

var output string
var outputFormatIsText bool
var outputField string
var StructDefine string
var dryRun bool

func init() {
	MergeCommand.Flags().StringVarP(&output, "output", "o", "output.data", "output filename")
	MergeCommand.Flags().BoolVarP(&outputFormatIsText, "formatIsText", "t", false,
		"if true,output is oneline one string, else output as  (default) json")
	MergeCommand.Flags().StringVarP(&outputField, "field", "f", "", "the specific name of field you need export."+
		"if == \"\" will make all exported ")
	MergeCommand.Flags().StringVarP(&StructDefine, "struct", "s", os.Getenv("STRUCT_DEFINE"),
		"struct define like: \n"+
			"(type output) struct { Name   string `select:\"domain\"`;Source string `select:\"source\"`}\n")
	MergeCommand.Flags().BoolVarP(&dryRun, "dry", "d", false, "dry run. Only generate out.go file")
	cmd.RootCommand.AddCommand(MergeCommand)
}

var InputFiles []InputFile

var MergeCommand = &cobra.Command{
	Use:   "merge",
	Short: "merge is process your pre-add files",
	PreRun: func(cmd1 *cobra.Command, args []string) {
		if cmd.FromConfig != "" {
			output = GlobalConfig.GetString("output.filename")
			output = GlobalConfig.GetString("output.filename")
			outputFormatIsText = GlobalConfig.GetBool("output.formatIsText")
			StructDefine = GlobalConfig.GetString("output.structDefine")
			outputField = GlobalConfig.GetString("output.outputField")
		}
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
		if err != nil {
			panic(err)
		}
		err = tx.Execute(file, s)
		if err != nil {
			panic(err)
		}
		// force close file to save
		file.Close()

		if !dryRun {
			log.Println("Run out.go")
			exec.Command("go", "run", "out.go").Run()
		} else {
			log.Println("DryRun mod is setting.")
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
