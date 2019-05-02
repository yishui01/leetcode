func isValid(s string) bool {
    lens := len(s)
    if lens == 0 {
        return true
    }
    //数据结构、栈、先进后出
    maps := map[byte]byte{
        ')':'(',
        ']':'[',
        '}':'{',
    }
    stack := make([]byte, lens)
    i:=0
    for _,v := range s  {
        tmp := byte(v)
        if tmp == '(' || tmp == '[' || tmp == '{' {
            //左括号入右括号出
            stack[i] = tmp
            i++
        } else {
            if i >= 1 && stack[i-1] == maps[tmp] {
                i--
            } else {
                return false
            }
        }
    }
    
    if i == 0 {
        return true
    }
    return false
    
}
