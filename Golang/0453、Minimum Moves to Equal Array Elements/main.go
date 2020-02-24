
//O(nlogn)
//思考增加过程，得出总次数=最小的数和每一个数字的差值 累加
//先排序，再计算
/*func minMoves(nums []int) int {
    sort.Ints(nums)
    res := 0
    for i:=0;i<len(nums);i++{
        res += nums[len(nums)-1-i]-nums[0]
    }
    return res
}*/


//上面的方法需要排序， 想下是否可以不排序直接计算所有差值
//直接O(n)找出最小值，同时将所有数字累加，最后减去最小值*n，是不是就是上面的方法计算的结果
func minMoves(nums []int) int {
    m := math.MaxInt32
    res := 0
    for _,v:= range nums {
        if v < m {
            m = v
        }
        res += v
    }
    return res-m*len(nums)
}
