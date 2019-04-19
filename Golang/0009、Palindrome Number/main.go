func isPalindrome(x int) bool {
   //这题后面要求不用转换成字符换的方式，感觉有点困难，看了大神解法，get
    if x < 0 {
        return false
    }
    if x < 10 {
        return true
    }
    reverse := 0
    y := x
    for y> 0 {
        reverse = reverse * 10 + y%10
        y/=10
    }
    
    return reverse == x
    
}
