package leetcodego
import(
	"sort"
)
func successfulPairs(spells []int, potions []int, success int64) []int {
    //TC O(max(nlogn, m*logn)), where m is spells.size(), n is potions.size(), SC O(1)
    sort.Ints(potions)
    pairs := []int{}
    for _,spell:= range spells{
        left, right :=0, len(potions) - 1
        for left <= right{
            mid := left + (right - left)/2
			// Casting spell to int64 to ensure multiplication doesn't overflow.
            if int64(spell*potions[mid]) >= success{//this potion would succed, go left to find the potion with less strength
                right = mid - 1
            }else{
                left = mid + 1
            }
        }
        pairs = append(pairs, len(potions) - left);//all potions on the right of the left potion(inclusive) would be successful 
    }
    return pairs;
}