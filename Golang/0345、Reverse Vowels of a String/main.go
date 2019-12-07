func reverseVowels(s string) string {
    sli := []byte(s)

    left,right := 0,len(s)-1
    
    isFind := false
    for left < right {
        if isFind{
            if isVow(sli[right]){
                sli[left],sli[right] = sli[right],sli[left]
                isFind = false
                left++
            }
            right--
        } else {
            if isVow(sli[left]){
                isFind = true
            } else {
                left++
            }
        }
    }
    return string(sli)
}

func isVow(c byte) bool {
    arr := []byte{'a','e','i','o','u','A','E','I','O','U'}
    for _,v := range arr {
        if v == c {
            return true
        }
    }

    return false
}
