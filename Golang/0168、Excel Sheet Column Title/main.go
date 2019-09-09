func convertToTitle(n int) string {
  //从末尾往前逐位求解就行
    res := make([]byte,0)
    for n > 0 {
        n-- //这个真TM精髓取模
        res = append([]byte{byte('A'+n%26)},res...)
        //这里就不要把n再加上1了，因为加上来相除的话本来n是26，除完之后还有1，又来循环，本来是Z，现在变成AZ了，错误，所以n不用+1
        n/=26
    }
    
    return string(res)
    
}
