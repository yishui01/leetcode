func restoreIpAddresses(s string) []string {
    //这种需要排列值的，只能用回溯
    //4段，每段取1-3个数字，每次取完判断是否合法，不合法就直接丢弃，取到第四段的时候直接判断剩下的数字
    
    if s == "" {
        return []string{}
    }
    res:=make([]string, 0)
    recursion(s, 0, "", &res)
    return res
    
}

func recursion(s string, level int , tmpRes string, res *[]string) {
    if level == 4 {
        *res = append(*res, tmpRes[1:])
        return 
    }
    
    if len(s) == 0 {
        return 
    }
    
    if level == 3 {
        
        //这是第4段了，那就是直接看剩下的数字，不要截取了,注意剩下的数字不能是0开头的，除非是只是一个0，那就OK，像010这样的就直接G
        if val, ok := strconv.Atoi(s); ok == nil && val >= 0 && val <= 255 && (s[0] != '0' || len(s) == 1) {
            //合格就进入下一层
            recursion(s, level+1, tmpRes + "."+s, res)
        }
        return 
    }
    
     //分别尝试截取1位、2位、3位
     //首先截取一位是无论如何都可以的,0开头的也行，就是一个0嘛
    recursion(s[1:], level+1, tmpRes + "."+string(s[0]), res)
     //再来两位的
     if s[0] != '0' && len(s) >= 2{
            recursion(s[2:], level+1, tmpRes + "."+s[:2], res)
     }
    //再来三位的
     if s[0] != '0' && len(s) >= 3{
         if val, ok := strconv.Atoi(s[:3]); ok == nil && val >= 0 && val <= 255 {
             recursion(s[3:], level+1, tmpRes + "."+s[:3], res)
         }
    }
    
    
}


