func canConstruct(ransomNote string, magazine string) bool {
    //credit to josh-ferrell, TC O(n^2), SC O(1)
    for _,v := range ransomNote{
        if strings.Count(ransomNote, string(v)) > strings.Count(magazine, string(v)){
            return false;
        }
    }
    return true
}

func canConstruct(ransomNote string, magazine string) bool {
    //TC O(n), SC O(n)
    charMap := make(map[rune]int)//using map to track count
    for _,ch := range magazine{
        charMap[ch]++
    }
    for _,ch := range ransomNote{
        charMap[ch]--
    }
    for _, count := range charMap{
        if count < 0{
            return false
        }
    }
    return true
}