func merge(intervals [][]int) [][]int {
    //这能怎么办，两两合并咯
    //两层for循环合并，合并成功替换两个元素为合并后的数组
    res := make([][]int, 0)
    
    lens := len(intervals)
    if lens <= 1 {
        return intervals
    }
    
    for i:=0; i < lens; i++ {
        now := intervals[i]
        for j:=0;j < lens;j++ {
            if now[0] == intervals[j][0] && now[1] == intervals[j][1] {
                continue
            }
            //两两比较，合并成功时就将两者都替换成合并后的数组
            if nums, ok := goMerge(now, intervals[j]); ok {
            //合并成功，进行下一层合并
            intervals[j] = nums
            now = nums
            }
        }
    }
    
    for i:=0; i < lens; i++ {
        flag := 1
        for _,v := range res {
            if v[0] == intervals[i][0] && v[1] == intervals[i][1] {
                flag = 0
                break
            }
        }
        if flag == 1 {
            res = append(res, intervals[i])
        }
        
    }
    
    return res
}

//合并成功返回合并后的数组和true
//合并失败返回空数组和false
func goMerge(one, two []int) ([]int, bool) {
    if one[0] > two[1] || one[1] < two[0] {
        return []int{}, false
    }
    res := make([]int, len(one))
    if one[0] <= two[0] {
        res[0] = one[0]
    } else {
        res[0] = two[0]
    }
    
    if one[1] >= two[1] {
        res[1] = one[1]
    } else {
        res[1] = two[1]
    }
    
    return res, true
}
