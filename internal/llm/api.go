package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const OPENAIURL = "https://api.openai.com/v1/chat/completions"

func GetResponseFormat() ResponseFormat {
	return ResponseFormat{
		Type: "json_schema",
		JSONSchema: JSONSchema{
			Name: "organize_files",
			Schema: SchemaStructure{
				Type: "object",
				Properties: Properties{
					Folders: FolderStructure{
						Type:        "object",
						Description: "Mapping of folder names to file lists.",
						PatternProperties: map[string]ArrayDefinition{
							"^[a-zA-Z0-9_\\- ]+$": {
								Type:        "array",
								Description: "List of files in the folder.",
								Items: ItemDef{
									Type: "string",
								},
								MinItems:    1,
								UniqueItems: true,
							},
						},
						AdditionalProperties: false,
					},
				},
			},
		},
	}
}

func createChatPayload(fileNames []string) ChatRequest {
	userPrompt := fmt.Sprintf("List of directories: %s", fileNames)
	responseFormat := GetResponseFormat()
	request := ChatRequest{
		Model: "gpt-4o",
		Messages: []Message{
			{Role: "system", Content: SYSTEM_PROMPT},
			{Role: "user", Content: userPrompt},
		},
		Response_format: responseFormat,
	}

	return request
}

func buildRequestBody(requestPayload ChatRequest) io.Reader {
	jsonData, err := json.Marshal(requestPayload)
	if err != nil {
		panic(err)
	}

	return bytes.NewReader(jsonData)
}

func parseChatMessage(resp []byte, dest *ResponseObject) {
	err := json.Unmarshal(resp, &dest)
	if err != nil {
		panic(err)
	}
}

func buildRequest(request ChatRequest) *http.Request {
	body := buildRequestBody(request)
	apiKey := os.Getenv("OPENAI_API_KEY")
	req, err := http.NewRequest("POST",
		OPENAIURL,
		body,
	)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	return req
}

func SendChatAPIMessage(fileNames []string) string {
	request := createChatPayload(fileNames)
	req := buildRequest(request)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	dest := &ResponseObject{}
	parseChatMessage(bytes, dest)

	return dest.Choices[0].Message.Content
}
