package leetcodego
import(
	"sort"
)
func deckRevealedIncreasing(deck []int) []int {
    //TC O(nlogn), SC O(n)
    sort.Ints(deck)
    deckSlice, ans := []int{}, []int{}

    //reverse the reveal and pop operation
    for i := len(deck) - 1; i >= 0; i--{
        deckSlice = append(deckSlice, deck[i])
        if i > 0{
            cur := deckSlice[0]
            deckSlice = deckSlice[1:]
            deckSlice = append(deckSlice, cur)
        }
    }

    for i := len(deck) - 1; i >= 0; i--{
        ans = append(ans, deckSlice[i])
    }

    return ans
}