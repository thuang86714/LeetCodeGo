/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func middleNode(head *ListNode) *ListNode {
    //TC O(n), SC O(1)
    if head == nil || head.Next == nil{
        return head
    }
    slow, fast := head, head
    countLen, target := 0, 0
    for fast.Next != nil{
        countLen++
        fast = fast.Next 
    } 
    if countLen%2 == 1{
        target = countLen/2 +1
    }else{
        target = countLen/2
    }
    for i := 0; i < target; i++{
        slow = slow.Next
    }
    return slow
}