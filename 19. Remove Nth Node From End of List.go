/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcodego
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//TC O(n), SC O(1)
    dummy := ListNode{0, head}
    fast, slow := &dummy, &dummy
    for i:= 0; i < n; i++{
        fast = fast.Next
    }
    for fast.Next != nil{
        fast = fast.Next
        slow = slow.Next
    }
    temp := slow.Next
    slow.Next = temp.Next
    return dummy.Next
}