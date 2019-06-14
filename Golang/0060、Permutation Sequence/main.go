func getPermutation(n int, k int) string {
    //列出全部的组合G了，超时
    if n == 0 {
        return ""
    }
    if n == 1 {
        return strconv.Itoa(n)
    }
    res := make([]byte,n)
    nums := make([]byte,n)
    for i := 0; i <n; i++ {
        nums[i] = byte(i+1) + '0'
    }
    fact := getFact(n-1) //最高位每个数字出现的次数
   
    k-- //这里k从1开始的，不好通过除法直接得出索引，减个1就可以了
    
    for i:=0; i < n; i++ {
        if len(nums)== 0 || fact == 0{
            break
        }
        index:= k/fact //当前在当前区域的第几号位
        k %= fact //下一段的总数
        res[i] = nums[index]
        nums = append(nums[:index], nums[index+1:]...)
        if n-i-1 > 0 {
            fact /= n-i-1
        }
        
    }
    
    return string(res)
}

//获取阶乘值
func getFact(n int) int {
    if n == 0 {
        return 0
    }
    res := 1
    for i:=0; i < n;i++ {
        res *= (n-i)
    }
    
    return res
}
