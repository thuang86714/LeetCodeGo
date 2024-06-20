package leetcodego

func minDays(bloomDay []int, m int, k int) int {
    //TC O(n*log(max - min)), SC O(1)
    if len(bloomDay) < m*k{
        return -1
    }
    left, right := 1, int(1e9)

    for _, bloom := range bloomDay{
        left = min(left, bloom)
        right = max(right, bloom)
    }

    canMakeBouquets := func(days int) bool {
        bouquets, flowers := 0, 0
        for _, bloom := range bloomDay {
            if bloom <= days {
                flowers++
                if flowers == k {
                    bouquets++
                    flowers = 0
                }
            } else {
                flowers = 0
            }
        }
        return bouquets >= m
    }

    for left <= right {
        mid := left + (right-left)/2
        if canMakeBouquets(mid) {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return left

}