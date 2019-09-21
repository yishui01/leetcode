func reverseBits(num uint32) uint32 {
    //从后往前逐位取出，
    //如果这位是1，就把1左移到相应的位置，与res取或
    var res uint32 = 0
    for i:=0;i<32;i++{
        now := num >> uint32(i) & 1
        if now == 1 {
            res = res | 1 << uint32(32-i-1)
        }
    }
    return res
}
