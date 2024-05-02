package leetcodego

import(
	"unicode",
    "regexp",
    "strings",
)

func isPalindrome(s string) bool {
    newStr := strings.ToLower(s)
    left, right := 0, len(s) - 1
    for left < right{
        for left < right && !unicode.IsLetter(rune(newStr[left])) && !unicode.IsNumber(rune(newStr[left])){
            left++
        }
        for left < right && !unicode.IsLetter(rune(newStr[right])) && !unicode.IsNumber(rune(newStr[right])){
            right--
        }
        if newStr[left] != newStr[right]{
            return false
        }
        left++
        right--
    }
    return true
}

//credit to golang official


func reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
        }
    return string(r)
}

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    re, _ := regexp.Compile(`[^a-zA-Z0-9]+`)
    new_text := re.ReplaceAllString(s, "")
    return new_text == reverse(new_text)
}