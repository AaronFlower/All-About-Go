package mymath

import "testing"

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("Division failed!")
	} else {
		t.Log("Passed")
	}
}

func Test_Division_2(t *testing.T) {
	if _, e := Division(6, 0); e == nil {
		t.Error("Division dit not work as expected!")
	} else {
		t.Log("One test passed.", e)
	}
}
