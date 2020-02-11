func toHex(num int) string {
    if num == 0 {
        return "0"
    }
    //4位4位的取
    str := "0123456789abcdef"
    mask := 15  //末尾4个1  或者0xf
    res := ""
    for i:=0;i<8;i++{
        //4位一截
        tmp := num & mask
        num = num >> 4
        res = string(str[tmp])+res
    }

    //把res前导0去掉
    z := 0
    for k,v := range res {
        if v == '0' {
            z = k+1
        } else {
            break
        }
    }

    return res[z:]
}
