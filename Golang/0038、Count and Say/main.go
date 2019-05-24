func countAndSay(n int) string {
    //这个和裴波那契的递归一毛一样╮(╯▽╰)╭
    if n == 1 {
        return "1"
    } else {
        return translate(countAndSay(n-1))
    }
}

func translate(s string) string{
    tmp := ' '
    nums:= 0
    res := make([]byte, 0);
    if len(s) == 1 {
        return "11" //长度为1只有一种能可能，就是 "1"
    }
    for _,v := range s {
        if tmp != ' ' && v != tmp {
            //新一轮替换
            res = append(res, byte(nums+'0'))
            res = append(res, byte(tmp))
            nums = 1
            tmp = v
        } else {
            tmp = v
            nums++
        }
    }
    if nums != 0 {
        res = append(res, byte(nums+'0'))
        res = append(res,  s[len(s)-1])
    }
 
    return string(res)
}


