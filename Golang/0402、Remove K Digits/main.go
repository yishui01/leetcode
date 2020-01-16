func removeKdigits(num string, k int) string {
    if len(num) == k {
        return "0"
    }

    //看答案用栈，新数字比栈顶大直接压栈，比栈顶小直接弹出，弹完继续比较下一个栈顶数字，直到栈空或者栈顶小于新数字，将新数字压栈
    b := []rune(num)
    stack := make([]rune,0)
    popCount := 0

    for k,v := range b {
         if popCount >= k {
                break
         }
        if len(b) == 0 || b[len(b)-1] < v {
            stack = append(stack,v)
        } else {
            for len(b) != 0 && b[len(b)-1] <= v { //等于栈顶也直接弹出
                b = b[:len(b)-1]
                popCount++
            }
        }
    }

    if popCount != k {
        k-popCount
    }

}
