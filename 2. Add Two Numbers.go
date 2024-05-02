package leetcodego
/*
type ListNode struct {
     Val int
    Next *ListNode
}
*/

 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    dummy := new(ListNode)
    cur := dummy
    for l1 != nil || l2 != nil || carry != 0{
        if l1 != nil{
            carry += l1.Val
            l1 = l1.Next
        }
        if l2 != nil{
            carry += l2.Val
            l2 = l2.Next
        }
        cur.Next = &ListNode{carry%10, nil}
        cur = cur.Next
        carry /= 10
    }
    return dummy.Next
}