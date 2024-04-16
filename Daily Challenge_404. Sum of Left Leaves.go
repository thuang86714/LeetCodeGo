package leetcodego

// Pair holds a TreeNode and a boolean to indicate if it is a left child.
type pair struct {
    Node   *TreeNode
    IsLeft bool
}

// sumOfLeftLeaves calculates the sum of left leaves in the binary tree.
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }

    nodeSlice := []pair{{Node: root, IsLeft: false}}
    sum := 0

    for len(nodeSlice) > 0 {
        // Pop the first element from the slice
        curPair := nodeSlice[0]
        nodeSlice = nodeSlice[1:]

        // Process the current node
        if curPair.Node != nil {
            // Check if it's a leaf node and a left child
            if curPair.Node.Left == nil && curPair.Node.Right == nil && curPair.IsLeft {
                sum += curPair.Node.Val
            }

            // Append right and left children to the slice
            if curPair.Node.Right != nil {
                nodeSlice = append(nodeSlice, pair{Node: curPair.Node.Right, IsLeft: false})
            }
            if curPair.Node.Left != nil {
                nodeSlice = append(nodeSlice, pair{Node: curPair.Node.Left, IsLeft: true})
            }
        }
    }

    return sum
}

/*
//DFS
// findAllLeft performs a DFS to find and sum all left leaves.
func (s *Solution) findAllLeft(node *TreeNode, isLeft bool) {
    if node == nil {
        return
    }

    if node.Left == nil && node.Right == nil && isLeft {
        s.sum += node.Val
        return
    }

    s.findAllLeft(node.Right, false)
    s.findAllLeft(node.Left, true)
}

// sumOfLeftLeaves initiates the DFS and returns the sum of all left leaves.
func (s *Solution) sumOfLeftLeaves(root *TreeNode) int {
    s.findAllLeft(root, false)
    return s.sum
}
*/