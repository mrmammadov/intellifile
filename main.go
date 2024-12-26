package main

import (
	"intellifile/cmd/cli"
	"intellifile/internal/fileops"
	"intellifile/internal/llm"
)

func main() {
	path := cli.Execute()

	fileNames := fileops.ReadFolder(path)
	llm.SendChatAPIMessage(fileNames)
}
