func findMaximumXOR(nums []int) int {
    //又跪了...，玄学构造res，服气
    //利用异或的性质 a ^ b = c  =>   a ^ c = b , c ^ b = a
    //那这样的话我们可以一位一位的构造res
    //res每增加一位，就去和nums所有数字的对应位数的前缀全部异或一遍，看异或的结果是不是存在于这些前缀里面
    //如果res这一位为1成立的话，那res和这些前缀异或的结果肯定是这些前缀里面某一个，没有的话res这一位只有一种可能，为0
    //假设前缀数组里面有[a,b]  假设我构造出一个res，  如果res是这数组里面两个数异或出来的，我res ^ a的时候，他的结果肯定是b（肯定能在前缀数组中找到）
    
    mask := 0
    res := 0
    for i:=31;i>=0;i--{
        mask = mask | (1 << i)
        perfixMap := make(map[int]bool,len(nums)) //在知道map容量的情况下，提前给map分配好容量是一个不错的优化
        for _,v := range nums {
            perfixMap[v & mask] = true
        }
        
        tmp := res | (1 << i)  //假设res这一位为1
        
        for k,_ := range perfixMap{
            if perfixMap[k ^ tmp]{
                res = tmp //那么res就真的为1
                break
            }
        }
    }

    return res
}

