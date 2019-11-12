func isAdditiveNumber(num string) bool {
    if len(num) < 3{
        return false
    }
    
    if num[0] == '0' {
        for i:=1;i<len(num);i++{
            if helper(num[i+1:],"0",num[1:i+1]) { 
                return true
            }
        }
    } else {
        for i:=0;i<len(num)-1;i++{
            if num[i+1] == '0' {
                if helper(num[i+2:],num[:i+1],"0"){
                    return true
                }
            } else {
                for j:=i+1;j<len(num);j++{
                   //fmt.Println(num[:i+1],"---",num[i+1:j+1],"---",num[j+1:])
                    if helper(num[j+1:],num[:i+1],num[i+1:j+1]){ //这里有可能传个空的s进去
                        return true
                    } 
                }
            }
        }
    }
    
    return false
    
}

func helper(s,n1,n2 string) bool {
    if len(s) < len(n1) || len(s) < len(n2) { //所以这里不能通过判断s为空来返回true
        return false
    }
    
    num1,_ := strconv.Atoi(n1)
    num2,_ := strconv.Atoi(n2)
    sumStr := strconv.Itoa(num1+num2)
    
    if len(s) < len(sumStr) {
        return false
    }
    
    if s[:len(sumStr)] == sumStr {
        if len(s) == len(sumStr) { //要在这里判断，刚好匹配完就直接返回
            return true
        }
        //fmt.Println("下一轮",n2,sumStr,s[len(sumStr):])
        return helper(s[len(sumStr):],n2,sumStr)
    }
    
    return false
    
}
