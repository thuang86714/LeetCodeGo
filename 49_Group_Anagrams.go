/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

 

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
*/
package leetcodego
func groupAnagrams(strs []string) [][]string {
    mp := map[[26]int][]string{}
    for _, s :=range strs{
        k := [26]int{}
        for i:= 0; i < len(s); i++{
            k[s[i]- 'a']++
        }
        mp[k] = append(mp[k], s)
    }
    result := [][]string{}
    for _, v := range mp{
        result = append(result, v)
    }
    return result
}