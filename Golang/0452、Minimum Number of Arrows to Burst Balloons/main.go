func findMinArrowShots(points [][]int) int {
    //按照结束坐标排序，贪心计算箭的数量，如果下一个气球的开始坐标不在当前的结束坐标范围内，那就箭数量+1，同时扩张结束坐标
    if len(points) <= 1 {
        return len(points)
    }

    sort.Slice(points,func(i,j int) bool{
        return points[i][1] < points[j][1]
    })

    end := points[0][1]
    count := 1

    for k,v :=  range points {
        if k == 0 {
            continue
        }
        if v[0] > end {
            count++
            end = v[1]
        }
    }

    return count
    
}
