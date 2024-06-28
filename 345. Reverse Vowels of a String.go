//credit to tr1nity
package leetcodego
import(
    "unicode"
)
func isVowel(c rune) bool{
    c = unicode.ToLower(c)
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
func reverseVowels(s string) string {
    chars := []rune(s)
    left, right := 0, len(chars) - 1
    for left < right{
        for left < right && !isVowel(chars[left]){
            left++
        }
        for left < right && !isVowel(chars[right]){
            right--
        }
        chars[left], chars[right] = chars[right], chars[left]
        left++
        right--
    }
    return string(chars)
}