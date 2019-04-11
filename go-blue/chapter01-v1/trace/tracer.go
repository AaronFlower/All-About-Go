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

// New create a tracer to trace info
func New(out io.Writer) Tracer {
	return &tracer{out: out}
}
