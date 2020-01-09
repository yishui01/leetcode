func maxRotateFunction(A []int) int {
    //跪了跪了，直接看想不出规律，看了答案才知道规律，好吧，智商是硬伤哈哈哈哈
    /*
    F(0) = 0A + 1B + 2C +3D

    F(1) = 0D + 1A + 2B +3C

    F(2) = 0C + 1D + 2A +3B

    F(3) = 0B + 1C + 2D +3A

    那么，我们通过仔细观察，我们可以得出下面的规律：

    sum = 1A + 1B + 1C + 1D

    F(1) = F(0) + sum - 4D

    F(2) = F(1) + sum - 4C

    F(3) = F(2) + sum - 4B

    那么我们就找到规律了, F(i) = F(i-1) + sum - n*A[n-i]，可以写出代码如下：
    */
    lens := len(A)
    sum := 0
    t := 0
    for i:=0;i<lens;i++{
        sum += A[i]
        t += i*A[i]
    }
    res := t
    for i:=lens-1;i >= 0;i--{
     t = t + sum - lens * A[i]
     res = max(res,t)
    }

    return res

}

func max(a, b int)int {
    if a > b {
        return a
    }
    return b
}
