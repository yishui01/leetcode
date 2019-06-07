func myPow(x float64, n int) float64 {
    //好吧开始我以为这题是移位运算，但是我还是太单纯，呵呵，float怎么移位你告诉我？？？
    //把n对折递归乘下去
    if n == 0 {
        return 1
    }
    if n == 1 {
        return x
    }
    if n == -1 {
        return 1/x
    }
    
    half := myPow(x, n/2) //n次方对折后的结果，结果看n是单数还是双数
    if n%2 == 0 {
        //双数就直接返回
        return half * half
    } else {
        if n < 0 {
            return half * half / x
        }
        return half * half * x
    }
}
