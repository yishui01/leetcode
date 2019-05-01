func letterCombinations(digits string) []string {
    if len(digits)==0{
        return []string{}
    }
	var res []string
	maps := make(map[string]string)
	maps["2"] = "abc"
	maps["3"] = "def"
	maps["4"] = "ghi"
	maps["5"] = "jkl"
	maps["6"] = "mno"
	maps["7"] = "pqrs"
	maps["8"] = "tuv"
	maps["9"] = "wxyz"
	dfsForPhone("", strings.Split(digits, ""), &res, maps)
	return res

}
func dfsForPhone(c string, lastS []string, res *[]string, maps map[string]string) {
	if len(lastS) == 0 {
		*res = append(*res, c)
		return
	}
	digit := string(lastS[0])
	letters := maps[digit]
	for i := 0; i < len(letters); i++ {
		letter := string(letters[i])
		dfsForPhone(c+letter, lastS[1:], res, maps)
	}
}

/*
func letterCombinations(digits string) []string {
    //组合的全排列
    if len(digits) == 0 {
        return []string{}
    }
    maps := map[byte][]string{
        '2':[]string{"a","b","c"},
        '3':[]string{"d","e","f"},
        '4':[]string{"g","h","i"},
        '5':[]string{"j","k","l"},
        '6':[]string{"m","n","o"},
        '7':[]string{"p","q","r","s"},
        '8':[]string{"v","t","u"},
        '9':[]string{"x","w","y","z"},
    }
    pre,_:= maps[digits[0]]
    
    for k,v := range digits {
        if k == 0 {
            continue
        }
        pre = join(pre,maps[byte(v)])
    }
    return pre
}

func join(pre,item []string)[]string{
    res := make([]string,0)
    for k,_ := range pre {
        for _,v2:=range item {
            res = append(res,pre[k]+v2)
        }
    }
    return res
}
*/