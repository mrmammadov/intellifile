package llm

type ChatRequest struct {
	Model           string         `json:"model"`
	Messages        []Message      `json:"messages"`
	Response_format ResponseFormat `json:"response_format"`
}

type ResponseFormat struct {
	Type       string     `json:"type"`
	JSONSchema JSONSchema `json:"json_schema"`
}

type JSONSchema struct {
	Name   string          `json:"name"`
	Schema SchemaStructure `json:"schema"`
}

type SchemaStructure struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}

type Properties struct {
	Folders FolderStructure `json:"folders"`
}

type FolderStructure struct {
	Type                 string                     `json:"type"`
	Description          string                     `json:"description"`
	PatternProperties    map[string]ArrayDefinition `json:"patternProperties"`
	AdditionalProperties bool                       `json:"additionalProperties"`
}

type ArrayDefinition struct {
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Items       ItemDef `json:"items"`
	MinItems    int     `json:"minItems"`
	UniqueItems bool    `json:"uniqueItems"`
}

type ItemDef struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choices struct {
	Index   int
	Message Message
}

type ResponseObject struct {
	Model   string
	Choices []Choices
}
