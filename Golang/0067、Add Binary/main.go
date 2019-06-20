func addBinary(a string, b string) string {
    lenA := len(a)
    lenB := len(b)
    
    sa := []byte(a)
    sb := []byte(b)
    isFlow := 0
    
    res := make([]byte, 0)
    
    for lenA-1 >=0 || lenB-1 >=0 {
        var tmp byte
        if lenA-1 >=0 && lenB-1 >=0 {
            tmp = sa[lenA -1] - '0' + sb[lenB -1] - '0' + byte(isFlow)
        } else if lenA-1 >=0 {
            tmp = sa[lenA -1] - '0' + byte(isFlow)
        } else if lenB-1 >=0 {
            tmp = sb[lenB -1] - '0' + byte(isFlow)
        }
        if tmp == 2 {
            tmp = 0
            isFlow = 1
        } else if tmp == 3 {
             tmp = 1
            isFlow = 1
        } else {
            isFlow = 0
        }
        res = append([]byte{tmp+'0'}, res...)
        
        lenA--
        lenB--
    }
    
    if isFlow == 1 {
        res = append([]byte{'1'}, res...)
    }
    
    return string(res)
    
}
