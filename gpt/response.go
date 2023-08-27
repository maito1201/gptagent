package gpt

type ResponseData struct {
	Choices []struct {
		Message struct {
			Content       *string `json:"content"`
			*FunctionCall `json:"function_call"`
		} `json:"message"`
	} `json:"choices"`
}

type FunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}
