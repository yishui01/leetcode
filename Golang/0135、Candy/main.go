func candy(ratings []int) int {
 //参考大佬解法，先全部初始化为1，从左往右一次，从右往左一次
    lens := len(ratings)
    if lens == 0 {
        return 0
    }
    
    candies := make([]int, lens)
    
    candies[0] = 1
    //从左往右，保证递增的是对的
    for i:=0; i < lens-1; i++ {
        if ratings[i+1] > ratings[i] {
            candies[i+1] = candies[i]+1
        } else {
            candies[i+1] = 1
        }
    }
    
    res := candies[lens-1]
    
    //从右往左，解决掉那些左边比右边大但是左边的糖果却没有比右边大的，那就左边糖果=右边糖果+1
    for i:=lens-2; i >=0; i-- {
        if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
            candies[i] = candies[i+1]+1
        }
        
        res += candies[i]
    }
    return res
}
