func findNthDigit(n int) int {
    if n == 0 {
        return 0
    }
    level := 1 //第几个区段 每一个区段的数字位数是相同的
    count := 0 //总数目
    for count < n {
        tmp := getCount(level)
        if count + tmp <= n {
            count += tmp
            level++
        } else {
            break
        }
        
    }
    //跳出for循环之后的res就在level这个区段内的
    //那具体是这个区段内第几个数字？
    //这个区段的起始数、当前count已经消耗了多少数字，还剩多少数字，剩下的这些数字能在这个区段内排到第多少个数，余多少位
    start := int(math.Pow(10,float64(level-1)))  //本区段第一个数
    diff := n - count //还有diff个数字可以消耗，本区段每个数长度为level
    di := diff / level
    mod := diff % level
    
    res := start+di
    if mod == 0 {
        //res - 1 的最后一位数字
        return (res-1) % 10
    }

   //res的从左往右第mod位数字，那也就是从右往左的第level - mod+1位数字
   p := level - mod + 1
   i := 0
   r := 0
   for i < p {
       r = res % 10
       res /= 10
       i++
   }

    return r
}

//返回当前区段的数字个数 * 每个数字的长度（不包括当前区段之前的区段）
func getCount(level int) int {
    return level * (int(math.Pow(10,float64(level))) - int(math.Pow(10,float64(level-1))))
}


/*
//看下大佬的几行就搞定，我的写那么长。。。
func findNthDigit(n int) int {
    len, cnt, start := 1,9,1

    for n > len * cnt {
        n-=len*cnt
        len++
        cnt *= 10
        start *= 10
    }
    start += (n-1) / len
    d := (n-1) % len
    //从左往右第d位就是数字
    s := strconv.Itoa(start)
    return int(s[d] - '0')
}
*/
