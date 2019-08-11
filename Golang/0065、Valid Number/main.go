func isNumber(s string) bool {
    if len(s) == 0 {
        return false
    }
    s = trim(s) //移除首尾空格
    seeNum := false //是否为数字
    seeDot := false //是否有逗号 
    seeE := false //是否有e
    
    lens := len(s)
   
    for i:=0; i < lens; i++ {
        if s[i] >= '0' && s[i] <= '9' {
            seeNum = true
        } else if s[i] == 'e' { 
            if !seeNum || seeE { //e不能重复出现且只能跟在数字后面
                return false
            }
            seeE = true
            seeNum = false //e后面必须要接一个数字才算合法
        } else if s[i] == '+' || s[i] == '-' {  //+ - 只能出现在最前面或者e的后面
            if i !=0 && s[i-1] != 'e' {
                return false
            }
        } else if s[i] == '.' { //.的位置可以任意，但是不能重复出现可出现在e的后面
            if  seeDot || seeE {
                return false
            }
            seeDot = true
        } else {
            return false
        }
    }
    
    return seeNum
}

func trim(s string) string {
    lens := len(s)
    left := 0
    leftNum := false
    right := lens-1
    rightNum := false
    for i:=0; i < lens; i++ {
        if left >= right {
            break
        }
        if s[i] == ' ' && !leftNum{
            left++
        } else {
            leftNum = true
        }
        if s[lens-1-i] == ' ' && !rightNum{
            right--
        } else {
            rightNum = true
        }
    }
  
    return s[left:right+1]
}
