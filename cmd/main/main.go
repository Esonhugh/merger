package main

import (
	"merger/cmd"
	_ "merger/cmd/process"
)

func main() {
	cmd.RootCommand.Execute()
}
