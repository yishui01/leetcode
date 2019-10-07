func calculate(s string) int {
/*解法一：利用栈，遇到加减符号直接将前一个数字push到栈中，遇到乘除就pop计算结果再push进去，最后把栈内的数字全部累加起来，
感觉这个方法还是比下面自己写出来的要好理解一些
*/
   var sign byte = '+'
    num := 0
    stack := make([]int,0)
    sum := 0
    
    s = strings.Replace(s," ","",-1) //去掉空格
    
    for i:=0;i<len(s);i++ {
        if s[i] >='0' && s[i] <= '9'{
            num = num*10+int(s[i]-'0') 
        }
        
        if i == len(s) - 1 || s[i] < '0' {
            switch sign {
                case '+': stack = append(stack,num)
                case '-': stack = append(stack,-num)
                case '*': stack[len(stack)-1] = stack[len(stack)-1] * num
                case '/': stack[len(stack)-1] = stack[len(stack)-1] / num
            }
            num = 0
            sign = s[i]
        }
        
    }
    
    for _,v := range stack {
        sum += v
    }
    
    return sum
    
}



//解法二：两个模式，一个加减模式一个乘除模式
//加减法的运算先保存到栈中
//乘除法直接进行运算
// func calculate(s string) int {
    
//     isAdd := 1 //符号位
//     sum := 0 
//     tmpMul := 0
//     num :=0
    
//     for i:=0;i<len(s);i++ {
//         if s[i] == ' '{
//             continue
//         }
//         if s[i] >= '0' && s[i] <= '9' {
//             num = num*10 + int(s[i]-'0') 
//         } else {
//             //到这里来num就是运算符之前的那个数字
//              if isAdd == 1 || isAdd == -1 {
//                 num *= isAdd //先把正负号加上
//             }
            
//             if s[i] == '+' || s[i] == '-' {
//                  if isAdd == 1 || isAdd == -1 { //如果之前的也是+或者-，那就直接累加num的值
//                          sum += num
//                     } else { //否则把之前的乘除法表达式的值累加上去
//                         if isAdd == 2{
//                             sum += tmpMul*num
//                         } else {
//                             sum += tmpMul/num
//                         }
//                 }
//             } else {
//                 if isAdd == 1 || isAdd == -1 { //如果是乘除法表达式的第一个数字，那就直接复制
//                        tmpMul =  num
//                     } else { //否则就进行累计
//                           if isAdd == 2{
//                            tmpMul *= num
//                         } else {
//                            tmpMul /= num
//                         }
//                     }
//             }
            
//             switch s[i]{
//                 case '+':
//                      isAdd = 1
//                 case '-':
//                      isAdd = -1  
//                 case '*':
//                     isAdd = 2
//                 case '/':
//                     isAdd = -2
//             }
//            num = 0
//         }
//     }
    
//     //将最后一个操作数累计到sum中
//     if isAdd == 1{
//         sum += num
//     } else if isAdd == -1 {
//          sum -= num
//     }else if isAdd == 2 {
//         sum += tmpMul * num
//     } else {
//         sum += tmpMul / num
//     }

//     return sum
    
// }
