func insert(intervals [][]int, newInterval []int) [][]int {
    //我想问一句，这很难吗？确定要合并的范围，合并后返回就行了，很难吗？
    startIndex := -1 //合并的起点
    endIndex := -1 //合并的终点
    
    lens := len(intervals)
    if lens == 0 {
        return [][]int{newInterval}
    }
    if len(newInterval) == 0 {
        return intervals
    }
    
    for i:=0; i < lens; i++ {
        if intervals[i][0] > newInterval[1] || intervals[i][1] < newInterval[0] {
            //无交集
        } else {
            //有交集
            if startIndex == -1 {
                startIndex = i
                endIndex = i
            }
            if endIndex < i {
                endIndex = i
            }
        }
    }
    
     res := make([][]int, 0)
    
    if startIndex == -1 {
        res = append(intervals, newInterval)
        sort.Slice(res, func (i, j int) bool {
            return res[i][0] < res[j][0]
         })
        return res
    }
    
    //合并startIndex到endIndex的元素（包括newInterval本身也参与合并）
    flag := 0
    for i:=0; i < lens; i++ {
        if i<startIndex || i>endIndex {
            res = append(res, intervals[i])
        } else {
            if flag == 0 {
                start := -1
                end := -1
                if intervals[startIndex][0] <= newInterval[0] {
                    start = intervals[startIndex][0]
                } else {
                    start = newInterval[0]
                }
                
                if intervals[endIndex][1] >= newInterval[1] {
                    end = intervals[endIndex][1]
                } else {
                    end = newInterval[1]
                }
                
                tmp := []int{start, end}
                res = append(res, tmp)
                flag = 1
            }
        }        
    }
    
    return res
}
