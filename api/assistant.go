package api

type AssistantCreateRequest struct {
	Instructions string              `json:"instructions"`
	Name         string              `json:"name"`
	Model        string              `json:"model"`
	Tools        []map[string]string `json:"tools"`
}

type AssistantCreateResponse struct {
	ID           string      `json:"id"`
	Object       string      `json:"object"`
	CreatedAt    int         `json:"created_at"`
	Name         string      `json:"name"`
	Description  interface{} `json:"description"`
	Model        string      `json:"model"`
	Instructions string      `json:"instructions"`
	Tools        []struct {
		Type string `json:"type"`
	} `json:"tools"`
	FileIds  []interface{} `json:"file_ids"`
	Metadata struct{}      `json:"metadata"`
	Error    struct {
		Message string `json:"message"`
	} `json:"error"`
}
