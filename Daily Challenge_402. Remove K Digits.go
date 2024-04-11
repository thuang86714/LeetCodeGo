package leetcodego
import(
	"strings"
)
func removeKdigits(num string, k int) string {
    if len(num) <= k {
        return "0"
    }

    if k == 0 {
        return num
    }

    var numSlice []byte

    for i := 0; i < len(num); i++ {
        for k > 0 && len(numSlice) > 0 && num[i]-'0' < numSlice[len(numSlice)-1]-'0' {
            numSlice = numSlice[:len(numSlice)-1]
            k--
        }
        numSlice = append(numSlice, num[i])
        // Prevent appending leading zeros to the result
        if len(numSlice) == 1 && num[i] == '0' {
            numSlice = numSlice[:0]
        }
    }

    // If k is still positive, remove the last k digits
    for k > 0 && len(numSlice) > 0 {
        numSlice = numSlice[:len(numSlice)-1]
        k--
    }

    // Convert numSlice to string
    result := string(numSlice)

    // If the result is empty, return "0", otherwise return result
    if result == "" {
        return "0"
    }

    // Remove leading zeros from the final result
    result = strings.TrimLeft(result, "0")

    // If trimming leaves an empty string, it means the number was all zeros
    if result == "" {
        return "0"
    }

    return result
}