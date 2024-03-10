func intersection(nums1 []int, nums2 []int) []int {
    //TC O(m+n), SC O(m+n)
    dict1, dict2 := make(map[int]bool), make(map[int]bool)
    res := []int{}
    for _, num := range nums1{
        dict1[num] = true
    }
    for _, num := range nums2{
        if dict1[num]{
            dict2[num] = true
        }
    }

    for num := range dict2{
        res = append(res, num)
    }
    return res
}