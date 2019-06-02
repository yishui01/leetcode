func isMatch(s string, p string) bool {
  
    iStart := -1
    pStart := -1
    lenp := len(p)
    lens := len(s)
    i:=0
    j:=0
    for i < lens {
        if j<lenp && (s[i] == p[j] || p[j] == '?'){
            i++
            j++
        } else if j<lenp && p[j] == '*' {
            iStart = i
            pStart = j
            j++
        } else if iStart >=0 {
            j=pStart+1
            i=iStart+1
            iStart++
        } else {
            return false
        }
    }
    for z:=j; z < len(p); z++{
        if p[z] != '*' {
            return false
        }
    }
    
    return true
}
