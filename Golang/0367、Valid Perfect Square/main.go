func isPerfectSquare(num int) bool {
    n := 0
    for n*n <= num {
        if n*n == num {
            return true
        }
        n+=1
    }
    return false
}
