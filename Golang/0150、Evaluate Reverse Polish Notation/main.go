func evalRPN(tokens []string) int {
    //栈，从数字末尾向前遍历，符号直接压入栈中，数字也是，只不过数字在压入栈中之后看是否凑够两个连续的数了，凑够了就开心消消乐，两个数加上后面的那个符号计算出结果基础压栈，继续消消乐
    //遍历完成 && 消消乐完成的数字就是最终的答案
    
    lens := len(tokens)
    stack := make([]string, 0)
    for i:=lens-1; i >=0; i-- {
        if !isNum(tokens[i]) {
            stack = append(stack, tokens[i])
        } else {
            //当前的是数字
            if len(stack) > 0 && isNum(stack[len(stack)-1]){
                //栈顶的也是个数字，，开始消消乐...
                nums := tokens[i]
                for len(stack) > 0 && isNum(stack[len(stack)-1]) {
                    //一直消消乐到没有两个连续的数字在栈中
                    tmpRes := getRes(nums, stack[len(stack)-1],stack[len(stack)-2])
                    nums = strconv.Itoa(tmpRes)
                    stack = stack[:len(stack)-2] //去掉刚刚消掉的两个
                }
                //这下可以把最后的结果压到栈中了
                stack = append(stack, nums)
            } else {
                //栈顶的是符号，直接压栈
                stack = append(stack, tokens[i])
            }
        }
    }
    res,_ := strconv.Atoi(stack[0]) 
    return res
    
}

func isNum(s string) bool {
    if s == string("+") || s == string("-") || s == string('*') || s == string("/") {
        return false
    }
    return true
}

func getRes(c, d string, notation string)int {
    a,_:=strconv.Atoi(c)
    b,_:=strconv.Atoi(d)
    switch(notation){
        case "+":
            return a+b
        case "-":
            return a-b
        case "*":
            return a*b
        case "/":
            return a/b
        default:panic("Invalid notation")
    }
    return 0
}
