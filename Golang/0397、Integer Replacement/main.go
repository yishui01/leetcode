func integerReplacement(n int) int {
    //这题有坑，以为-1永远比+1好，实际上对于15这种数来说，+1的效果更好


    //方法二：判断当前这个数是该加1还是减1
    res := 0
    for n != 1 {
        if n % 2 == 0 {
            n /= 2
        } else { 
            //如果n是奇数（代表二进制最后一位为1）

            //并且n & 2为2，代表二进制到底第二位为1，那么此时+2=1就会出现末尾两个0的情况，末尾两个0实际上就是可以除两次的那种，也就是4的倍数
            if n & 2 == 2 && n != 3 {
                n++
            } else {
                n--
            }
        }
        res++
    }

    return res

    //方法一：dfs 执行用时 4ms
    /*if n == 1 {
        return 0
    }

    if n % 2 == 0 {
        return 1+integerReplacement(n/2)
    }

    return min(1+integerReplacement(n-1),1+integerReplacement(n+1))*/
}

func min(a,b int)int{
    if a < b {
        return a
    }
    return b
}
