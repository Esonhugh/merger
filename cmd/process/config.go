package process

import "merger/cmd"

var GlobalConfig = cmd.GlobalConfig

type InputFile struct {
	FileName      string
	Type          string
	Data          map[string]string
	Csv_Delimiter string
}

type GoCodeTemplate struct {
	StructDefine  string
	InputFileList []InputFile
	Output        OutputOption
}

type OutputOption struct {
	FileName string
	IsText   bool
	Field    string
}
