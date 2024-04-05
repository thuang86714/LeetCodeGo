package leetcodego
import(
	"unicode"
)
func makeGood(s string) string {
    result := []rune{}
    for _, ch := range s {
        if len(result) > 0 && unicode.ToLower(ch) == unicode.ToLower(result[len(result)-1]) &&
            ch != result[len(result)-1] {
            // Pop the last character if current and last characters are a bad pair.
            result = result[:len(result)-1]
        } else {
            // Otherwise, append the current character to the result.
            result = append(result, ch)
        }
    }
    return string(result)
}