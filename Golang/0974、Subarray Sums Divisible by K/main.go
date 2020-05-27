func subarraysDivByK(A []int, K int) int {
    //这种子数组，一般要考虑滑动窗口或者是前缀和
    //再利用同余定理 a%k == b %k  => (a-b)%k==0
    //题目要求(a-b)%k == 0 (其中a和b都是前缀和)
    //那就转换为 a%k == b%k
    //用hash记录前缀和%k的值，值为key，次数为value
    // (a + b) mod c = (a mod c + b mod c) mod c
    //还要考虑模完之后为负数的情况，比如一个模完为3，另一个模完为-1，
    maps := make(map[int]int,0)
    maps[0] = 1
    res := 0
    sum := 0
    for _,v:= range A {
        sum += v
        t:= sum % K
        if t < 0 {
            t += K
        }
        res+=maps[t]
        maps[t]++
    }
    return res
}

