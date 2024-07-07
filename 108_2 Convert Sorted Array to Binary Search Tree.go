package leetcodego

func sortedArrayToBST2(nums []int) *TreeNode {
    //TC O(n), SC O(logn)
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := TreeNode{nums[mid], sortedArrayToBST(nums[:mid]), sortedArrayToBST(nums[mid+1:])}
	return &root
}