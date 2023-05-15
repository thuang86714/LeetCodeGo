/*
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

 

Example 1:

Input: x = 2.00000, n = 10
Output: 1024.00000
*/
func myPow(x float64, n int) float64 {
    return backtrack(x, n, 1)
}

func backtrack(x float64, n int, res float64) float64{
    if n==0{
        return res
    }

    if n%2 != 0{
        if n >0{
            res *= x
        }else{
            res /= x
        }
    }

    return backtrack(x*x, n/2, res)
}