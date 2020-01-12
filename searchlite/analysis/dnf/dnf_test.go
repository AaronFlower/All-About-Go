package dnf

import (
	"fmt"
	"testing"
)

func TestSecondConjunctionList(t *testing.T) {

}

func TestAd(t *testing.T) {
	assign1 := Assignment{
		Attribute: 1,
		Belong:    true,
		Value:     1,
	}

	assign2 := Assignment{
		Attribute: 2,
		Belong:    false,
		Value:     1,
	}

	conj1 := Conjunction{assignments: []Assignment{assign1}}
	conj2 := Conjunction{assignments: []Assignment{assign2}}
	conj3 := Conjunction{assignments: []Assignment{assign1, assign2}}

	dnf1 := DNF{conjunctions: []Conjunction{conj1, conj2, conj3}}
	ad1 := Ad{
		Id:  1,
		dnf: dnf1,
	}
	fmt.Printf("%v", ad1)
}
