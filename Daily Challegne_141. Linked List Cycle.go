/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 package leetcodego
 func hasCycle(head *ListNode) bool {
    //TC O(n), SC O(1)
    if head == nil{
        return false
    }
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast{
            return true
        }
    }
    return false;
}