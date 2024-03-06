/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

 func copyRandomList(head *Node) *Node {
    //map + lined list, TC O(n), SC O(n)
    if head == nil{
        return head
    }
    nodeMap := make(map[*Node]*Node)//<old node, new node>

    pOld := head
    for pOld != nil{
        nodeMap[pOld] = &Node{Val: pOld.Val}
        pOld = pOld.Next
    }
    pOld = head
    for pOld != nil{
        nodeMap[pOld].Next = nodeMap[pOld.Next]
        nodeMap[pOld].Random = nodeMap[pOld.Random]
        pOld = pOld.Next
    }
    return nodeMap[head]
}