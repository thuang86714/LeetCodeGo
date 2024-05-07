package leetcodego

func deleteNode2(node *ListNode) {
    //TC O(n), SC O(1)
    for node != nil{
        nextVal := node.Next.Val
        node.Val = nextVal
        if node.Next.Next == nil{
            node.Next = nil
        }
        node = node.Next
    }
}