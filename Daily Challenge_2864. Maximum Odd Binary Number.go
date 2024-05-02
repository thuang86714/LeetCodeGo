package leetcodego

import (
    "strings"
)

func maximumOddBinaryNumber(s string) string {
    //TC O(n), SC O(1)
    oneCount, zeroCount := 0, 0
    for _, num := range s{
        if num == '1'{
            oneCount++
        }else{
            zeroCount++
        }
    }
    ans := strings.Repeat("1", oneCount - 1) + strings.Repeat("0", zeroCount) + "1"
    return ans
}