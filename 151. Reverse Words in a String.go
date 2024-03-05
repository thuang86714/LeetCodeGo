func reverseWords(s string) string {
    //credit to vibhushit
    //Returns: A slice of substrings of str or an empty slice if str contains only white space.
    words := strings.Fields(s) 

    left, right := 0, len(words) - 1

    for left < right{
        words[left], words[right] = words[right], words[left]
        left++
        right--
    }

    // Join the reversed words back into a single string and trim space
    //TrimSpace remove all leading and trailing zeros
    return strings.TrimSpace(strings.Join(words, " "))
}