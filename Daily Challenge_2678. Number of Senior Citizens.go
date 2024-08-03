package leetcodego
import(
	"strconv"
)

// countSeniors counts the number of people older than 60 based on the details provided
func countSeniors(details []string) int {
	count := 0
	for _, detail := range details {
		numByte := []byte{}
		numByte = append(numByte, detail[11], detail[12])
		tempStr := string(numByte)
		num, _ := strconv.Atoi(tempStr)
		if num > 60 {
			count++
		}
	}
	return count
}
