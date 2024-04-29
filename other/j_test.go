package other

import "testing"

type LinkNode struct {
	next *LinkNode
	val  int
}

func FirstMeet(head1, head2 *LinkNode) *LinkNode {
	if head1 == nil || head2 == nil {
		return nil
	}

	len1 := 0
	len2 := 0
	n := 0
	cur1 := head1
	cur2 := head2

	for cur1.next != nil {
		cur1 = cur1.next
		len1++
	}
	for cur2.next != nil {
		cur2 = cur2.next
		len2++
	}

	//10 7  | -5
	// 5 2 | 5
	if len1 < len2 {
		n = len2 - len1
		for n > 0 {
			cur2 = cur2.next
			n--
		}
	} else {
		n = len1 - len2
		for n > 0 {
			cur1 = cur1.next
			n--
		}
	}

	for cur1.val != cur2.val {
		cur1 = cur1.next
		cur2 = cur2.next
	}

	return cur1
}

func TestJ(t *testing.T) {

}
