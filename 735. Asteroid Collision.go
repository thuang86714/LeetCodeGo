package leetcodego

import (
	"math"
)

func asteroidCollision(ast []int) []int {
	var asteroidStack []int
	for _, a := range ast {
		collision := false
		for len(asteroidStack) > 0 && a < 0 && asteroidStack[len(asteroidStack)-1] > 0 {
			if abs(a) > asteroidStack[len(asteroidStack)-1] {
				asteroidStack = asteroidStack[:len(asteroidStack)-1]
				continue
			} else if abs(a) == asteroidStack[len(asteroidStack)-1] {
				asteroidStack = asteroidStack[:len(asteroidStack)-1]
			}
			collision = true
			break
		}
		if !collision && (a > 0 || len(asteroidStack) == 0 || asteroidStack[len(asteroidStack)-1] < 0) {
			asteroidStack = append(asteroidStack, a)
		}
	}
	return asteroidStack
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
