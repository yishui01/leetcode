func lastRemaining(n int) int {

    //参考题解：https://leetcode-cn.com/problems/elimination-game/solution/gui-lu-de-xun-zhao-by-luo-ben-zhu-xiao-man-tou/
    
    //将原问题转化为求找数列第一个数的变化规律
    //最后得出
    if n == 0 {
        return 0
    }

    first := 1 //第一个数
    factor := 1 //公差
    isLeft := true //方向
    isEven := (n&1) == 0 //长度是否为偶数

    for n != 1 {
        if isLeft || !isEven {  //若从左向右或者从右向左时为奇数，第一个数变为原来的第二个
            first += factor 
        }
        isLeft = !isLeft                                //变向
        n = n >> 1                                      //长度/2
        factor = factor << 1                            //公差翻倍
        isEven = ((n&1) == 0)                           //判断奇偶
    }

    return first
}
