func diffWaysToCompute(input string) []int {
  //按照符号分割，分而治之
    res := make([]int,0)
    for i:=0;i<len(input);i++{
        if input[i] == '+' ||input[i] == '-' || input[i] =='*'{
            //分割成左右两部分，递归计算
            left := diffWaysToCompute(input[:i])
            right := diffWaysToCompute(input[i+1:])
            
            //到这里开始计算
            for z:=0;z<len(left);z++{
                for k:=0;k<len(right);k++{
                    switch input[i]{
                        case '+':
                        res = append(res,left[z]+right[k])
                        case '-':
                        res = append(res,left[z]-right[k])
                        case '*':
                        res = append(res,left[z]*right[k])
                    }
                }
            }
            
        }
    }
    
    //如果res为空，代表s字符串本身就是个数字，直接加进去即可
    if len(res) == 0 {
        num,_ := strconv.Atoi(input)
        res = append(res,num)
    }
    
    return res
    
}
