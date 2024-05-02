package leetcodego

func isPalindrome(x int) bool {
    if x < 0 || (x > 0 && x%10 == 0){
        return false
    }
    sum := 0
    for x > sum{
        sum = sum*10 + x%10
        x /= 10
    }
    return (x == sum) || (sum/10 == x)
}