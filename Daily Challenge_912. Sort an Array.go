//Heap Sort TC O(nlogn), SC O(1)
//Merge Sort TC O(nlogn), SC O(n)
package leetcodego
func merge2(nums1, nums2 []int) []int {
    result := []int{}
    idx1, idx2 := 0, 0
    for idx1 < len(nums1) && idx2 < len(nums2){
        if nums1[idx1] > nums2[idx2]{
            result = append(result, nums2[idx2])
            idx2++
        }else{
            result = append(result, nums1[idx1])
            idx1++
        }
    }

    for idx1 < len(nums1){
        result = append(result, nums1[idx1])
        idx1++
    }
    for idx2 < len(nums2){
        result = append(result, nums2[idx2])
        idx2++
    }
    return result
}
func mergeSort(nums []int) []int {
    if len(nums) == 1{
        return nums
    }
    mid := len(nums)/2
    nums1 := mergeSort(nums[:mid])
    nums2 := mergeSort(nums[mid:])

    return merge2(nums1, nums2)
}
func sortArray(nums []int) []int {
    return mergeSort(nums)
}