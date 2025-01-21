package job

import "context"

type Processor interface {
	Process(job Job) error
}

type stringProcessor struct {
	ctx context.Context
}

func NewStringProcessor(ctx context.Context) Processor {
	return &stringProcessor{}
}

func (p *stringProcessor) Process(job Job) error {

	return nil
}
