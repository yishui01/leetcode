func maxCoins(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}
	var left, right, max, cur, _left, _right int
	for i := 1; i <= length; i++ {
		for j := 0; j <= length-i; j++ {
			left = j
			right = j + i - 1
			// fmt.Println(left, right)
			max = 0
			cur = 0
			if left-1 < 0 {
				_left = 1
			} else {
				_left = nums[left-1]
			}
			if right+1 >= length {
				_right = 1
			} else {
				_right = nums[right+1]
			}
			//k是戳破气球的位置
			for k := left; k <= right; k++ {
				cur = _left * nums[k] * _right
				if k-1 >= left {
					cur += dp[left][k-1]
				}
				if k+1 <= right {
					cur += dp[k+1][right]
				}
				if max < cur {
					max = cur
				}
				//fmt.Println(left, right, cur, max)
			}
			dp[left][right] = max
		}
		// fmt.Println(dp)
	}
	return dp[0][length-1]
}
