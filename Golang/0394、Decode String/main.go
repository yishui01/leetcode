func decodeString(s string) string {
    //虽然没做过，但是一看就是个dfs,天然dfs ╮(╯▽╰)╭
    res := ""
    for i:=0;i<len(s);i++{
        if s[i] >= '1' && s[i] <= '9' {
            //数字,找出数字的全部部分
            start := i+1
            for s[start] >= '0' && s[start] <= '9' {
                start++
            }
            num,_ := strconv.Atoi(s[i:start])
            i = start-1 //将i移到数字的最后一位，准备递归

            //找出当前数字管辖的范围
            left := 0
            j := i+1;
            for  j<len(s) {
                if s[j] == '[' {
                    left++
                } else if s[j] == ']' {
                    left--
                }
                if left == 0 {
                    break;
                }
                j++
            }

            
            //拆出这个大段的索引为i+2 ~ j-1 进行dfs
            for num > 0 {
                res += decodeString(s[i+2:j])
                num--
            }

            //这里将i移动到右括号位置处，因为for循环还会自动i++，刚好移动到下一位新的字符串上
            i = j
        } else {
            //不是数字，这里也遇不到括号，因为全部被摘除了，直接拼接
            res += string(s[i])
        }
    }
    return res 
}
