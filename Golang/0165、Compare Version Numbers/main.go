func compareVersion(version1 string, version2 string) int {
    //以小数点为分隔符逐个比较，之前还加了去掉前导0的操作，然而golang可以把0008这种直接转成8，那就没必要去掉前导0了，直接进行比较即可
    sli1 := strings.Split(version1,".")
    sli2 := strings.Split(version2, ".")
    
    len1 := len(sli1)
    len2 := len(sli2)
    
    for i:=0;i < len1; i++ {
        n1,_:=strconv.Atoi(sli1[i])
        if len2 > i {
            n2,_:=strconv.Atoi(sli2[i])
            if n1 > n2 {
                return 1
            } else if n1 < n2 {
                return -1
            }
        } else {
            if n1 != 0 {
                return 1
            }
        }
    }
    
    if len2 > len1 {
        for i:=len1; i < len2; i++ {
            n2,_:=strconv.Atoi(sli2[i])
            if n2 != 0 {
                return -1
            }
        }
    }
    
    return 0
}
