func singleNumber(nums []int) int {
    //这个看了大佬解法才知道的，每一位出现的次数总和与3取模，出现三次的改为最后必定为0，出现一次的改位必定为1，golang的int默认是64位的...
    //具体是这样，每个数字都是64位的int类型，那就把每个数字按位比较，怎么按位比较？就是第一个数的1号位+第二个数的1号位+第三个数的3号位+....
    //加完之后得到一个1号位的总和，这个总和有个特点，就是和3取模的时候，要么是0要么是1.因为有个数字是出现一次的，假设这个数字的这个位为1，
    //那么其余数字都是出现三次，其余数字有该位为1的也有该位为0的，为1的加起来是3，多个为1的一起算就是3的倍数，为0的就是0，这些最后和3取模都会
    //被化成0，余数就是出现一次的那个数字的该位的值，求出那个数字该位的值之后，再把他左移到它该在的位置，res一直取 或，因为是一位一位的算，
    //所以不用考虑取或的时候 位重复相互干扰，所以最终res就是那个出现一次的数字一位一位的通过取 或 运算堆起来的
    lens := len(nums)
    var res int 
    var tmp int 
    for i:=0; i < 64; i++ {
        tmp = 0
        for j:=0; j < lens; j++ {
            tmp += (nums[j] >> uint(i)) & 1
        }
        res |=  (tmp % 3) << uint(i) //一个数在这一位出现3次那就是3，两个数在这位出现，那tmp总和就是6，反正最后会和3取模后会剩余个1或者0（0是因为那个出现一次的数该位也是0）
    }
    
    return res
}


//还有一个神级解法，反正我是没看懂
//这解法对于出现3次、6次、9次等等全部可以用这个,还没上面那个取模的好用呢，取模的改下模的值就能对应出现的次数
/**
func singleNumber(nums []int) int {
    var seen_once, seen_twice int
    seen_once, seen_twice = 0,0
    
    for _ , value := range nums {
        seen_once = (seen_once ^ value) &^ seen_twice 
        seen_twice = (seen_twice ^ value) &^ seen_once
    }
    
    return seen_once
}
**/
