package api

import (
	"encoding/json"
)

type MessageRequest struct {
	Role     string            `json:"role"`
	Content  string            `json:"content"`
	FileIDs  []string          `json:"file_ids,omitempty"`
	Metadata map[string]string `json:"metadata"`
}

type MessageResponse struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	CreatedAt int    `json:"created_at"`
	ThreadID  string `json:"thread_id"`
	Role      string `json:"role"`
	Content   []struct {
		Type string `json:"type"`
		Text struct {
			Value       string       `json:"value"`
			Annotations []Annotation `json:"annotations"`
		} `json:"text"`
		ImageFile struct {
			FileID string `json:"file_id"`
		} `json:"image_file"`
	} `json:"content"`
	FileIds     []interface{} `json:"file_ids"`
	AssistantID string        `json:"assistant_id"`
	RunID       string        `json:"run_id"`
	Metadata    struct{}      `json:"metadata"`
}

type Annotation struct {
	Type         string `json:"type"`
	Text         string `json:"text"`
	FileCitation struct {
		FileID string `json:"file_id"`
		Quote  string `json:"quote"`
	} `json:"file_citation"`
	StartIndex int `json:"start_index"`
	EndIndex   int `json:"end_index"`
}

func (mr MessageResponse) Dump() string {
	js, err := json.MarshalIndent(&mr, "", "\t")
	if err != nil {
		return err.Error()
	}
	return string(js)
}

type MessageListResponse struct {
	Object  string            `json:"object"`
	Data    []MessageResponse `json:"data"`
	FirstID string            `json:"first_id"`
	LastID  string            `json:"last_id"`
	HasMore bool              `json:"has_more"`
}

func (ml MessageListResponse) Dump() string {
	js, err := json.MarshalIndent(&ml, "", "\t")
	if err != nil {
		return err.Error()
	}
	return string(js)
}
