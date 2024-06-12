package leetcodego

func sortColors2(nums []int)  {
    //Dutch National Flag algo, TC O(n), SC O(1)
    low, mid, high := 0, 0, len(nums)-1
    for mid <= high{
        if nums[mid] == 0{
            nums[mid], nums[low] = nums[low], nums[mid]
            low++
            mid++
        }else if nums[mid]==1{
            mid++
        }else if nums[mid]==2{
            nums[mid], nums[high] = nums[high], nums[mid]
            //mid++
            high--
        }
    }
}

