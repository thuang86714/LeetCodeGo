package leetcodego

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList3(head *ListNode) *ListNode {
    //TC O(n), SC O(n)
    if head == nil{
        return nil
    }
    tempNodeSlice := []*ListNode{}
    for head != nil{
        tempNodeSlice = append(tempNodeSlice, head)
        head = head.Next
    }

    lens := len(tempNodeSlice)

    for i := lens - 1; i >= 1; i--{
        tempNodeSlice[i].Next = tempNodeSlice[i - 1]
    } 
    tempNodeSlice[0].Next = nil// Fix to properly terminate the list

    return tempNodeSlice[lens - 1]// Return new head, which is the last element
}

func reverseList(head *ListNode) *ListNode {
	//TC O(n), SC O(1)
    var prev *ListNode = nil
    current := head

    for current != nil {
        nextTemp := current.Next // Store next node
        current.Next = prev      // Reverse current node's pointer
        prev = current           // Move pointers one position ahead
        current = nextTemp
    }
    return prev // New head of the reversed list
}
