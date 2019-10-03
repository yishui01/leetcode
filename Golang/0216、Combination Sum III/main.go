func combinationSum3(k int, n int) [][]int {
    if k == 0 {
        return [][]int{}
    }
    
    res := make([][]int, 0)
    
    helper(k,n,0,1,0,[]int{},&res)
    
    return res
}

func helper(k,n,level,start,tmpSum int,tmpSlice []int, res *[][]int) {
    if k == level {
        if n == tmpSum {
            dist := make([]int,k)
            copy(dist,tmpSlice)
            *res = append(*res, dist)
        }
        return 
    }
    
    for i:=start;i<=9;i++{
        tmpSlice = append(tmpSlice,i)
        helper(k,n,level+1,i+1,tmpSum+i,tmpSlice,res)
        tmpSlice = tmpSlice[:len(tmpSlice)-1]
    }
    
}

