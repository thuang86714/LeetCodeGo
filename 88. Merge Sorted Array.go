package leetcodego

func merge(nums1 []int, m int, nums2 []int, n int)  {
    //TC O(m+n), SC O(1)
    i, j, k := m - 1, n - 1, m + n - 1
    //traverse from the back of slice(from high val to low val)
    for ; i >= 0 &&  j >= 0 && k >= 0; k--{
        if nums1[i] > nums2[j]{
            nums1[k] = nums1[i]
            i--
        }else{
            nums1[k] = nums2[j]
            j--
        }
    }

    for ; i >= 0 && k >= 0;{
        nums1[k] = nums1[i]
        k--
        i--
    } 

    for ; j >= 0 && k >= 0;{
        nums1[k] = nums2[j]
        k--
        j--
    }
}