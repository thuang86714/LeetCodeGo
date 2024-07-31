package leetcodego

func minHeightShelves(books [][]int, shelfWidth int) int {
	heightMemo := make(map[int]int)

	var findMinHeight func(int) int
	findMinHeight = func(idx int) int {
		if idx >= len(books) {
			return 0
		}
		if _, exist := heightMemo[idx]; exist {
			return heightMemo[idx]
		}

		height, curIdx, curWidth, curLevelHeight := 1000001, idx, shelfWidth, books[idx][1]
		for curIdx < len(books) && curWidth-books[curIdx][0] >= 0 {
			curWidth -= books[curIdx][0]
			curLevelHeight = max(curLevelHeight, books[curIdx][1])
			curIdx++
			height = min(height, curLevelHeight+findMinHeight(curIdx))
		}

		heightMemo[idx] = height
		return height
	}
	return findMinHeight(0)
}
