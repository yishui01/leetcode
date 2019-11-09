func removeInvalidParentheses(s string) []string {
    //bfs，一个字符一个字符的移除，每移除一个字符产生一个新的字符串看是否合法，合法加到res中，不合法就继续入队
    //直到第一个res找到，那么此时确定了res的长度（所有res的长度必定一样），可以对接下来遍历的字符串进行剪枝，小于等于res又不合法的字符串
    //就不用继续移除一个字符然后入队了，直接跳过
    //注意：入队时需要检查该字符串是否已经出现过，避免重复
    
    queue := []string{s}
    maps := make(map[string]bool)
    
    res := []string{}
    resLen := -1
    
    for len(queue) > 0{
        t := queue[0]
        queue = queue[1:]
        if isValid(t){
            res = append(res,t)
            resLen = len(t)
            continue
        }
        if len(t) <= resLen { //这时候t不管怎么删都不合法了
            continue
        }
        for i:=0;i<len(t);i++{
            if t[i] != '(' && t[i] != ')' {
				continue
			}
            //顺序删除t中的某一个字符
            tc := t[:i]+t[i+1:]
            if _,ok := maps[tc];!ok {
                maps[tc] = true
                queue = append(queue,tc)
            }
        }

    }
    
    return res
}

func isValid(s string) bool {
    count:=0
    for i:=0;i<len(s);i++ {
        if s[i] == '('{
            count++
        } else if s[i] == ')'{
            count--
        }
        if count < 0 {
            return false
        }
    }
    
    return count == 0
}
