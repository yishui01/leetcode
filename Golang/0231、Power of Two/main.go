func isPowerOfTwo(n int) bool {
    //2的次幂的二进制只有一个1
    count := 0
    for n > 0 {
        if n & 1 == 1{
            count++
        }
        if count > 1 {
            return false
        }
         n = n >> 1
    }
    
    return count == 1
}
