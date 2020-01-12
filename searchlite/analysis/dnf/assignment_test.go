package dnf

import (
	"fmt"
	"testing"
)

func TestAssignment_String(t *testing.T) {
	a := Assignment{
		Attribute: 1,
		Belong:    true,
		Value:     1,
		attrName:  "age",
		valueName: "[10 to 19]",
	}

	fmt.Printf("%s\n", a)
	a.Belong = false
	fmt.Printf("%s\n", a)
}


