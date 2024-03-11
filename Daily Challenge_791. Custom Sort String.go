package leetcodego

import (
	"strings"
)

func customSortString(order string, s string) string {
    //TC O(n + m), where n and m are the length of order and s respectively, SC O(1)
    charCount := make([]int, 26)//count the appearance of char

    for _, ch := range s{//then we know the overall appearance of s
        charCount[ch - 'a']++;
    }

    var ans strings.Builder//an empty string as answer

    // Construct the answer based on the order.
    for _, ch := range order{
        if charCount[ch - 'a'] > 0{
            ans.WriteString(strings.Repeat(string(ch), charCount[ch - 'a']))
            charCount[ch - 'a'] = 0
        }
    }

    // Add remaining characters not in order.
    for i := 0; i < 26; i++{
        if(charCount[i] > 0){
            ans.WriteString(strings.Repeat(string('a' + i), charCount[i]))
        }
    }

    return ans.String()
}