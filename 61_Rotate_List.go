/*
Given the head of a linked list, rotate the list to the right by k places.
Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 /**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 package leetcodego
func rotateRight(head *ListNode, k int) *ListNode {
    if head ==nil || k==0{
        return head
    }
    size := 1
    tail, cur := head, head
    for tail.Next !=nil{
        size++
        tail = tail.Next
    }
    k %= size
    if k==0{return head}
    for i:=0; i < size-k-1; i++{
        cur = cur.Next
    }
    newhead := cur.Next 
    cur.Next = nil
    tail.Next=head
    return newhead

}