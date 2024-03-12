func minEatingSpeed(piles []int, h int) int {
    //k has an interval [1, *max_element(piles.begin(), miles.end())];
    //need to find the k that KoKo could cosume all bananas just in time
    //TC O(n*logm), where n is the len of piles and m is the max number in piles, SC O(1)
    left, right := 1, 1000000000
    for left <= right{
        k := left + (right - left)/2
        sumTime := 0
        for _, banana := range piles{
            sumTime += (banana + k - 1)/k//there's no in-build ceil() for int in golang
        }
        if sumTime <= h{//KoKo eats too slow, needs to slow down
            right = k - 1
        }else{
            left = k + 1//KoKo eats too slow, needs to speed up
        }
    }
    return left
}