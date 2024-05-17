package leetcodego

func generateParenthesis(n int) []string {
    result := make([]string, 0)
    current := make([]byte, n*2)
    backtrack2(&result, n, 0, 0, current)
    return result
}

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

 

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]
*/
func backtrack2(res *[]string, n int, left int, right int, cur []byte){ 
    if left+right == n*2{
        *res = append(*res, string(cur))
    }
    if left < n {
        cur[left+right] = '('
        backtrack2(res, n, left+1, right, cur)
    }

    if right < left {
        cur[left+right] = ')'
        backtrack2(res, n, left, right+1, cur)
    }
}