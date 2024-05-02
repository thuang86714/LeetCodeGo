package leetcodego

func minOperations(nums []int, k int) int {
	for _, num := range nums {
		k ^= num
	}
	return countSetBits(k)
}

func countSetBits(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}