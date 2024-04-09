package leetcodego

func timeRequiredToBuy(tickets []int, k int) int {
    //credit to lzl124631x, TC O(n), SC O(1)
    step := 0
    for idx, ticket := range tickets{
        if idx <= k{
            step += min(ticket, tickets[k])
        }else{
            step += min(ticket, tickets[k] - 1)
        }
    }
    return step
}