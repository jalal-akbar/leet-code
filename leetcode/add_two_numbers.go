package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Value int
	Next  *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	sum := 0
	for cur := result; l1 != nil || l2 != nil || sum != 0; {
		if l1 != nil {
			sum += l1.Value
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Value
			l2 = l2.Next
		}
		cur.Next = &ListNode{Value: sum % 10}
		sum /= 10
		cur = cur.Next
	}
	return result.Next
}
