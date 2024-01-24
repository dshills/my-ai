package assistant

import (
	"strings"
	"time"

	"github.com/dshills/my-ai/api"
)

type Thread interface {
	SendMessage(string) error
}

type _thread struct {
	ID          string
	Object      string
	CreatedAt   int
	Metadata    map[string]string
	apiKey      string
	assistantID string
	Content     []api.MessageResponse
	runQueue    []string
	msgStore    messageStore
	data        chan runData
	done        chan bool
}

func (t *_thread) create(apikey, assistantid string, data chan runData, done chan bool) error {
	t.apiKey = apikey
	t.assistantID = assistantid
	t.data = data
	t.done = done
	t.msgStore = newMessageStore()

	api := api.New(apikey)
	ctresp, err := api.CreateThread()
	if err != nil {
		return err
	}
	t.ID = ctresp.ID
	t.Object = ctresp.Object
	t.CreatedAt = ctresp.CreatedAt
	t.Metadata = ctresp.Metadata
	go t.monitorRuns()

	return nil
}

func (t *_thread) SendMessage(line string) error {
	msg := api.MessageRequest{
		Role:    "user",
		Content: line,
	}
	api := api.New(t.apiKey)
	content, err := api.PostMessage(t.ID, msg)
	if err != nil {
		rd := runData{Error: err.Error()}
		t.data <- rd
		return err
	}
	t.Content = append(t.Content, *content)
	t.run()
	return nil
}

func (t *_thread) run() error {
	rr := api.RunRequest{
		AssistantID: t.assistantID,
	}
	api := api.New(t.apiKey)
	runresp, err := api.Run(t.ID, rr)
	if err != nil {
		rd := runData{Error: err.Error()}
		t.data <- rd
		return err
	}
	t.runQueue = append(t.runQueue, runresp.ID)
	t.checkRunStatus()
	return nil
}

func (t *_thread) getResponse(runid string) (string, error) {
	api := api.New(t.apiKey)
	msgs, err := api.RunResponse(t.ID)
	if err != nil {
		return "", err
	}
	t.msgStore.AddMessageList(msgs)
	cl := t.msgStore.GetSortedUndisplayed()
	builder := strings.Builder{}
	for _, con := range cl {
		for _, msg := range con.Content {
			builder.WriteString(msg.Text.Value + "\n")
		}
		t.msgStore.SetDisplayed(con.ID)
	}
	if msgs.HasMore {
		return builder.String(), nil
	}
	t.runQueue = t.runQueue[1:]
	return builder.String(), nil
}

type runData struct {
	Message string
	Error   string
}

func (t *_thread) monitorRuns() error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-t.done:
			return nil
		case <-ticker.C:
			if len(t.runQueue) == 0 {
				continue
			}
			t.checkRunStatus()
		}
	}
}

func (t *_thread) checkRunStatus() {
	api := api.New(t.apiKey)
	isComplete, err := api.RunStatus(t.ID, t.runQueue[0])
	switch {
	case err != nil:
		rd := runData{Error: err.Error()}
		t.data <- rd
	case isComplete:
		msg, err := t.getResponse(t.runQueue[0])
		rd := runData{Message: msg}
		if err != nil {
			rd.Error = err.Error()
		}
		t.data <- rd
	}
}
