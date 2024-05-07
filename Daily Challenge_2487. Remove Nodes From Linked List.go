package leetcodego

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNodes(head *ListNode) *ListNode {
    //TC O(n), SC O(n)
    curMaxValSlice := []int{}
    dummy := ListNode{Val:0, Next: head}
    dummyPtr, prevPtr := &dummy, &dummy
    curMax, idx := 0, 0
    //record the val of each node with curMaxValSlice
    for head != nil{
        curMaxValSlice = append(curMaxValSlice, head.Val)
        head = head.Next
    }

    //from tail to head, compare to get the local max val 
    for i := len(curMaxValSlice) - 1; i >= 0; i--{
        curMax = max(curMax, curMaxValSlice[i])
        curMaxValSlice[i] = curMax
    }

    //if cur node val < curMaxVal, then prevPtr.Next = prevPtr.Next.Next
    for prevPtr.Next != nil{
        if prevPtr.Next.Val < curMaxValSlice[idx]{
            prevPtr.Next = prevPtr.Next.Next
        }else if prevPtr.Next.Val == curMaxValSlice[idx]{
            prevPtr = prevPtr.Next
        }
        idx++
    }
    return dummyPtr.Next
}

/*
TC O(n), SC O(1)
 func removeNodes(head *ListNode) *ListNode {
    head = reverseList(head)
    temp := head
    for temp != nil {
        if temp.Next != nil && temp.Val > temp.Next.Val {
            temp.Next = temp.Next.Next
        } else {
            temp = temp.Next
        }
    }
    return reverseList(head)
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        next := head.Next
        head.Next = prev
        prev, head = head, next
    }
    return prev
}
*/