func intToRoman(num int) string {
     m1:= map[int]byte{
         1: 'I',
         5: 'V',
         10:'X',
         50:'L',
         100:'C',
         500:'D',
         1000:'M',
     }
    
    res_arr := make([]byte,0)
    //从后往前，一个个算
    for i:=1;num > 0;i*=10{
        div := num % 10
        
        switch{
            case div == 0 :
                
            case div >=1 && div <=3:
                for j:=0;j<div;j++{
                    res_arr = append(res_arr,m1[i])
                } 
            case div == 4:
                res_arr = append(res_arr,m1[i*10/2])
                res_arr = append(res_arr,m1[i])
            case div >=5 && div <=8:
                for j:=0; j < div-5;j++ {
                    res_arr = append(res_arr,m1[i])
                }
                res_arr = append(res_arr,m1[i*5])
            case div == 9:
                res_arr = append(res_arr,m1[i*10])
                res_arr = append(res_arr,m1[i])
        }
        num  = num/10
    }
    
    return reverse(res_arr)
}


func reverse(str []byte)string {
    len_str := len(str)
    for i:=0;i < len_str/2;i++ {
        tmp := str[i]
        str[i] = str[len_str-1-i]
        str[len_str-1-i] = tmp
    }
    return string(str)
}
