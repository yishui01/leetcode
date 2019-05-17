func longestValidParentheses(s string) int {
    //这个貌似很简单,左括号入栈右括号出栈,注意题意是要你返回最长的那个子串，不是全部的
    //每次遍历到左括号的时候，将下标，记住是index入栈，
    //遇到右括号的时候，尝试pop，如果pop失败，那么将下一个元素的值作为start，
    //pop成功后，看栈是否为空，为空则i-start+1比对max，不为空则i-栈顶元素下标比对max
    
    count := 0
    start := 0
    maxCount :=0
    arr := make([]int, 0)
    
    for k,v := range s {
        if v == '(' {
            arr = append(arr, k) 
        } else {
            lens := len(arr)
            if lens > 0 {
                //有pop
                arr = arr[:lens-1]
                if len(arr) == 0 {
                    //栈为空
                    count = k-start+1
                } else{
                     count = k-arr[len(arr)-1]
                }
                if maxCount < count {
                    maxCount = count
                }
            } else {
                //无pop，下一个元素为start
                start = k+1
            }
           
        }
        
    }
    
    return maxCount
}
