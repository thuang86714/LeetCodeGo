package leetcodego

func pairSum(head *ListNode) int {
    //TC O(n), SC O(n)
    nodeValSlice := []int{}
    for head != nil{
        nodeValSlice = append(nodeValSlice, head.Val)
        head = head.Next
    }
    maxSum, n := 0, len(nodeValSlice)
    for i:= 0; i <= n/2 - 1; i++{
        maxSum = max(maxSum, nodeValSlice[i] + nodeValSlice[n - 1 - i])
    }
    return maxSum
}