func fractionToDecimal(numerator int, denominator int) string {
    //关键点在于怎么判断是否进入循环
    //如果逐位判断的话很难，暂时无法解决，
    //替代方案是用取模的方法进行判断，如果取到的模（余数）一样，那就是进入循环
    //还有，可以写成分数形式的都是有理数，有理数要么有限、要么是无限循环，所以不用考虑无限不循环小数
    if numerator == 0 || denominator == 0 {
        return "0"
    }
    num := int64(numerator)
    denum := int64(denominator)
    num = getAbs(num)
    denum = getAbs(denum)
    res := make([]byte,0)
    diviRes := make([]byte,0)
    if (numerator >0 && denominator < 0) || (denominator>0 && numerator <0){
        res = append(res,'-')
    }
    
    val := num/denum
    res = append(res,[]byte(strconv.Itoa(int(val)))...)
    reside := num%denum 
    maps := make(map[int64]int)
    isRepeat := false
    index := 0
    for reside != 0{
        //没有除尽
        if  val,ok := maps[reside];ok{
            //有相同的余数
           isRepeat = true
           index = val
           break
        }
      
        maps[reside] = index
        tmp := reside*10
        reside = tmp % denum
        tmp = tmp/denum
        diviRes = append(diviRes, byte(tmp+'0'))
        index++
    }
    if len(diviRes) == 0{
        //没有余数
        return string(res)
    }
    
    if isRepeat {
        //index代表商数组中，当前开始重复的索引值
        // r = diviRes[:index]+"("+ diviRes[index]+")"
        
        r :=make([]byte,index)
        copy(r,diviRes[:index])
        r = append(r,'(')
        r = append(r,diviRes[index:]...)
        r = append(r,')')
        res = append(append(res,'.'), r...)
    } else {
        res = append(append(res,'.'),diviRes...)
    }
    
     return string(res)
}

func getAbs(i int64) int64 {
    if i >=0 {
        return i
    }
    return -i
}
