package linklist

import "testing"

func TestNilList(t *testing.T) {
	list := newList([]int{})
	if list != nil {
		t.Error("Expect newList return nil")
	}
}

func TestList123(t *testing.T) {
	data := []int{1, 2, 3}
	list := newList(data)
	plist := list
	for _, v := range data {
		if v != plist.data {
			t.Errorf(" Expect %d = %d ", v, plist.data)
		}
		plist = plist.next
	}
}

func TestListReverse(t *testing.T) {
	data := []int{1, 2, 3}
	list := newList(data)
	plist := list

	for _, v := range data {
		if v != plist.data {
			t.Errorf(" Expect %d = %d ", v, plist.data)
		}
		plist = plist.next
	}

	rdata := []int{3, 2, 1}
	rlist := reverse(list)
	for _, v := range rdata {
		if v != rlist.data {
			t.Errorf(" Expect %d = %d ", v, rlist.data)
		}
		rlist = rlist.next
	}
}
