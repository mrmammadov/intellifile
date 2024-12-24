package llm

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseMessage struct {
	Role string
	Content string
}

type Choices struct {
	index int
	Message ResponseMessage
}

type ResponseObject struct {
	Model string
	Choices []Choices
}
