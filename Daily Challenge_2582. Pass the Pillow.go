package leetcodego

func passThePillow(n int, time int) int {
	tripCount := time / (n - 1)
	time %= (n - 1)

	if tripCount%2 == 0 {
		return 1 + time
	}
	return n - time
}