package leetcodego

// reverseList reverses the singly linked list.
func reverseList2(node *ListNode) *ListNode {
    var prev *ListNode
    curr := node
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}

// doubleIt doubles the integer stored in each node of the linked list.
func doubleIt(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    cur := reverseList(head)
    reverseNode := cur
    carry := 0
    for cur != nil {
        newVal := cur.Val*2 + carry
        cur.Val = newVal % 10
        carry = newVal / 10
        if cur.Next == nil {
            break
        }
        cur = cur.Next
    }
    if carry > 0 {
        cur.Next = &ListNode{Val: carry}
    }
    return reverseList(reverseNode)
}