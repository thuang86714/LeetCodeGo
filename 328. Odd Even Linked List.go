package leetcodego
func oddEvenList(head *ListNode) *ListNode {
	//TC O(n), SC O(1)
    if head == nil || head.Next == nil {
        return head
    }

    // Initialize dummy heads for odd and even lists
    oddDummy := &ListNode{}
    evenDummy := &ListNode{}

    // Initialize current pointers for odd and even lists
    oddCurrent, evenCurrent := oddDummy, evenDummy
    isOdd := true // Track whether the current node is odd or even

    // Iterate through the original list
    for head != nil {
        if isOdd {
            // Add the node to the odd list
            oddCurrent.Next = head
            oddCurrent = oddCurrent.Next
        } else {
            // Add the node to the even list
            evenCurrent.Next = head
            evenCurrent = evenCurrent.Next
        }
        head = head.Next // Move to the next node in the original list
        isOdd = !isOdd   // Switch between odd and even
    }

    // Connect the odd list to the even list
    oddCurrent.Next = evenDummy.Next
    evenCurrent.Next = nil // Terminate the even list

    return oddDummy.Next // The head of the rearranged list is the first odd node
}
