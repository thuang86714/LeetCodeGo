func kidsWithCandies(candies []int, extraCandies int) []bool {
	//TC O(n), SC O(1)
    curMax := slices.Max(candies)
    ans := make([]bool, len(candies))
    for i := 0; i < len(candies); i++{
        if candies[i] + extraCandies >= curMax{
            ans[i] = true
        }
    }
    return ans
}