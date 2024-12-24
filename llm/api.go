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

func buildRequestBody() io.Reader {
	request := ChatRequest{
		Model: "gpt-4o-mini",
		Messages: []Message{
			{Role: "user", Content: "The absolite difference between highest and lowest point on Earth, in km"},
		},
	}
	jsonData, err := json.Marshal(request)
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

func buildRequest() *http.Request {
	body := buildRequestBody()

	apiKey := os.Getenv("OPENAI_API_KEY")
	req, err := http.NewRequest("POST",
		OPENAIURL,
		body,
	)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + apiKey)

	return req
}

func SendChatMessage() {
	req := buildRequest()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	dest := &ResponseObject{}
	parseChatMessage(bytes, dest)
	fmt.Println(dest.Choices[0].Message.Content)
}
