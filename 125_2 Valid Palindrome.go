package leetcodego
import(
	"strings"
	"unicode"
)
func isPalindrome5(s string) bool {
	newStr := strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !unicode.IsLetter(rune(newStr[left])) && !unicode.IsNumber(rune(newStr[left])) {
			left++
		}
		for left < right && !unicode.IsLetter(rune(newStr[right])) && !unicode.IsNumber(rune(newStr[right])) {
			right--
		}
		if newStr[left] != newStr[right] {
			return false
		}
		left++
		right--

	}
	return true
}