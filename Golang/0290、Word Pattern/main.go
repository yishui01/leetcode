func wordPattern(pattern string, str string) bool {
        maps := make(map[byte]string)
        mapsStr := make(map[string]bool)
        strs := strings.Split(str," ")
        
        if len(pattern) != len(strs) {
            return false
        }
        
        for k,v := range strs{
            if val,ok := maps[pattern[k]];ok {
                if val != v {
                    return false
                }
            } else {
                if _,ok := mapsStr[v];ok {
                    return false
                }
                maps[pattern[k]] = v
                mapsStr[v] = true
            }
        }

    return true
    
}
