package leetcodedynamicprogramming

type ListNode struct {
    Val int
    Next *ListNode
}

 func hasCycle(head *ListNode) bool {
    //TC O(n), SC O(1)
    slow, fast := head, head
    for fast != nil && fast.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next
        if fast == slow {
            return true
        }
    }
    return false
}