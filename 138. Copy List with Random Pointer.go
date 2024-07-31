package leetcodego

// Node is the definition for a linked list node with a random pointer
type Node3 struct {
    Val    int
    Next   *Node3
    Random *Node3
}

// copyRandomList copies a linked list with random pointers
func copyRandomList(head *Node3) *Node3 {
    // map + linked list, TC O(n), SC O(n)
    if head == nil {
        return head
    }

    nodeMap := make(map[*Node3]*Node3) // <old node, new node>

    pOld := head
    // First pass: copy all the nodes
    for pOld != nil {
        nodeMap[pOld] = &Node3{Val: pOld.Val}
        pOld = pOld.Next
    }

    pOld = head
    // Second pass: assign Next and Random pointers
    for pOld != nil {
        nodeMap[pOld].Next = nodeMap[pOld.Next]
        nodeMap[pOld].Random = nodeMap[pOld.Random]
        pOld = pOld.Next
    }

    return nodeMap[head]
}
