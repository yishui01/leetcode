func romanToInt(s string) int {
    len_str := len(s)
   
    res:=0
    for i,j:=0,1;i<len_str;i,j=i+1,j*10{
        tmp := s[len_str-1-i]
        switch{
            case tmp == 'I':
            if len_str-i < len_str && (s[len_str-i] == 'V' || s[len_str-i] == 'X') {
                    continue
                } else {
                    res +=1
                }
            case tmp == 'V':
                if len_str-2-i >=0 && s[len_str-2-i] == 'I'{
                    res += 4
                } else {
                    res += 5
                }
            case tmp == 'X':
                if len_str-i < len_str && (s[len_str-i] == 'L' || s[len_str-i] == 'C') {
                    continue
                }
                if len_str-2-i >=0 && s[len_str-2-i] == 'I'{
                    res += 9
                } else {
                    res += 10
                }
            case tmp == 'L':
               if len_str-2-i >=0 && s[len_str-2-i] == 'X'{
                    res += 40
                } else {
                    res += 50
                }
            case tmp == 'C':
                if len_str-i < len_str && (s[len_str-i] == 'D' || s[len_str-i] == 'M') {
                    continue
                }
                if len_str-2-i >=0 && s[len_str-2-i] == 'X'{
                    res += 90
                } else {
                    res += 100
                }
            case tmp == 'D':
                if len_str-2-i >=0 && s[len_str-2-i] == 'C'{
                    res += 400
                } else {
                    res += 500
                }
            case tmp == 'M':
                if len_str-2-i >=0 && s[len_str-2-i] == 'C'{
                    res += 900
                } else {
                    res += 1000
                }
            
        }
    }
    return res
}