func isHappy(n int) bool {
    maps := make(map[int]bool)
    return helper(n,maps)
}

func helper(n int,maps map[int]bool) bool {
    if n == 0 {
        return false
    } else if n==1{
        return true
    }
    
    if _,ok := maps[n];ok{
        return false
    }
    maps[n] = true
    
    next := 0
    for n > 0 {
        next += (n%10)*(n%10)
        n /= 10
    }
    
    return helper(next, maps)
}
