/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcodego
 func reverseBetween(head *ListNode, left int, right int) *ListNode {
    //TC O(n), SC O(n)
    dummy := &ListNode{0, head}
    pLeft, prev := dummy, dummy
    nodeSlice := []*ListNode{}
    for i := 0; i < left; i++{
        prev = pLeft
        pLeft = pLeft.Next
    }
    for i:= left; i <= right; i++{
        nodeSlice = append(nodeSlice, pLeft)
        pLeft = pLeft.Next
    }
    lIdx, rIdx := 0, len(nodeSlice) - 1
    for lIdx <= rIdx{
        nodeSlice[lIdx], nodeSlice[rIdx] = nodeSlice[rIdx], nodeSlice[lIdx]
        lIdx++
        rIdx--
    }
    for _, node := range nodeSlice{
        prev.Next = node
        prev = prev.Next
    }
    prev.Next = pLeft
    return dummy.Next

}