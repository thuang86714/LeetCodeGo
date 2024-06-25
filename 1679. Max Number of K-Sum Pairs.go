package leetcodego

func maxOperations(nums []int, k int) int {
    //TC O(n), SC O(n)
    dict := make(map[int]int)
    for _, num := range nums{
        dict[num]++
    }
    ans := 0
    for _, num := range nums{
        if dict[num] > 0{
            dict[num]--
            if dict[k - num] > 0{
                ans++
                dict[k - num]--
            }else{
                dict[num]++
            }
        }
    }
    return ans
}

func maxOperations(nums []int, k int) int {
    // Creatjing a map to count occurrences of each number
    countMap := make(map[int]int)
    for _, num := range nums {
        countMap[num]++
    }

    operations := 0
    for num, count := range countMap {
        if count == 0 || countMap[k-num] == 0 {
            continue // Skip if no pairs can be formed
        }

        if num*2 == k { // Special case for pairs of the same number
            operations += count / 2
        } else { // General case
            pairs := min(count, countMap[k-num])
            operations += pairs
            // Update counts to reflect pairs formation
            countMap[num] -= pairs
            countMap[k-num] -= pairs
        }
    }

    return operations
}