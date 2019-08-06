

func isPalindrome(s string) bool {
    //这...一头一尾进行比对就行了啊...很难吗？
    //。。。好吧做完之后我才发现要想bug free还真是不简单，好容易就漏了
    lens := len(s)
    if lens <= 1 {
        return true
    }
    left := 0
    right := lens-1
    
    
    for i:=0; i < lens; i++ {
        if right <= left {
            return true
        }
        if (s[left] < '0' || s[left] > '9') &&  (s[left] < 'a' || s[left] > 'z') && (s[left] < 'A' || s[left] > 'Z'){
            //left不及格
            left++
            continue
        }
        if (s[right] < '0' || s[right] > '9') &&  (s[right] < 'a' || s[right] > 'z') && (s[right] < 'A' || s[right] > 'Z'){
            //right不及格
            right--
            continue
        }
        
        if s[left] == s[right] {
            left++
            right--
            continue
        }
        
        if (s[left] >= 'a' && s[left] <= 'z') || (s[left] >= 'A' && s[left] <= 'Z'){
            if s[right] >= 'a' && s[right] <= 'z'{
                if s[right] - (97-65) != s[left]{
                    return false
                } else {
                     left++
                    right--
                    continue
                }
            } else if s[right] >= 'A' && s[right] <= 'Z'{
                if s[right] + (97-65) != s[left]{
                    return false
                } else {
                     left++
                     right--
                     continue
                }
            } else {
                return false
            }
           
        } else {
            return false
        }
    }
    
    
    return true
}
/**
//看了下大佬的算法，就是这么清晰简洁，真的功力还是差太远了,这里有个细节，函数参数是rune类型
func isAlphanumeric(c rune) (rune, bool){
    if c >= '0' && c <= '9' {
        return c, true
    }else if c >= 'a' && c <= 'z' {
        return c, true
    }else if c >= 'A' && c <= 'Z' {
        return c-'A'+'a', true
    }
    return c, false
}

func isPalindrome(s string) bool {
    
    i,j := 0, len(s)-1
    for j > i {
        var front, end rune
        if c, ok := isAlphanumeric(rune(s[i])); ok{
            front = c
        }else {
            i++
            continue
        }
        if c, ok := isAlphanumeric(rune(s[j])); ok{
            end = c
        }else {
            j--
            continue
        }
        if front != end {
            return false
        }else{
            i++
            j--
        }
    }

    return true
}
**/
