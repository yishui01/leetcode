func reverseWords(s string) string {
    //这...要O1的额外空间，golang里面是字符串常量、、不能原地交换、、做不到，很伤
    //最简单的新开一个字符串数组，把字单词一个个加进去，最后拼接起来
    lens := len(s)
    if lens == 0 {
        return s
    }
    left := lens - 1
    right := lens - 1
    words := make([]string, 0)

    for i:=lens-1; i >=0; i-- {
        if s[i] == ' ' { 
            left--
            right--
        } else {
            if s[right] == ' '{
                right-- //保证right永远指向的是单词的末尾而不是空格
            }
            left--
        }
         if left >=0 && right >= 0 && s[right] != ' ' && s[left] == ' ' { 
             words = append(words, s[left+1:right+1])
             right = left
        }
    }
    if s[0] != ' '{
        words = append(words, s[:right+1])
    }
    return strings.Join(words," ")
}

/***
func reverseWords(s string) string {
    words := strings.Split(s, " ")
    ans := make([]string, 0, len(words))
    for i := len(words)-1; i >= 0; i-- {
        if words[i] != "" {
            ans = append(ans, words[i])
        }
    }
    return strings.Join(ans, " ")
}
***/
