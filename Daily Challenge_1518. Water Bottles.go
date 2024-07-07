package leetcodego

func numWaterBottles(numBottles int, numExchange int) int {
	ans := 0
	residualBottle := 0
	for numBottles > 0 {
		ans += numBottles
		curBottle := (numBottles + residualBottle)
		residualBottle = curBottle % numExchange
		numBottles = curBottle / numExchange
	}
	return ans
}