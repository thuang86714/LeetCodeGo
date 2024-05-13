package leetcodego

import (
	"strconv"
	"sort"
)

func findRelativeRanks(score []int) []string {
    //TC O(nlogn), SC O(n)
    ans := make([]string, len(score))
	//a custom struct to store index in the score slice and score
    indexedScoreSlice := make([]struct{
        curIndex int
        curScore int
    }, len(score))

	//iterate to fill up indexedScoreSlice
    for i := 0; i < len(score); i++{
        indexedScoreSlice[i].curScore = score[i]
        indexedScoreSlice[i].curIndex = i
    }

	//custom sort by curScore
    sort.Slice(indexedScoreSlice, func(i, j int)bool{
        return indexedScoreSlice[i].curScore > indexedScoreSlice[j].curScore
    })

	//using switch to make it more consice
    for rank, pair := range indexedScoreSlice{
        switch rank{
            case 0:
                ans[pair.curIndex] = "Gold Medal"
            case 1:
			    ans[pair.curIndex] = "Silver Medal"
		        case 2:
			    ans[pair.curIndex] = "Bronze Medal"
		    default:
			    ans[pair.curIndex] = strconv.Itoa(rank + 1)
        }
    }
    return ans
}