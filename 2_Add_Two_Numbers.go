/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
Example:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 package leetcodego
 func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
    carry, dummy := 0, new(ListNode);//new() would return pointer
    for node:= dummy; l1 != nil || l2 != nil || carry >0; node = node.Next{
        if l1 != nil{
            carry += l1.Val
            l1 = l1.Next
        }
        if l2 != nil{
            carry += l2.Val
            l2 = l2.Next
        }
        node.Next = &ListNode{carry%10, nil}
        carry /= 10
    }
    return dummy.Next
}