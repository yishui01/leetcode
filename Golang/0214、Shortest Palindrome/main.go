func shortestPalindrome(s string) string {
    //在头部添加字符使得整体变为回文,那么先找出头部开始能成的最大回文，可以想象成原本是一个完整的回文串，但是前半段里面有以小部分被人砍了，现在要找回来
    //查找从0到i可以成回文的最大长度
    //1、先将整体反转得到recv
    //2、找到 s[0:n-i] == recv[i:] 的i点，此时 s[0:n-i]或者recv[i:]就是从头部开始的最大回文串
    //3、接下来将recv[0:i]加到最大回文串前面，构成完整回文
    lens := len(s)
    recv := reverse(s)
    i:=0
    for i < lens && s[:lens-i] != recv[i:] {
        i++
    }
 
    return recv[:i]+s
    
}

func reverse(s string) string {
    sli := []byte(s)
    lens := len(s)
    for i:=0;i<lens/2;i++{
        sli[i],sli[lens-1-i] = sli[lens-1-i],sli[i]
    }
    return string(sli)
}
