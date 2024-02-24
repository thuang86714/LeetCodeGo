package leetcodedynamicprogramming

func removeElement(nums []int, val int) int {
	//TC O(n), SC O(1)
    slow := 0
    fmt.Printf("cap = %d, len = %d", cap(nums), len(nums))
	/*
	unanswered:
	with test case nums = [2, 2, 2], the cap(nums) = 4, while len(nums) = 3
	*/
    for i := 0; i < len(nums); i++{
        if nums[i] != val{
            nums[slow] = nums[i]
            slow++
        }
    }
    return slow
}