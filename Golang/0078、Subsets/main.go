func subsets(nums []int) [][]int {
    //一样是递归回溯
    lens := len(nums)
    if lens == 0 {
        return [][]int{}
    }
    
    res := make([][]int,0)
    res = append(res, []int{})
    
    for i:=0; i < lens; i++ {
        recursion(i,0,0, lens,nums, []int{}, &res)
    }
    
    return res
}

func recursion(targetLength, start, level,lens int,nums []int, tmpRes []int, res *[][]int) {
    if level == targetLength+1 { 
        dist := make([]int, len(tmpRes))
        copy(dist, tmpRes)
        *res = append(*res, dist)
        return 
    }
    
    for i:= start; i <lens; i++ {
        recursion(targetLength, i+1,level+1, lens, nums, append(tmpRes, nums[i]), res)
    }
}
