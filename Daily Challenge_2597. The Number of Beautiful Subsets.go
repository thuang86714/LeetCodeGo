package leetcodego
import(
	"math"
	"sort"
)
func beautifulSubsets(A []int, k int) int {
	// Initialize count slice with k empty maps
	count := make([]map[int]int, k)
	for i := range count {
		count[i] = make(map[int]int)
	}
	
	// Count occurrences of each element in A, grouped by modulo k
	for _, a := range A {
		count[a % k][a]++
	}

	res := 1
	
	// Process each group of elements with the same modulo k value
	for i := 0; i < k; i++ {
		prev, dp0, dp1 := 0, 1, 0
		keys := make([]int, 0, len(count[i]))
		
		// Extract the keys (unique elements) in the current group
		for key := range count[i] {
			keys = append(keys, key)
		}
		
		// Sort the keys
		sort.Ints(keys)

		// Iterate through the sorted keys
		for _, a := range keys {
			v := int(math.Pow(2, float64(count[i][a])))  // Calculate power of 2 for the count of the current element
			
			if prev + k == a {
				dp0, dp1 = dp0 + dp1, dp0 * (v - 1)
			} else {
				dp0, dp1 = dp0 + dp1, (dp0 + dp1) * (v - 1)
			}
			prev = a
		}
		res *= dp0 + dp1
	}
	return res - 1
}