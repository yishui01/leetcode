func trailingZeroes(n int) int {
    //原来阶乘末尾0的个数就是5的个数
    if n == 0 {
        return 0
    }
    return n/5 + trailingZeroes(n/5)
}

