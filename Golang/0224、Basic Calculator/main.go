func calculate(s string) int {
    rangeSign := 1
    nowSign := 1
    sum := 0
    tmp := 0
    stack := make([]int,0)
    
    for i:=0;i<len(s);i++{
        if s[i] >= '0' && s[i] <= '9' {
            tmp = tmp*10 + int(s[i]-'0')
        } else {
            sum += rangeSign * nowSign * tmp
            tmp = 0
            
            switch(s[i]){
                case '+':
                    nowSign = 1
                case '-':
                    nowSign = -1
                case '(':
                    //现在要进入括号领域，要把括号之前的范围sign保留下来
                    stack = append(stack,rangeSign)
                    //再把当前的符号累计起来，用于得出负负得正
                    rangeSign *= nowSign
                    //累计完成之后消除nowSign的影响
                    nowSign = 1
                case ')':
                    //括号结束，还原rangeSign
                rangeSign = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
            }
        }
        
    }
    
    //把最后一个数字加上去
    sum += rangeSign * nowSign * tmp
    
    return sum
}
