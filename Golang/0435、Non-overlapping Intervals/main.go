func eraseOverlapIntervals(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }
    //按照结尾，从小到大排序
    sort.Slice(intervals,func(i,j int)bool{
        return intervals[i][1] < intervals[j][1]
    })

    //优先选择早点结束的，将原问题转化为，在一条路上尽可能铺更多的格子，那么我们只要保证每个格子的end尽可能早点结束，好让下一个格子进来
    start := intervals[0]
    res := 0
    for i:=1;i<len(intervals);i++{
        if intervals[i][0] >= start[1] {
            start = intervals[i]
        } else {
            res++
        }
    }

    return res
}
