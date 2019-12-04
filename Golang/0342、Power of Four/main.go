func isPowerOfFour(num int) bool {
    if num < 1 {
        return false
    }
    if num & (num-1) == 0 {
       //现在已经是2的次幂了，就是说整个二进制中只有一个1，接下来有两种判断方法，
       //1、与0x55555555相与（1010101010101010101010101010101）如果得到的数还是其本身，则可以肯定其为4的次方数：
            //解释（因为4的次幂后面的0的个数都是偶数个，1-0,100-4,10000-16,1000000-64），所以和那个0x5555555相与就是确定1的位置，由此判断出1后面0的个数
    
        //2、(num-1)%3 == 0 看是否能被3整除，能的就符合条件，我真跪了
        
        return (num-1)%3 == 0
    }

    return false
}

/*
func isPowerOfFour(num int) bool {
    if num < 1 {
        return false
    }
    if num & (num-1) == 0 {
        //只有一个1，已经是2的次幂了
        //接下来判断这个1后面的0是不是偶数个即可
        res := 0
        for {
            tmp := num & 1
            num = num >> 1
            if tmp == 0 {
                res++
            } else {
                break
            }
        }
        return res % 2 == 0
    }

    return false
}*/
