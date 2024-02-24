func mySqrt(x int) int {
    l, r := 0, x
    for l < x{
        mid := l + (r - l)/2
        if((mid*mid) <= x && (mid+ 1)*(mid + 1) > x){
            return mid
        }else if mid*mid > x{
            r = mid - 1
        }else{
            l = mid + 1
        }

    }
    return l
}