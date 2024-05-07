package leetcodego

func numRescueBoats(people []int, limit int) int {
    //TC O(nlogn), SC O(1)
    sort.Ints(people)
    ans, left, right := 0, 0, len(people) - 1
    
    for left <= right{
        if people[left] + people[right] <= limit{
            left++
        }
        right--
        ans++
    }
    return ans
}