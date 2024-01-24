package assistant

import (
	"sort"
	"sync"

	"github.com/dshills/my-ai/api"
)

type contentList []api.MessageResponse

func (cl contentList) Len() int           { return len(cl) }
func (cl contentList) Swap(i, j int)      { cl[i], cl[j] = cl[j], cl[i] }
func (cl contentList) Less(i, j int) bool { return cl[i].CreatedAt < cl[j].CreatedAt }

type messageStore struct {
	messages  contentList
	displayed map[string]bool
	m         sync.RWMutex
}

func (ms *messageStore) GetSortedUndisplayed() contentList {
	cl := contentList{}
	for _, m := range ms.messages {
		if !ms.HasDisplayed(m.ID) {
			cl = append(cl, m)
		}
	}
	sort.Sort(cl)
	return cl
}

func (ms *messageStore) HasDisplayed(msgid string) bool {
	ms.m.RLock()
	defer ms.m.RUnlock()
	_, ok := ms.displayed[msgid]
	if ok {
		return true
	}
	return false
}

func (ms *messageStore) SetDisplayed(msgid string) {
	ms.m.Lock()
	defer ms.m.Unlock()
	ms.displayed[msgid] = true
}

func (ms *messageStore) AddContentList(cl contentList) {
	ms.messages = append(ms.messages, cl...)
}

func (ms *messageStore) AddMessageList(ml *api.MessageListResponse) {
	cl := contentList{}
	for _, con := range ml.Data {
		if con.Role == "assistant" {
			cl = append(cl, con)
		}
	}
	ms.AddContentList(cl)
}

func newMessageStore() messageStore {
	return messageStore{displayed: make(map[string]bool)}
}
