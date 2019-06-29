func combine(n int, k int) [][]int {
    //C42 => cnk
    
    //经典递归回溯
    
    if n == 0 || k == 0 {
        return [][]int{}
    }
    
    nums :=  make([]int, 0)
    for i:=0; i < n; i++ {
        nums = append(nums,i+1)
    }
    
    res := make([][]int, 0)
    
    recursion(k,0,0,len(nums),[]int{},&res)
    
    return res
    
}

func recursion(k, level, start, lens int, tmpRes []int, res *[][]int) {
    if level == k {
        dist := make([]int, k)
        copy(dist, tmpRes)
        *res = append(*res, dist)
        return 
    }
    
    for i:= start; i < lens; i++ {
        recursion(k,level+1,i+1,lens, append(tmpRes, i+1),res)
    }
    
}
