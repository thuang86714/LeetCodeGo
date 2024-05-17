/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
*/
package leetcodego

func Min(x, y int) int{
    if x >= y{
        return y
    }else{
        return x
    }
}
func Max(x, y int) int{
    if x <= y{
        return y
    }else{
        return x
    }
}
func maxArea2(height []int) int {
    water, l, r := 0, 0, len(height)-1
    for l < r{
        h := Min(height[r], height[l])
        water = Max(water, (r-l)*h)
        for height[l] <= h && l < r{
            l++
        }
        for height[r] <= h && l < r{
            r--
        }
    }
    
    return water
}