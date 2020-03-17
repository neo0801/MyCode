package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	top := make([]int, n)
	piles := 0
	for i := 0; i < n; i++ {
		poker := nums[i]

		left, right := 0, piles
		for left < right {
			mid := (left + right) / 2
			if top[mid] > poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == piles {
			piles++
		}
		top[left] = poker
	}
	return piles
}

func lengthOfLIS2(nums []int) int {
	n := len(nums)
	dp := make([]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = 1.0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = math.Max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0.0
	for i := 0; i < n; i++ {
		res = math.Max(res, dp[i])
	}
	return int(res)
}
