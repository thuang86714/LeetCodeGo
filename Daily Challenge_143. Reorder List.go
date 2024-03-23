package leetcodego 

 func reorderList(head *ListNode)  {
    if head == nil|| head.Next == nil|| head.Next.Next == nil{
        return
    }
    //divide the linked list into 2 halves-->slow fast algo
    slow, fast, prev, temp := head, head, head, head
    for fast != nil && fast.Next != nil{
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    prev.Next = nil

    //reverse the second half-->in place
    prev = nil
    for slow != nil{
        temp = slow.Next
        slow.Next = prev
        prev = slow
        slow = temp
    }

    //iterate both halves and combine into one linked list
    // Now, 'prev' is the head of the reversed second half
    firstHalf, secondHalf := head, prev
    for firstHalf != nil && secondHalf != nil{
        temp1 := firstHalf.Next // Store next node of first half
        temp2 := secondHalf.Next // Store next node of second half

        firstHalf.Next = secondHalf // Link first half node to second half node
        if temp1 != nil { // Ensure we do not link from secondHalf to nil when firstHalf is at the end
            secondHalf.Next = temp1
        }

        // Move pointers forward
        firstHalf = temp1
        secondHalf = temp2
    }
}