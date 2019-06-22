func mySqrt(x int) int {
    if x <= 1 {
        return x
    }
    
    //二分吧,平方根只会小于等于x/2,所以在(1,x/2] 这个区间找,
    
    
    left := 0
    right := x
    mid := 0
    for left <= right {
        mid = (left + right) / 2
        sum := mid * mid
        if sum == x {
            return mid
        }
        if sum < x {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    
    //最后mid就是答案
    if mid * mid > x {
        return mid - 1
    }
    return mid
    
    
}
