func canConstruct(ransomNote string, magazine string) bool {
    maps := make(map[rune]int)
    for _,v := range ransomNote {
        maps[v]++
    }

    for k,v := range maps {
        if strings.Count(magazine,string(k)) < v {
			return false
		}
    }

    return true

}
