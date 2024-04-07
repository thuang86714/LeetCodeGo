package leetcodego
import(
	"unicode"
)
func minRemoveToMakeValid(s string) string {
    //TC O(n), SC O(n)
    pairCount, parenthesesSlice, result := 0, []rune{}, ""
    //first pass we count the number of valid parentheses
    for _, char := range s{
        if unicode.IsLetter(char){
            continue
        }
        if char == '('{
            parenthesesSlice = append(parenthesesSlice, char)
        }else{
            if len(parenthesesSlice) != 0 && parenthesesSlice[len(parenthesesSlice) - 1] == '('{
                parenthesesSlice = parenthesesSlice[:len(parenthesesSlice) - 1]
                pairCount++
            }
        }
    }

    leftCount, rightCount := pairCount, pairCount
    //second pass we build the string according to what we get in the first pass
    for _, char := range s{
        if unicode.IsLetter(char){
            result += string(char)
        }
        if char == '(' && leftCount > 0{
            result += string(char)
            leftCount--;
        }
        if char == ')' && rightCount > 0 && rightCount > leftCount{
            result += string(char)
            rightCount--
        }
    }
    return result

}


/*

func minRemoveToMakeValid(s string) string {
    var stack []int // Use a slice to simulate a stack for tracking '(' indices
    var removeIndex = make(map[int]bool) // Track indices of characters to remove

    // Single pass through the string
    for i, char := range s {
        if char == '(' {
            stack = append(stack, i) // Push the index onto the stack
        } else if char == ')' {
            if len(stack) == 0 {
                // Mark ')' for removal if there's no matching '('
                removeIndex[i] = true
            } else {
                // Pop the stack if there's a matching '('
                stack = stack[:len(stack)-1]
            }
        }
    }

    // Any indices left in the stack are '(' without a matching ')'
    for _, index := range stack {
        removeIndex[index] = true
    }

    // Build the result string by skipping characters marked for removal
    var result strings.Builder
    for i, char := range s {
        if !removeIndex[i] && (unicode.IsLetter(char) || char == '(' || char == ')') {
            result.WriteRune(char)
        }
    }

    return result.String()
}
*/