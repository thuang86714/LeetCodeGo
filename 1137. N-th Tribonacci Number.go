package leetcodego

func tribonacci2(n int) int {
    //credit to srgbl, TC O(n), SC O(1)
    triset := [3]int{0, 1, 1}
    if n < 3 {
        return triset[n]
    }
    for i := 3; i <= n; i++ {
        triset[i % 3] = triset[0] + triset[1] + triset[2]
    }
    return triset[n % 3]
}

func tribonacci(n int) int {
    //TC O(n), SC O(1)
    if n <= 2{
        if n == 0{
            return 0
        }else{
            return 1
        }
    }

    threeStep, twoStep, oneStep, ans := 0, 1, 1, 0
    for i := 3; i <= n; i++{
        ans = threeStep + twoStep + oneStep
        threeStep = twoStep
        twoStep = oneStep
        oneStep = ans
    }
    return ans
}