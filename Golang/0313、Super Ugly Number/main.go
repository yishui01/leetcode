func nthSuperUglyNumber(n int, primes []int) int {
    //每个primes元素一个队列，用一个数组保存每个队列的进度，找出每一轮的最小值，放到res中,将该队列的索引++
    if n == 0 {
        return 0
    }
    res := []int{1}
    indexQueue := make([]int,len(primes))
    minIndex := 0
    for len(res) < n {
        min := math.MaxInt32
        for i:=0;i<len(primes);i++{
            tmp := res[indexQueue[i]]*primes[i]
            if tmp <= min {
                min = tmp
                minIndex = i
            }
        }
       
        indexQueue[minIndex]++
         if min == res[len(res)-1] { //结果数组中可能有重复的所以要过滤，并且重复的感觉一定连在一起，因为是从小往大递增的，不可能出到某一个大数的时候出来一个重复的小数
            continue 
         } else {
             res = append(res,min)
         }
        
    }
    return res[len(res)-1]

}



