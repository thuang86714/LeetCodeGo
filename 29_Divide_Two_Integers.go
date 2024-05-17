/*
Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.

Return the quotient after dividing dividend by divisor.

Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For this problem, if the quotient is strictly greater than 231 - 1, then return 231 - 1, and if the quotient is strictly less than -231, then return -231.

 

Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Explanation: 10/3 = 3.33333.. which is truncated to 3.*/
package leetcodego
import(
    "math"
)
func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1{
        return math.MaxInt32
    }
    quotient := 0
    divid := abs(dividend)
    divs := abs(divisor)
    for divid >= divs {
        sub := divs
        add := 1
        for divid >= sub<<1{
            sub <<= 1
            add <<= 1
        }
        divid -= sub
        quotient += add
        /*
        For anyone confused about the bitshifts (<<), the idea is that it makes the subtractions and additions faster by multiplying each value by a higher power of 2 with each shift, requiring fewer iterations. This is logically equivalent to this more intuitive code but much faster:
        */
    }
    sign := (dividend < 0) == (divisor < 0)
    if sign{
        return quotient 
    }
    return -quotient
}
