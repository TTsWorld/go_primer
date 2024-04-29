package other

import "testing"

type ListNode struct {
	data int
	next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	prev := &ListNode{}
	current := head
	for current.next != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	return prev

}

func TestD(t *testing.T) {
	reverseList(&ListNode{})
}
