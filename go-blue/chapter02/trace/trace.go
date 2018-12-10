package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describing an object capable of
// tracing events throughout code.
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

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {
}

// New creates a tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// Off creates a tracer that will ignore calls to it.
func Off() Tracer {
	return &nilTracer{}
}
