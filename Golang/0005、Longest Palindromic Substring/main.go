func longestPalindrome(s string) string {
    
    //解题思路：从每个字符开始往两边扩散，每次扩散都比对字符两边扩展的新字符是否相等
    
    var maxLen int = 0;
    var long_start = 0;
    var long_end = 0;
    
    if len(s) < 2 {
        return s
    }
    
    for i:=0; i<len(s); i++{
        findPal(&maxLen, &s, i,i,&long_start,&long_end)
        findPal(&maxLen, &s, i,i+1,&long_start,&long_end)
    }
    return s[long_start:long_end+1]
    
}

func findPal(maxLen *int, s *string, start int, end int, long_start *int, long_end *int) {
    for {
        if start >=0 && end < len(*s) && (*s)[start] == (*s)[end] {
            start--;
            end++;
        } else {
            if end-start-1 > *maxLen {
                *maxLen = end-start-1
                *long_start = start+1
                *long_end = end-1
            }
            break;
        }
    }
    
}