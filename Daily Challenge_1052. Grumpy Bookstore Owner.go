package leetcodego

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	totalSatisfied := 0
	n := len(customers)

	// Calculate the initial satisfaction without using any grumpy technique
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			totalSatisfied += customers[i]
		}
	}

	// Calculate the maximum increase in satisfied customers by using the grumpy technique
	maxIncrease := 0
	currentIncrease := 0

	// Calculate the initial window
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			currentIncrease += customers[i]
		}
	}
	maxIncrease = currentIncrease

	// Slide the window across the array
	for i := minutes; i < n; i++ {
		if grumpy[i] == 1 {
			currentIncrease += customers[i]
		}
		if grumpy[i - minutes] == 1 {
			currentIncrease -= customers[i - minutes]
		}
		maxIncrease = max(maxIncrease, currentIncrease)
	}

	return totalSatisfied + maxIncrease
}