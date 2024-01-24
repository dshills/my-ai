package assistant

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dshills/goai/api"
)

type Assistant interface {
	Info() string
	Start() error
}

type _assistant struct {
	openAIKey    string
	model        string
	tool         string
	name         string
	instructions string
	AssistantID  string
	threads      []Thread
	activeThread Thread
	data         chan runData
	done         chan bool
}

func (a *_assistant) Info() string {
	return fmt.Sprintf(
		"%q created %q using model %s\nType \"quit\" or \"exit\" to end",
		a.name,
		a.instructions,
		a.model,
	)
}

func (a *_assistant) createThread() error {
	t := _thread{}
	if err := t.create(a.openAIKey, a.AssistantID, a.data, a.done); err != nil {
		return err
	}
	a.threads = append(a.threads, &t)
	a.activeThread = &t
	return nil
}

func (a *_assistant) create() error {
	a.data = make(chan runData)
	a.done = make(chan bool)
	acr := api.AssistantCreateRequest{
		Instructions: a.instructions,
		Name:         a.name,
		Model:        a.model,
	}
	mp := make(map[string]string)
	mp["type"] = ToolCodeInterpreter
	acr.Tools = append(acr.Tools, mp)

	api := api.New(a.openAIKey)
	acresp, err := api.CreateAssistant(acr)
	if err != nil {
		return err
	}

	a.AssistantID = acresp.ID

	return nil
}

func (a *_assistant) stopThreads() {
	for range a.threads {
		a.done <- true
	}
}

func (a *_assistant) Start() error {
	defer a.stopThreads()
	for {
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		line = strings.TrimSpace(line)
		switch line {
		case "":
			continue
		case "quit", "exit":
			return nil
		}
		a.activeThread.SendMessage(line)
		fmt.Println("hmmm...")
		select {
		case rd := <-a.data:
			if rd.Error != "" {
				fmt.Println(rd.Error)
			} else {
				fmt.Println(rd.Message)
			}
		}
	}
}

func NewAssistant(apikey, model, name, instructions string) (Assistant, error) {
	a := _assistant{openAIKey: apikey, model: model, name: name, instructions: instructions}
	if err := a.create(); err != nil {
		return &a, err
	}
	if err := a.createThread(); err != nil {
		return &a, err
	}
	return &a, nil
}
