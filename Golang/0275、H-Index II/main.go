func hIndex(citations []int) int {
    //一说要用log就是二分
    lens := len(citations)
    left := 0
    right := lens - 1
    
    for left <= right {
        mid := (left + right) / 2
        if citations[mid] == lens-mid {
            return lens-mid
        }
        
        if citations[mid] > lens - mid {
            mid -= 1
            right = mid
        } else {
            mid += 1
            left = mid
        }
    }
    
    return lens-left
}
