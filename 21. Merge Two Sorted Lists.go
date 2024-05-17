package leetcodego

 func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
    //credit to dorlib
    dummy := &ListNode{}
    cur := dummy
    for list1 != nil && list2 != nil{
        if list1.Val < list2.Val{
            cur.Next = list1
            list1 = list1.Next
        }else{
            cur.Next = list2
            list2 = list2.Next
        }
        cur = cur.Next
    }

    if list1 == nil{
        cur.Next = list2
    }else{
        cur.Next = list1
    }

    return dummy.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    //credit to brianchiang_tw, recursive
    // Base cases
    if (l1 == nil) && (l2 == nil){
        
        // Both l1 and l2 are empty
        return nil
		
    } else if l1 == nil{
	
        // Only l1 is empty
        return l2
		
    } else if l2 == nil{
	
        // Only l2 is empty
        return l1
    }
    
    
    // General cases
    if l1.Val < l2.Val{
        
        // l1 is smaller than l2
        l1.Next = mergeTwoLists(l1.Next, l2)
        return l1
        
    } else{
        
        // l2 is smaller than l1
        l2.Next = mergeTwoLists(l1, l2.Next)
        return l2
        
    }
    
}