package linklist

type node struct {
	data int
	next *node
}

type linkList *node

func newList(data []int) *node {
	headNode := &node{-1, nil}
	tail := headNode

	for _, v := range data {
		tail.next = &node{v, nil}
		tail = tail.next
	}
	return headNode.next
}

func reverse(list linkList) linkList {
	tail := node{-1, nil}
	p := list
	for p != nil {
		pnext := p.next
		p.next = tail.next
		tail.next = p
		p = pnext
	}
	return tail.next
}
