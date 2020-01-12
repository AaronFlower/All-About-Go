package dnf

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConjunction_Size(t *testing.T) {

	age1 := Assignment{
		Attribute: 1,
		Belong:    true,
		Value:     2,
		attrName:  "age",
		valueName: "20s",
	}

	age2 := Assignment{
		Attribute: 1,
		Belong:    true,
		Value:     3,
		attrName:  "age",
		valueName: "30s",
	}

	gender := Assignment{
		Attribute: 2,
		Belong:    true,
		Value:     1,
		attrName:  "gender",
		valueName: "male",
	}

	city := Assignment{
		Attribute: 3,
		Belong:    false,
		Value:     1,
		attrName:  "city",
		valueName: "SH",
	}

	conj := Conjunction{assignments: []Assignment{age1, age2, gender, city}}
	assert.Equal(t, 3, conj.Size())
	fmt.Printf("%s\n", conj.String())
}
