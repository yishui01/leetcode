func addStrings(num1 string, num2 string) string {
    res := make([]byte,0)
    up := 0 //是否进位
    for i:=0;i<len(num1) || i < len(num2);i++{
        n1,n2 := 0,0
        if i < len(num1){
            n1 = int(byte(num1[len(num1)-1-i]) - '0')
        }
        if i < len(num2){
            n2 = int(byte(num2[len(num2)-1-i]) - '0')
        }
        sum := n1+n2+up
        if sum >= 10 {
            up = 1
            sum %= 10
        } else {
            up = 0
        }
        res = append(res,byte(sum + '0'))
    }

    if up == 1 {
        res = append(res,'1')
    }

    for i:=0;i<len(res)/2;i++{
        res[i],res[len(res)-1-i] = res[len(res)-1-i],res[i]
    }

    return string(res)
}
