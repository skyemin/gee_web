package main

func trap(height []int) int {
	/*	sum := 0
		for i := 1; i < len(height)-1; i++ {
			//找出左边最高
			maxLeft := 0
			for j := i - 1; j >= 0; j-- {
				if height[j] > maxLeft {
					maxLeft = height[j]
				}
			}
			maxRight := 0
			for j := i + 1; j < len(height); j++ {
				if height[j] > maxRight {
					maxRight = height[j]
				}
			}
			min := Min(maxLeft, maxRight)
			if min > height[i] {
				sum += min - height[i]
			}
		}
		return sum*/

	sum := 0
	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))

	for i := 1; i < len(height)-1; i++ {
		maxLeft[i] = Max(maxLeft[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = Max(maxRight[i+1], height[i+1])
	}
	for i := 1; i < len(height)-1; i++ {
		min := Min(maxLeft[i], maxRight[i])
		if min > height[i] {
			sum += min - height[i]
		}
	}
	return sum
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
