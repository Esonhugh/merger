package template

var MainTemplate = `package main

import (
	{{if .Output.IsText -}}
	"fmt"
	{{else}}
	"encoding/json"
	{{- end}}
	"github.com/esonhugh/sculptor"
	"os"
	"sync"
)

type FinalStruct {{.StructDefine}}

// Demo of merge subfinder and oneforall
func main() {
	wg := &sync.WaitGroup{}

	{{range $i, $Elem := .InputFileList }}
	doc_{{$i}} := sculptor.NewDataSculptorWithWg("{{$Elem.FileName}}", wg).
		SetDocType(sculptor.DocumentType("{{$Elem.Type}}")).
		{{range $Key, $Value := $Elem.Data -}}
		SetQuery("{{$Key}}", "{{$Value}}").
		{{- end}}
		SetTargetStruct(&FinalStruct{}){{if eq $Elem.Type "csv" -}}.SetCSVDelimiter('{{$Elem.Csv_Delimiter}}')
		{{end}}
	{{- end}}

	common_output := sculptor.Merge({{range $i, $Elem := .InputFileList -}} 
			{{if ne $i 0}} ,{{end}}doc_{{$i}}
	{{- end}})
	{{range $i, $Elem := .InputFileList }}
	doc_{{$i}}.Do()
	{{- end}}
	
	go func() {
		wg.Wait()
		close(common_output)
	}()

	OutPutFile, err := os.Create("{{.Output.FileName}}")
	if err != nil {
		panic(err)
	}
	defer OutPutFile.Close()
	for i := range common_output {
		{{ if .Output.IsText}}
			OutPutFile.WriteString(fmt.Sprintf("%v\n", i.(*FinalStruct){{if ne .Output.Field ""}}.{{.Output.Field}}{{end}}))
		{{else}}
    		jsonString, err := json.Marshal(i.(*FinalStruct))
    		if err != nil {
				panic("json marshal error:"+ err.Error())
			}
			OutPutFile.WriteString(string(jsonString) + "\n")
		{{end}}
	}
}
`
