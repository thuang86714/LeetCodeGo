package leetcodego

var keyboard = map[rune]string{
    '2':"abc", '3' :"def", '4': "ghi", '5':"jkl", '6':"mno", '7':"pqrs", '8':"tuv", '9':"wxyz",
}

func letterCombinations(digits string) []string {
    var result []string
    if len(digits) == 0{
        return result
    }
    findCombo(digits, "", 0, &result)
    return result
}

func findCombo(digits, curString string, idx int, res *[]string){
    if idx == len(digits){
        *res = append(*res, curString)
        return
    }

    for _, char := range keyboard[rune(digits[idx])]{
        findCombo(digits, curString + string(char), idx + 1, res)
    }
}