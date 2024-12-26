package display

import (
	"encoding/json"
	"fmt"
)

func DisplayFolders(content string) {
	var result struct {
		Folders map[string][]string `json:"folders"`
	}
	// fmt.Println(dest.Choices[0].Message.Content)
	err := json.Unmarshal([]byte(content), &result)
	if err != nil {
		panic(err)
	}

	for index, value := range result.Folders {
		fmt.Printf("The Folder Name: %s -  The files: %s\n", index, value)
	}
}
