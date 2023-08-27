package gpt

type RequestData struct {
	Model     string `json:"model"`
	Messages  `json:"messages"`
	Functions `json:"functions"`
}

type Messages []struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Functions []struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  `json:"parameters"`
}

type Parameters struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required"`
}

type Property struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
