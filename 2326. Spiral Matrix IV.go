/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcodego
 func spiralMatrix(m int, n int, head *ListNode) [][]int {
    matrix := make([][]int, m)
    for i := 0; i < m; i++{
        matrix[i] = make([]int, n)
    }
    top, bottom, left, right := 0, m - 1, 0, n - 1
    for top <= bottom && left <= right{
        //traverse the top row
        for j := left; j <= right && top <= bottom; j++{
            if head != nil{
                matrix[top][j] = head.Val;
                head = head.Next;
            }else{
                matrix[top][j] = -1;
            } 
        }
        top++;

        //traverse the last col
        for i := top; i <= bottom && left <= right; i++{
            if head != nil{
                matrix[i][right] = head.Val;
                head = head.Next;
            }else{
                matrix[i][right] = -1;
            }
        }
        right--;

        //traverse the bottom row
        for j := right; j >= left && top <= bottom; j--{
            if head != nil{
                matrix[bottom][j] = head.Val;
                head = head.Next;
            }else{
                matrix[bottom][j] = -1;
            }
        }
        bottom--;

        for i := bottom; i >= top && left <= right; i--{
            if head != nil{
                matrix[i][left] = head.Val;
                head = head.Next;
            }else{
                matrix[i][left] = -1;
            }
        }
        left++;
    }
    return matrix
}