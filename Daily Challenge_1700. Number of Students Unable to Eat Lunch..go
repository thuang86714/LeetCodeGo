package leetcodego

func countStudents(students []int, sandwiches []int) int {
    //credit to lee215, TC O(n), SC O(1)
    size, idx := len(students), 0
    preferenceCount := make([]int, 2)
    for _, preference := range students{
        preferenceCount[preference]++
    }
    for ; idx < size && preferenceCount[sandwiches[idx]] > 0; idx++{
        preferenceCount[sandwiches[idx]]--
    }
    return size - idx;
}