func validUtf8(data []int) bool {
    //两个模式：一个是独立模式、一个是尾部模式（一个字符未完时）
    //用一个变量++来记录当前还需要找到几个尾部，为0则为独立模式，大于0则为尾部模式
    sum := 0
    var tmp [8]int
    for _,v := range data {
        for i:=0;i<8;i++{
            tmp[7-i] = v & 1 //从后往前填充，保证顺序
            v = v >> 1
        }
       if sum == 0 {
           //独立模式
           if tmp[0] == 0 {
               //还是独立模式
               continue
           } else {
               tmpCount := 0 //首位为1的个数
               j := 0
               for tmp[j] != 0 {
                   tmpCount++
                   j++
                   if tmpCount > 4 {
                        return false
                    }
               }
               if tmpCount < 2 {
                    return false
               }
               sum = tmpCount-1
           }
       } else {
           //尾部模式
           if tmp[0] != 1 || tmp[1] != 0 {
               return false
           }
           sum--
       }
    }

    return sum == 0
    
}
