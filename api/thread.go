package api

type CreateThreadResponse struct {
	ID        string            `json:"id"`
	Object    string            `json:"object"`
	CreatedAt int               `json:"created_at"`
	Metadata  map[string]string `json:"metadata"`
}
