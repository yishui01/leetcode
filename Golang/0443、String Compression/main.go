func compress(chars []byte) int {
    lens := len(chars)
    if lens <= 1 {
        return lens
    }
    
    res := 1
    nowChar := chars[0]
    nowCount := 1
    for i:=1;i<lens;i++{
        if chars[i] == nowChar {
            nowCount++
        } else {
            if nowCount > 1 {
                strNum := strconv.Itoa(nowCount)
                for _,v := range strNum {
                    chars[res] = byte(v)
                    res++
                }
            }
          
            chars[res] = chars[i]
            res++
            nowChar = chars[i]
            nowCount = 1
        }
    }

    if nowCount > 1 {
            strNum := strconv.Itoa(nowCount)
            for _,v := range strNum {
                chars[res] = byte(v)
                res++
            }
    }

    return res

}
