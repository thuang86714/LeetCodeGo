/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.
Example:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Example 2:

Input: head = [1], n = 1
Output: []
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := ListNode{0, head}
    left, right := &dummy, &dummy
    for i:=0; i < n; i++{
        right = right.Next
    }
    for right.Next !=nil{
        left = left.Next
        right = right.Next
    }
    next_one := left.Next
    left.Next = next_one.Next
    return dummy.Next
}