package leetcodego

func checkValidString(s string) bool {
    //credit to hiepit, TC O(n), SC O(1)
    minOpenPairCount, maxOpenPairCount := 0, 0
    for _, char := range s{
        if char == '('{
            minOpenPairCount++
            maxOpenPairCount++
        }else if char == ')'{
            minOpenPairCount--
            maxOpenPairCount--
        }else if char == '*'{
            minOpenPairCount-- // if `*` become `)` then openCount--
            maxOpenPairCount++// if `*` become `(` then openCount++
            // if `*` become `` then nothing happens
            // So openCount will be in new range [cmin-1, cmax+1]
        }
        if maxOpenPairCount < 0{
            return false
        }

        minOpenPairCount = max(0, minOpenPairCount)
    }
    return minOpenPairCount == 0
}