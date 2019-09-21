func hammingWeight(num uint32) int {
    //这....有难度？
    res := 0
    if num == 0 {
        return 0
    }
    for i:=0;i<32;i++{
        if num >> uint32(i) & 1 == 1 {
            res++
        }
    }
    
    return res
}
