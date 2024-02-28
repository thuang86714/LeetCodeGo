package leetcodego
import "testing"

func groupAnagrams(strs []string) [][]string {
    //using hashmap, with key is int slice, and slice of string as value
    //TC O(m*n), where m is len(strs) and n is avg len of all words in strs
    //SC O(m)
    charStrMap := map[[26]int][]string{}
    for _, s:= range strs{
        k := [26]int{}
        for _, ch := range s{
            k[ch - 'a']++
        }
        charStrMap[k] = append(charStrMap[k], s)
    }

    result := [][]string{}
    for _, v := range charStrMap{
        result = append(result, v)
    }
    return result
}