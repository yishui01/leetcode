func canMeasureWater(x int, y int, z int) bool {
    if z == 0 {
        return true
    }
   if z <= x+y {
        return z % gcd(x,y) == 0
   }
    return false
}

func gcd(a, b int)int{
    if b == 0 {
        return a
    }
    return gcd(b,a%b)
}
