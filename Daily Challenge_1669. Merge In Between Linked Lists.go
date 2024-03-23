type ListNode struct {
    Val int
    Next *ListNode
}

 func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    //TC O(n), SC O(1)
    dummy := &ListNode{0, list1}
    left, right, cur, nodeCount := dummy, dummy, list2, -1
    var prev *ListNode
    var nxt *ListNode
    for nodeCount < b{
        if nodeCount < a{
            prev = left
            left = left.Next
        }
        if nodeCount < b{
            right = right.Next
        }
        nodeCount++
    }
    nxt = right.Next
    prev.Next = list2
    for cur.Next != nil{
        cur = cur.Next
    }
    cur.Next = nxt

    return dummy.Next
}