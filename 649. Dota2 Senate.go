package leetcodego

func predictPartyVictory(senate string) string {
    //credit to caihao0727mail, TC O(n), SC O(n)
    rIdxSlice, dIdxSlice, n := []int{}, []int{}, len(senate)
    for idx, ch := range senate{
        if ch == 'R'{
            rIdxSlice = append(rIdxSlice, idx)
        }else{
            dIdxSlice = append(dIdxSlice, idx)
        }
    }

    for len(rIdxSlice) != 0 && len(dIdxSlice) != 0{
        rIdx, dIdx := rIdxSlice[0], dIdxSlice[0]
        rIdxSlice = rIdxSlice[1:]
        dIdxSlice = dIdxSlice[1:]
        if rIdx < dIdx{
            rIdxSlice = append(rIdxSlice, rIdx + n)//the operation plus n means the corresponding senate can keep "alive" to the next round.
        }else{
            dIdxSlice = append(dIdxSlice, dIdx + n)
        }
    }

    if len(rIdxSlice) > len(dIdxSlice){
        return "Radiant"
    }
    return "Dire"
}