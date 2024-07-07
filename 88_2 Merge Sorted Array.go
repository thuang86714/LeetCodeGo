package leetcodego

func merge3(nums1 []int, m int, nums2 []int, n int)  {
    idx1, idx2, curIdx := m - 1, n - 1, m + n - 1
    for ; idx1 >= 0 && idx2 >= 0 && curIdx >= 0; curIdx--{
        if nums1[idx1] > nums2[idx2]{
            nums1[curIdx] = nums1[idx1]
            idx1--
        }else{
            nums1[curIdx] = nums2[idx2]
            idx2--
        }
    }

    for idx1 >= 0 && curIdx >= 0{
        nums1[curIdx] = nums1[idx1]
        curIdx--
        idx1--
    }

    for idx2 >= 0 && curIdx >= 0{
        nums1[curIdx] = nums2[idx2]
        curIdx--
        idx2--
    }
}