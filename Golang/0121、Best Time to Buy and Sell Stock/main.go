func maxProfit(prices []int) int {
    //O(n)一遍过优化版本
    lens := len(prices)
    res := 0
    if lens == 0 {
        return res
    }
    min := prices[0]
    for i:=0;i <lens; i++ {
        if prices[i] < min {
            min = prices[i]
        }
        if res < prices[i]-min{
            res = prices[i]-min
        }
    }
    return res
}

/**
   //每一个点看做是出售日期，那么到达这个出售日期只能是出售日期的前一个日期。那就好办了，找到出售日期前的最小的那个数字
    lens := len(prices)
    res := 0
    for i:=1;i < lens; i++{
        min := prices[i]
        for j:=i-1; j >=0;j--{
            if prices[j] < min {
                min = prices[j]
            }
        }
        if prices[i] - min > res{
            res = prices[i] - min
        }
    }
    
    return res
**/
