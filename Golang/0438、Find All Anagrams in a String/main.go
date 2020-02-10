func findAnagrams(s string, p string) []int {
    charIndexP := make([]int, 26)
	charIndexS := make([]int, 26)
	res := make([]int, 0, 16)
	for i := 0; i < len(p); i++ {
		charIndexP[p[i]-'a']++
	}

	// [l, r]
	l, r := 0, -1
    prev := false
	for r+1 < len(s) {
		r++
		charIndexS[s[r]-'a']++

		if r-l+1 > len(p) { //窗口过大，收缩左边界
			charIndexS[s[l]-'a']--
			l++
		}

		if r-l+1 == len(p) {
            if prev { //上一步是合格的话，本次只要判断左边移除的和右边新加的是不是同一个字符
                if s[l-1] == s[r] {
                    res = append(res, l)
                } else {
                    prev = false
                }
                continue
            }

			sign := true
			for i := 0; i < 26; i++ {
				if charIndexS[i] != charIndexP[i] {
					sign = false
					break
				}
			}
			if sign {
				res = append(res, l)
			}
            prev = sign
		}

	}

	return res 
}
