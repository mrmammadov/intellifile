package main

import (
	"intellifile/cmd/cli"
	"intellifile/internal/display"
	"intellifile/internal/fileops"
	"intellifile/internal/llm"
)

func main() {
	path := cli.Execute()

	fileNames := fileops.ReadFolder(path)
	response := llm.SendChatAPIMessage(fileNames)
	display.DisplayFolders(response)
}
