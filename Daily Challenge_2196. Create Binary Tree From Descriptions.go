package leetcodego

type NodePair struct {
    node    *TreeNode
    isChild int
}

func createBinaryTree(descriptions [][]int) *TreeNode {
    valTreeNodeMap := make(map[int]NodePair)
    var ans *TreeNode

    for _, description := range descriptions {
        rootVal := description[0]
        childVal := description[1]
        isLeft := description[2]

        var curRoot *TreeNode
        if _, exists := valTreeNodeMap[rootVal]; !exists {
            valTreeNodeMap[rootVal] = NodePair{node: &TreeNode{Val: rootVal}, isChild: 0}
        }
        curRoot = valTreeNodeMap[rootVal].node

        if _, exists := valTreeNodeMap[childVal]; !exists {
            valTreeNodeMap[childVal] = NodePair{node: &TreeNode{Val: childVal}, isChild: 1}
        } else {
            valTreeNodeMap[childVal] = NodePair{node: valTreeNodeMap[childVal].node, isChild: 1}
        }

        if isLeft == 1 {
            curRoot.Left = valTreeNodeMap[childVal].node
        } else {
            curRoot.Right = valTreeNodeMap[childVal].node
        }
    }

    for _, nodePair := range valTreeNodeMap {
        if nodePair.isChild == 0 {
            ans = nodePair.node
        }
    }

    return ans
}