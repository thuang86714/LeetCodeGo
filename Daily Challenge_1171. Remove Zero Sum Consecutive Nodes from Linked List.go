package leetcodego

type ListNode struct {
    Val int
    Next *ListNode
}

 func removeZeroSumSublists(head *ListNode) *ListNode {
    //credit to hiepit, TC O(n), SC O(n)
    dummy := &ListNode{0, head}
    cur := dummy
    prefix := 0
    seen := make(map[int]*ListNode)//<curPrefix, ListNode*>
    for cur != nil{
        prefix = prefix + cur.Val
        seen[prefix] = cur//this way we will only store the last node that have the same prefix sum
        cur = cur.Next
    }
    prefix = 0
    cur = dummy
    for cur != nil{
        prefix = prefix + cur.Val
        cur.Next = seen[prefix].Next//connect to the last node with the same prefix sum
        cur = cur.Next
    }
    return dummy.Next
}