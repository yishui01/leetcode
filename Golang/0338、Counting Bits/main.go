func countBits(num int) []int {
    res := make([]int,num+1)
    if num == 0 {
        return res
    }

    //偶数就是偶数/2的那个位的1的个数
    //奇数就是奇数/2的那个位的1的个数+1
    
    for i:=1;i<=num;i++{
       if i%2==0{
            res[i] = res[i/2]
        } else {
            res[i] = res[i/2]+1
        }
    }

    return res

}

//还有，i&(i-1) 可以将i的二进制的最后一位1置零，由此可以用于判断是否为2的指数(为0就是2的指数)以及循环使用循环判断i是否为0来统计i的1的个数

//res[i] = res[i & (i - 1)] + 1; 还有这个等式也吊炸天

//还有一个区间DP【2,3】可以推出【4,5,6,7】再通过【4,5,6,7】推出【8,9,10,11,12,13,14,15】
