func myAtoi(str string) int {
    //遍历字符串，将数字存到另一个数组中
    //再遍历另一个数组，累加转换
    lenS := len(str)
    if lenS <= 0 {
        return 0
    }
    is_start := 0 //是否开始记录
    sign := 1 //符号位，1为正，-1为负数
    
    num_arr := make([]byte, 0)
    for i:=0; i<lenS; i++{
        if is_start == 0 {
            //未开始记录时，跳过空格，遇到字符直接G
            if str[i] == ' '{
                continue
            } else if(str[i] == '+' || str[i] == '-') {
                if str[i] == '-' {
                    sign = -1
                }
                is_start = 1
            } else if(str[i] >= 48 && str[i] <= 57) {
               //数字
                num_arr = append(num_arr, str[i])
                is_start = 1
            } else {
                 return 0
            }
        } else {
            //开始记录了,除了数字以外的任何字符都跳出循环，包括空格
            if(str[i] >= 48 && str[i] <= 57) {
               //数字
                num_arr = append(num_arr, str[i])
            } else {
                 break
            }
        }
    }
    
    lenNum := len(num_arr)
    res := 0
    //开始累加数组
    for j:=0; j < lenNum; j++ {
        res+= (int(num_arr[j])-int('0'))*int(math.Pow(10,float64(lenNum-j-1)))
        if (num_arr[j] != '0' && math.Pow(10,float64(lenNum-j-1)) > float64(math.MaxInt32)) || res > math.MaxInt32 || res < math.MinInt32 {
            if sign == 1 {
                return math.MaxInt32
            }
            return math.MinInt32
        }
    }
    
    return res*sign
    
}
