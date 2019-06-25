func simplifyPath(path string) string {
    
    //这题花了我两个小时，还是各种调试才做出来的，谁能告诉我我TM为什么这么蠢，为什么？恩？
    
    lens := len(path)
    if lens == 0 {
        return ""
    }
    res := make([]byte, 0)
    tmpSlice := make([]byte, 0)
    
    isPeriod := 0
    isSlash := 0
    isLetter :=0
    
    slashs := make([]int, 0)
    nowIndex:= -1
    for _,v := range path {
        switch v {
            case '.':
                isPeriod++
                isSlash = 0
                tmpSlice = append(tmpSlice, '.')
            case '/':
                if isSlash > 0 { //之前是 '/' 跳过
                    continue
                }
                if isPeriod == 0 || isLetter != 0 || isPeriod > 2 { //有三种情况进行追加，1、前面没有点  2、前面有字母  3、点的数量超过两个 ，这三种都代表前面的是一个文件夹名称，不需要转换目录
                             tmpSlice = append(tmpSlice, '/')
                             dist := make([]byte, len(tmpSlice))
                             copy(dist,tmpSlice)
                             res = append(res, dist...)
                             nowIndex += len(tmpSlice)
                             slashs = append(slashs, nowIndex)
                } else if isPeriod == 2 { //到这里只考虑两个点的情况，一个点加斜杠直接忽略，两个点，又没有字母帮忙，那就是返回上一级了
                            //上级目录
                            if slashs[len(slashs)-1] > 0 {
                                nowIndex = slashs[len(slashs)-2]
                                res = res[:slashs[len(slashs)-2]+1]
                                slashs = slashs[:len(slashs)-1]
                            }
                    }
                tmpSlice = []byte{}
                isPeriod = 0
                isLetter = 0
                isSlash = 1
            default:
                //字母
                isLetter = 1
                isSlash = 0
                tmpSlice = append(tmpSlice, byte(v))
        }
         
    }
    
    //处理末尾是1个.或者两个.的情况
    if len(tmpSlice) > 0 {
        if len(tmpSlice) == 2 && tmpSlice[0] == '.' && tmpSlice[1] == '.'  {
                //返回到上一级
                 if slashs[len(slashs)-1] > 0 {
                            nowIndex = slashs[len(slashs)-2]
                            res = res[:slashs[len(slashs)-2]+1]
                            slashs = slashs[:len(slashs)-1]
                }
        } else if len(tmpSlice) == 1 && tmpSlice[0] == '.' {
            //当前目录，不用管
        } else {
            res = append(res,tmpSlice...)
        }
        
    }
    
    //去掉末尾的 '/'
    for len(res) > 1 && res[len(res)-1] == '/' {
        res = res[:len(res)-1]
    }
    return string(res)
}
