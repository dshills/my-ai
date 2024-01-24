package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const openaiBaseURL = "https://api.openai.com/v1"

const (
	EndpointAssistants      = "/assistants"
	EndpointThreads         = "/threads"
	EndpointThreadMessages  = "/threads/%THREADID%/messages"
	EndpointThreadRuns      = "/threads/%THREADID%/runs"
	EndpointThreadRunStatus = "/threads/%THREADID%/runs/%RUNID%"
)

type Client struct {
	apiKey string
	client http.Client
}

func New(apikey string) Client {
	return Client{apiKey: apikey}
}

func (a Client) send(method, endpoint string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("send: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.apiKey))
	req.Header.Add("OpenAI-Beta", "assistants=v1")

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Create assistant: %w", err)
	}
	defer resp.Body.Close()

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Create assistant: %w", err)
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("send: %v", resp.StatusCode)
	}

	return respbody, err
}

func (a Client) CreateAssistant(acr AssistantCreateRequest) (*AssistantCreateResponse, error) {
	js, err := json.MarshalIndent(&acr, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("Create assistant: %w", err)
	}
	apiurl := fmt.Sprintf("%v%v", openaiBaseURL, EndpointAssistants)
	body, err := a.send("POST", apiurl, strings.NewReader(string(js)))
	if err != nil {
		return nil, fmt.Errorf("Create assistant: %w", err)
	}
	acresp := AssistantCreateResponse{}
	if err := json.Unmarshal(body, &acresp); err != nil {
		return nil, fmt.Errorf("Create assistant: %w", err)
	}
	return &acresp, nil
}

func (a Client) CreateThread() (*CreateThreadResponse, error) {
	apiurl := fmt.Sprintf("%v%v", openaiBaseURL, EndpointThreads)
	body, err := a.send("POST", apiurl, nil)
	if err != nil {
		return nil, fmt.Errorf("Create thread: %w", err)
	}
	ctresp := CreateThreadResponse{}
	if err := json.Unmarshal(body, &ctresp); err != nil {
		return nil, fmt.Errorf("Create thread: %w", err)
	}
	return &ctresp, nil
}

func (a Client) PostMessage(threadid string, msg MessageRequest) (*MessageResponse, error) {
	js, err := json.MarshalIndent(&msg, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("Post message: %w", err)
	}
	ep := EndpointThreadMessages
	ep = strings.Replace(ep, "%THREADID%", threadid, 1)
	apiurl := fmt.Sprintf("%v%v", openaiBaseURL, ep)
	body, err := a.send("POST", apiurl, strings.NewReader(string(js)))
	if err != nil {
		return nil, fmt.Errorf("Post message: %w", err)
	}
	msgresp := MessageResponse{}
	if err := json.Unmarshal(body, &msgresp); err != nil {
		return nil, fmt.Errorf("Post message: %w", err)
	}
	return &msgresp, nil
}

func (a Client) Run(threadid string, rr RunRequest) (*RunResponse, error) {
	js, err := json.MarshalIndent(&rr, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("Run: %w", err)
	}
	ep := EndpointThreadRuns
	ep = strings.Replace(ep, "%THREADID%", threadid, 1)
	apiurl := fmt.Sprintf("%v%v", openaiBaseURL, ep)
	body, err := a.send("POST", apiurl, strings.NewReader(string(js)))
	if err != nil {
		return nil, fmt.Errorf("Run: %w", err)
	}
	runresp := RunResponse{}
	if err := json.Unmarshal(body, &runresp); err != nil {
		return nil, fmt.Errorf("Run: %w", err)
	}
	return &runresp, nil
}

func (a Client) RunStatus(threadid, runid string) (bool, error) {
	ep := EndpointThreadRunStatus
	ep = strings.Replace(ep, "%THREADID%", threadid, 1)
	ep = strings.Replace(ep, "%RUNID%", runid, 1)
	apiurl := fmt.Sprintf("%v%v", openaiBaseURL, ep)
	body, err := a.send("GET", apiurl, nil)
	if err != nil {
		return false, fmt.Errorf("RunStatus: %w", err)
	}
	runresp := RunStatusResponse{}
	if err := json.Unmarshal(body, &runresp); err != nil {
		return false, fmt.Errorf("RunStatus: %w", err)
	}
	if runresp.Status == "error" || (runresp.LastError != nil && *runresp.LastError != "") {
		return false, fmt.Errorf("%v", *runresp.LastError)
	}
	if runresp.Status == "completed" {
		return true, nil
	}
	return false, nil
}

func (a Client) RunResponse(threadid string) (*MessageListResponse, error) {
	ep := EndpointThreadMessages
	ep = strings.Replace(ep, "%THREADID%", threadid, 1)
	apiurl := fmt.Sprintf("%v%v", openaiBaseURL, ep)
	body, err := a.send("GET", apiurl, nil)
	if err != nil {
		return nil, fmt.Errorf("RunResponse: %w", err)
	}
	msgresp := MessageListResponse{}
	if err := json.Unmarshal(body, &msgresp); err != nil {
		return nil, fmt.Errorf("RunResponse: %w", err)
	}
	return &msgresp, nil
}
