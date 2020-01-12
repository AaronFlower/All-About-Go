package dnf

import (
	"fmt"
	"strings"
)

// Assignment defines a <attribute, value> assignment. Belong or Not Belong
type Assignment struct {
	Attribute uint
	Belong    bool
	Value     int
	attrName  string
	valueName string
}

func (a Assignment) String() string {
	belong := "∉"
	if a.Belong {
		belong = "∈"
	}
	return fmt.Sprintf("(%s, %s, %s)", a.attrName, belong, a.valueName)
}

// Conj defines a conjunction
type Conjunction struct {
	assignments []Assignment
}

func (c *Conjunction)String() string  {
	switch len(c.assignments) {
	case 0:
		return "[]"
	case 1:
		return c.assignments[0].String()
	default:
		assignStrs := make([]string, len(c.assignments))
		for i, a := range c.assignments {
			assignStrs[i] = a.String()
		}
		return fmt.Sprintf("[%s]", strings.Join(assignStrs, " ∧ "))
	}
}

// Size return the number of Assignment which the belong is true.
func (c *Conjunction)Size() int {
	count := 0
	for _, a := range c.assignments {
		if a.Belong {
			count++
		}
	}
	return count
}

// DNF defines a conjunction
type DNF struct {
	conjunctions []Conjunction
}

// Ad defines a ad
type Ad struct {
	Id  uint
	dnf DNF
}
