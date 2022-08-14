package news_manager

import (
	"errors"
	om "github.com/shshimamo/delinkcious/pkg/object_model"
)

const inMemoryMaxPageSize = 10

type userEvents map[string][]*om.LinkManagerEvent

type inMemoryNewsStore struct {
	userEvents userEvents
}

func (m *inMemoryNewsStore) GetNews(username string, startIndex int) (events []*om.LinkManagerEvent, nextIndex int, err error) {
	userEvents := m.userEvents[username]
	if startIndex > len(userEvents) {
		err = errors.New("Index out of bounds")
		return
	}

	pageSize := len(userEvents) - startIndex
	if pageSize > inMemoryMaxPageSize {
		pageSize = inMemoryMaxPageSize
		nextIndex = startIndex + inMemoryMaxPageSize
	} else {
		nextIndex = -1
	}

	events = userEvents[startIndex : startIndex+pageSize]
	return
}

func (m *inMemoryNewsStore) AddEvent(username string, event *om.LinkManagerEvent) (err error) {
	if username == "" {
		err = errors.New("user name can't be empty")
		return
	}

	if event == nil {
		err = errors.New("event can't be nil")
		return
	}

	if m.userEvents[username] == nil {
		m.userEvents[username] = []*om.LinkManagerEvent{}
	}

	m.userEvents[username] = append(m.userEvents[username], event)
	return
}

func NewInMemoryNewsStore() *inMemoryNewsStore {
	return &inMemoryNewsStore{userEvents{}}
}
