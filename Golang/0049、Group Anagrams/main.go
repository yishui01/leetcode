func groupAnagrams(words []string) [][]string {
	cache := make(map[[26]byte]int)
	result := make([][]string, 0, 0)
	for _, word := range words {
		var z [26]byte
		for i := range word {
			z[word[i]-'a']++
		}
		if index, ok := cache[z]; ok {
			result[index] = append(result[index], word)
		} else {
			result = append(result, []string{word})
			cache[z] = len(result) - 1
		}
	}
	return result
}


/****************************下面是本菜鸟的解法，无法直视1300ms+************************/
/*func groupAnagrams(strs []string) [][]string {
    lens := len(strs)
    res := make([][]string, 0)
    if lens == 0 {
        return res
    }
    
    //一个一个对比
    for i:=0; i < lens; i++ {
        flag := false
        for j:=0;j < len(res); j++ {
            if flag == true {
                break
            }
            if compare(strs[i], res[j][0]) {
                flag = true
                res[j] = append(res[j], strs[i])
            }
        }
        
        if flag == false {
            res = append(res, []string{strs[i]})
        }
    }
    
    return res
   
}

func compare(a,b string) bool {
    lensa := len(a)
    lensb := len(b)
    if lensa == lensb {
        sa := []byte(a)
        sb := []byte(b)
        mysort(sa)
        mysort(sb)
        for i:=0; i < lensa; i++ {
            if sa[i] != sb[i] {
                return false
            }
        }
        return true
    }
    
    return false
}

func mysort(chars []byte){
    lens := len(chars)
    for i:=0; i < lens-1; i++ {
        for j:=0; j < lens-1-i;j++ {
            if chars[j]  > chars[j+1] {
                tmp := chars[j]
                chars[j] = chars[j+1]
                chars[j+1] = tmp
            }
        }
    }
    
    
}


*/