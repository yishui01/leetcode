func maxProfit(prices []int) int {
    //这题看了半天，无解，看答案，贪心，只要前一天的价格比今天的低，今天直接卖了，再从今天开始买入，继续往后找
    res := 0
    lens := len(prices)
    if lens < 2 {
        return res
    }
    for i:=1; i < lens; i++ {
        if prices[i] > prices[i-1] {
            res += prices[i]-prices[i-1]
        }
    }
    return res
}
