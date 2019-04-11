package trace

import (
	"fmt"
	"io"
)

// Tracer traces the info
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

type nilTracer struct {
}

func (t *nilTracer) Trace(a ...interface{}) {}

// New create a tracer to trace info
func New(out io.Writer) Tracer {
	return &tracer{out: out}
}

// Off create a tracer that traces nothing.
func Off() Tracer {
	return &nilTracer{}
}
