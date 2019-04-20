func isMatch(s string, p string) bool {
    //递归解决
    lenS,lenP := len(s), len(p)
    if lenP == 0 {
        //正则都没了，看字符串还有没有，还有就G了
        return lenS == 0
    }
    if lenS == 0 && lenP == 0{
        return true
    }
    if lenP == 1 {
        return lenS == 1 && (p[0] == s[0] || p[0] == '.')
    }
   
    //接下来是正则表达式字符大于等于2的时候
    if p[1] != '*' {
        if lenS == 0 {
            return false
        }
        return (p[0] == '.' || p[0] == s[0]) && isMatch(s[1:], p[1:])
    }
    //如果第二个字符是*,要考虑*重复的次数，有可能是0次，也有可能是多次
    for {
        if(len(s)>0 && (p[0] == '.' || p[0] == s[0])) {
            if (isMatch(s,p[2:])) {
            //如果当前字符串能完美匹配之后的正则，那这个循环就没卵用，*重复次数为0
            //否则进行到下一次循环，代表至少重复了一次，s已经前进了一位
                return true
            }
       } else {
           //直到s前进到的字符不处于这个重复字符，那就跳出循环，将s继续和剩下的正则进行匹配
           break
       }
        s = s[1:]
    }
    //到最后有一段是循环模式匹配不了的，s已经在for中脱离了循环匹配的范围，将正则也后移两位（移到*的后面）继续匹配
    return isMatch(s,p[2:])
    
}
