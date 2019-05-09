func divide(dividend int, divisor int) int {
    //思路是将除数左移，每左移一次，扩大2倍，也就是乘以2，知道除数扩大到大于被除数
    //将除数退回一格，保持除数小于被除数，将被除数-除数的差值作为被除数，再重复第一行的操作
    //左移的次数作为2的指数，结果就是除数扩大的倍数，也就是商，可以巧妙的用1<<n来得出结果
    //比如被除数是3，除数是3，除数左移1位，为6，大于被除数，倒回一格，此时左移次数为0，n为0，左移0位，结果为1，商为1
    
    if dividend == math.MinInt32 && divisor == -1 {
        //将唯一一种溢出情况过滤
        return math.MaxInt32
    }
    
    if dividend == 0 {
        return dividend
    }
    
    if divisor == math.MinInt32 &&  dividend != math.MinInt32{
        //除数为最小负数，可以直接返回0, 因为没人能除过他
        return 0
    }
    if divisor == math.MinInt32 &&  dividend == math.MinInt32{
        return 1
    }
    
    flag := 1
    if dividend > 0 && divisor < 0  || dividend < 0 && divisor > 0 {
        flag = -1
    }
    
    res := 0
    tmpRes := 0

    if dividend == math.MinInt32 {
        //被除数为最小负数，不好转正数，那就先除掉一个除数
        dividend = dividend + int(math.Abs(float64(divisor)))
        tmpRes = 1
    }
   
    dividend = int(math.Abs(float64(dividend)))
    divisor = int(math.Abs(float64(divisor)))
    
    var tmpDivisor int
    var i uint
    for dividend >= divisor {
        i = 0
        tmpDivisor = divisor
        //迭代
        for dividend >= tmpDivisor {
            tmpDivisor = divisor << i
            i++
        }
        //跳出循环时i-2为小于被除数的点
        res += 1 << (i-2)
        //迭代完成
        //重新赋值被除数
        dividend = dividend - (tmpDivisor >> 1)
    }
    
    return (res + tmpRes)*flag
    
    
}
