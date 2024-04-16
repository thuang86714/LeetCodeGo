package leetcodego
import(
	"unicode"
)


func buildString(curString string, idx *int)string{
	result := ""
	for *idx < len(curString) && curString[*idx] != ']'{
		if !unicode.IsDigit(rune(curString[*idx])){
			result += string(curString[*idx])
			*idx++
		}else{
			repeat := 0
			for *idx < len(curString) && unicode.IsDigit(rune(curString[*idx])){
				repeat = repeat*10 + int(curString[*idx] - '0')
				*idx++
			}
			*idx++//skip '['
			tempString := buildString(curString, idx)
			*idx++
			for repeat > 0{
				repeat--
				result += tempString
			}
		}
	}
	return result
}
func decodeString(s string) string {
    idx := 0
    return buildString(s, &idx)
}