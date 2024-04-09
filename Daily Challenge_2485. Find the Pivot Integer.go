package leetcodego

func pivotInteger(n int) int {
    //binary search, TC O(logn), SC O(1)
    left, right := 1, n
    for left <= right{
        pivot := left + (right - left)/2
        leftSum := (1 + pivot)*pivot/2//sum of all number on the left of pivot and pivot itself
        rightSum := (pivot + n)*(n - pivot + 1)/2//sum of all number on the right of pivot and pivot itself
        if leftSum == rightSum{//given there is at most one solution
            return pivot
        }else if leftSum > rightSum{
            right = pivot - 1
        }else{
            left = pivot + 1
        }
    }
    return -1;
}