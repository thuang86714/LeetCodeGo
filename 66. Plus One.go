func plusOne(digits []int) []int {
    length := len(digits)
    if digits[length - 1] != 9{
        digits[length - 1]++
        return digits
    }

    for i := length - 1; i >= 0; i--{
        if i == length - 1{
            digits[i]++
        }

        if digits[i] == 10{
            if i == 0{
                digits[i] = 1
                digits = append(digits, 0)
            }else{
                digits[i] = 0
                digits[i - 1]++
            }
            
        }
    }
    return digits
}