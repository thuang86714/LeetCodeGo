/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 package leetcodego
 func isPalindrome3(head *ListNode) bool {
    //TC O(n), SC O(1)
    slow, fast, prev, temp := head, head, head, head
    for fast != nil && fast.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next
    }
    prev = slow
    slow = slow.Next
    prev.Next = nil//put an end to the first half

    for slow != nil{//reverse the second half
        temp = slow.Next
        slow.Next = prev
        prev = slow
        slow = temp
    }
    slow = prev//before this line, slow == nil
    for slow != nil{
        if head.Val != slow.Val{
            return false
        }
        head = head.Next
        slow = slow.Next
    }
    return true
}