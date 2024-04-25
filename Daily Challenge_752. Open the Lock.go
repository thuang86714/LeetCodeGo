package leetcodego

func increaseChar(cur rune)rune{
    if cur == '9'{
        return '0'
    }
    return cur + 1
}
func decreaseChar(cur rune)rune{
    if cur == '0'{
        return '9'
    }
    return cur - 1
}
func openLock(deadends []string, target string) int {
    deadendsMap := make(map[string]bool)
    for _, dead := range deadends{
        deadendsMap[dead] = true
    }
    if _, ok := deadendsMap["0000"]; ok{
        return -1
    }
    visitedMap := make(map[string]bool)
    visitedMap["0000"] = true

    toVisitSlice := []string{}
    toVisitSlice = append(toVisitSlice, "0000")
    stepCount := 0

    for len(toVisitSlice) != 0{
        curSize := len(toVisitSlice)
        for curSize > 0{
            curString := toVisitSlice[0]
            toVisitSlice = toVisitSlice[1:]
            if curString == target{
                return stepCount
            }
            
            for idx := range curString{
                chars := []rune(curString)
                chars[idx]= increaseChar(chars[idx])
                newString := string(chars)
                if _, found := deadendsMap[newString]; found{
                    continue
                }
                if _, found := visitedMap[newString]; !found{
                    toVisitSlice = append(toVisitSlice, newString)
                    visitedMap[newString] = true
                }

    	        chars[idx] = decreaseChar(decreaseChar(chars[idx]))
                newString = string(chars)
                if _, found := deadendsMap[newString]; found{
                    continue
                }
                if _, found := visitedMap[newString]; !found{
                    toVisitSlice = append(toVisitSlice, newString)
                    visitedMap[newString] = true
                }
                chars[idx] = increaseChar(chars[idx])
            }
            curSize--
        }
        stepCount++
    }
    return -1
}