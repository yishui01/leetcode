func rotate(nums []int, k int)  {
    //自己想到的只有暴力法
    //看了大佬答案有4中解法：暴力、额外数组一次性替换、环形替换、反转替换，这里使用反转替换
    lens:=len(nums)
    
    //1、先将数组全部反转
    //2、将前k个数字反转
    //3、将后面n-k个数字也反转
    
    if lens == 1 {
        return 
    }
    
   k %= lens //我真是服了，原来k还可以比lens大，要取下模才能进行运算，否则下面的算法会越界
    
   for i := 0; i < lens/2; i++ {
		nums[i], nums[lens-1-i] = nums[lens-1-i], nums[i]
	}

	for i := 0; i < k/2; i++ {
		nums[i], nums[k-i-1] = nums[k-i-1], nums[i]
	}

    //还有这里反转后面n-k个元素的时候要注意各种下标，否则很容易越界
	for i, j := k, 0; i < (lens+k)/2; i, j = i+1, j+1 {
		nums[i], nums[lens-1-j] = nums[lens-1-j], nums[i]
	}
}
