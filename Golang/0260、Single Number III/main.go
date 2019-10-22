func singleNumber(nums []int) []int {
    //这题我看完解答之后我给跪了
    //出现两次的异或到最后会被抵消，问题是这里面有两个出现一次的数字，
    //怎么把这两个数字区分开？
    //所有数字异或之后的结果实际上就是这两个数字异或的结果（因为其他数字都是出现两次，被抵消成0了）
    //拿着这个异或的结果，我们要把他做成一个mask，用这个mask来区分出这两个变量
    //找出这个异或结果其中为1的一位，这一位为1代表这两个数字在这一位是不同的
    //那么这是一个mask，将mask的这一位置为1，通过 res & (-res)的方式取出异或结果最右边为1的一位，
    //因为负数的表示是补码=反码+1，所以能得出上面的结果。
    //拿到这个mask就一个个去和nums中的元素异或就行了，异或成0的分到一组，累计异或，异或成1的放到一组累计异或
    //最后这两组的结果就是那两个数字
    
    lens := len(nums)
    if lens <= 2 {
        return nums
    }
    
    diff := 0
    for _,v := range nums {
        diff ^= v
    }
    
    mask := diff &(-diff)
    
    res := make([]int,2)
    
    for _,v := range nums {
        if v & mask == 0 {
            res[0] ^= v
        } else {
            res[1] ^= v
        }
    }
    
    return res
}
