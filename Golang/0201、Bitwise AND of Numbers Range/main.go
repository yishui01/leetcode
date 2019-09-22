//这题是找出m和n左边相同的公共部分，知道这个就好办了
//解析：m在+1的过程中不断向n靠近，二进制中每次+1都会使得末尾+1或者进位，那么只要有进位，就会出现0,0与任何数相与都是0，所以要找到没有被进位影响的高位
//高位后面填0就是结果

//解法一：通过一个mask值来找公共部分
// func rangeBitwiseAnd(m int, n int) int {
//     d := math.MaxInt32
//     for (m & d) != (n & d){
//         d=d << 1
//     }
//     return m & d
    
// }

//解法二，m，n自己不断右移，相等的时候就是结果
func rangeBitwiseAnd(m int, n int) int {
    var i uint
    for m != n{
       m = m >> 1
       n = n >> 1
       i++ 
    }
    return m << i //最后要把公共位回移到对应的高位
}

// //解法三，神级递归，原理和解法二差不多，也是一直递到n <= m的时候就归，并在每层return 的时候直接把位置还原回来，很6
// func rangeBitwiseAnd(m int, n int) int {
//     if n > m {
//        return  rangeBitwiseAnd(m >> 1, n >> 1) << 1
//     }
//     return m
// }

