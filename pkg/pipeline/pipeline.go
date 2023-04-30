package pipeline

import (
	"math"
	"sort"
	"sync"
)

type ProcessorPayload interface{}

type ProcessorFunc func(input ProcessorPayload) (output ProcessorPayload)

type Processor struct {
	Func     ProcessorFunc
	Priority int
}

type Pipeline struct {
	processors []Processor
	mu         sync.Mutex
}

func (p *Pipeline) Register(processor Processor) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// check if priority is set, if not set it to max int
	if processor.Priority == 0 {
		processor.Priority = math.MaxInt32 / 2
	}

	p.processors = append(p.processors, processor)
	sort.SliceStable(p.processors, func(i, j int) bool {
		return p.processors[i].Priority < p.processors[j].Priority
	})
}

func (p *Pipeline) Execute(input ProcessorPayload) ProcessorPayload {
	p.mu.Lock()
	defer p.mu.Unlock()
	output := input
	for _, processor := range p.processors {
		output = processor.Func(output)
	}
	return output
}
