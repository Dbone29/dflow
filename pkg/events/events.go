package events

import (
	"math"
	"sort"
	"sync"
)

type EventPayload interface{}

type EventFunc func(input EventPayload)

type EventHandler struct {
	Func     EventFunc
	Priority int
}

func (p *EventHandler) Handle(payload EventPayload) {
	p.Func(payload)
}

type Event struct {
	handlers []EventHandler
	mu       sync.Mutex
}

func (e *Event) Register(handler EventHandler) {
	e.mu.Lock()
	defer e.mu.Unlock()

	// check if priority is set, if not set it to max int
	if handler.Priority == 0 {
		handler.Priority = math.MaxInt32 / 2
	}

	e.handlers = append(e.handlers, handler)
	sort.Slice(e.handlers, func(i, j int) bool {
		return e.handlers[i].Priority < e.handlers[j].Priority
	})
}

func (e *Event) Trigger(payload EventPayload) {
	e.mu.Lock()
	defer e.mu.Unlock()
	for _, handler := range e.handlers {
		go handler.Handle(payload)
	}
}
