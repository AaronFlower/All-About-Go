package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Errorf("the New method should not return nil")
	} else {
		tracer.Trace("Hello Tracer.")
		if buf.String() != "Hello Tracer.\n" {
			t.Errorf("The tracer should not returns '%s'\n", buf.String())
		}
	}
}
