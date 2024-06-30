package leetcodego

func isSubsequence2(s string, t string) bool {
	idxT, idxS := 0, 0
	for idxS < len(s) && idxT < len(t){
        if s[idxS] == t[idxT]{
            idxS++
        }
        idxT++
	}
	return idxS == len(s)
}