package api

type RunRequest struct {
	AssistantID string `json:"assistant_id"`
}

type RunResponse struct {
	ID           string      `json:"id"`
	Object       string      `json:"object"`
	CreatedAt    int         `json:"created_at"`
	AssistantID  string      `json:"assistant_id"`
	ThreadID     string      `json:"thread_id"`
	Status       string      `json:"status"`
	StartedAt    interface{} `json:"started_at"`
	ExpiresAt    int         `json:"expires_at"`
	CancelledAt  interface{} `json:"cancelled_at"`
	FailedAt     interface{} `json:"failed_at"`
	CompletedAt  interface{} `json:"completed_at"`
	LastError    interface{} `json:"last_error"`
	Model        string      `json:"model"`
	Instructions string      `json:"instructions"`
	Tools        []struct {
		Type string `json:"type"`
	} `json:"tools"`
	FileIds  []interface{} `json:"file_ids"`
	Metadata struct{}      `json:"metadata"`
	Usage    interface{}   `json:"usage"`
}

type RunStatusResponse struct {
	ID           string      `json:"id"`
	Object       string      `json:"object"`
	CreatedAt    int         `json:"created_at"`
	AssistantID  string      `json:"assistant_id"`
	ThreadID     string      `json:"thread_id"`
	Status       string      `json:"status"`
	StartedAt    int         `json:"started_at"`
	ExpiresAt    interface{} `json:"expires_at"`
	CancelledAt  interface{} `json:"cancelled_at"`
	FailedAt     interface{} `json:"failed_at"`
	CompletedAt  int         `json:"completed_at"`
	LastError    *string     `json:"last_error"`
	Model        string      `json:"model"`
	Instructions interface{} `json:"instructions"`
	Tools        []struct {
		Type string `json:"type"`
	} `json:"tools"`
	FileIds  []string `json:"file_ids"`
	Metadata struct{} `json:"metadata"`
	Usage    struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
