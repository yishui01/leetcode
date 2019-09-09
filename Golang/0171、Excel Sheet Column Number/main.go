func titleToNumber(s string) int {
    //...这和之前那题有区别？
    res := 0
    p := 1
    for i:=len(s)-1;i >=0;i--{
        res += int(s[i]-'A'+1)*p
        p*=26
    }
    return res
}

// func titleToNumber(s string) int {
// 	res := 0
// 	for _, i := range s {
// 		res *= 26
// 		res += int(i - 'A' + 1)

// 	}
// 	return res
// }
