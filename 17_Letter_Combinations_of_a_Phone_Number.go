/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example 1:

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]*/
package leetcodego
func letterCombinations2(digits string) []string {
	//an alias for the unsigned integer 8 type ( uint8
    Map := map[byte][]string{
        '2': []string{"a", "b", "c"},
        '3': []string{"d", "e", "f"},
        '4': []string{"g", "h", "i"},
        '5': []string{"j", "k", "l"},
        '6': []string{"m", "n", "o"},
        '7': []string{"p", "q", "r", "s"},
        '8': []string{"t", "u", "v"},
        '9': []string{"w", "x", "y", "z"},
    }
    var result []string
    if digits == ""{
        return result
    }
    backtrack("", 0, &result, Map, digits)
    return result
}

func backtrack(curStr string, idx int, res *[]string, Map map[byte][]string, digits string){
    if idx == len(digits){
        *res = append(*res, curStr)
        return
    }
    for _, ch := range Map[digits[idx]] {
        backtrack(curStr+ch, idx+1, res, Map, digits)
    }
}