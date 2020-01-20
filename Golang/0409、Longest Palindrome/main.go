func longestPalindrome(s string) int {
    //一开始题目都看错了
    maps:=make(map[rune]int)
    for _,v := range s {
        maps[v]++
    }

    res := 0

    for _,v := range maps {
        res += v/2*2
        if res % 2 == 0 && v % 2 == 1 { //res正好是偶数个字符 && v 拥有奇数个字符，那就加一个中间字符。注意这里res 加了一次之后就不会再进入这个if了，因为上面res+=偶数
            res++
        }
    }

    return res
}
