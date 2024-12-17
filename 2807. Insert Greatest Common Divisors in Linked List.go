/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// GCD calculates the Greatest Common Divisor of two integers
package leetcodego
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return abs(a)
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		temp := &ListNode{Val: GCD(cur.Val, cur.Next.Val), Next: cur.Next}
		cur.Next = temp
		cur = temp.Next
	}
	return head
}