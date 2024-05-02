package leetcodego

func romanToInt(s string) int {
    //credit to klakovskiy
    prevVal, curVal, res := 0, 0, 0
    dict := map[uint8]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    for i := len(s) - 1; i >= 0; i--{
        curVal = dict[s[i]]
        if curVal < prevVal{
            res -= curVal
        }else{
            res += curVal
        }
        prevVal = curVal
    }
    return res
}