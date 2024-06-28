/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 package leetcodego
 func deleteMiddle(head *ListNode) *ListNode {
    //TC O(n), SC O(1)
    dummy := &ListNode{0, head}
    slow, fast := dummy, dummy
    for fast.Next != nil && fast.Next.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next
    }
    slow.Next = slow.Next.Next
    return dummy.Next
}