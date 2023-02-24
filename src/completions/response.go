package completions

type Response struct {
	Id      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Usage   Usage     `json:"usage,omitempty"`
	Choices []Choices `json:"choices"`
}

type Choices struct {
	Text         string   `json:"text,omitempty"`
	Index        int      `json:"index,omitempty"`
	Logprobs     Logprobs `json:"logprobs,omitempty"`
	FinishReason string   `json:"finish_reason,omitempty"`
}

type Logprobs struct {
	Tokens        []string             `json:"tokens,omitempty"`
	TokenLogprobs []float32            `json:"token_logprobs,omitempty"`
	TopLogprobs   []map[string]float32 `json:"top_logprobs,omitempty"`
	TextOffset    []int                `json:"text_offset,omitempty"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}
