func plusOne(digits []int) []int {
    lens := len(digits)
    isFlow := 0
    for i:=lens-1; i >=0; i--{
        tmp := digits[i] + 1 + isFlow
        if i != lens-1 {
             tmp = digits[i] + isFlow
        }
        
        if tmp >= 10 {
            digits[i] = 0
            isFlow = 1
        } else {
            isFlow = 0
            digits[i] = tmp
            break
        }
    }
    
    if isFlow == 1 {
        digits = append([]int{1}, digits...)
    }
    
    return digits
}
