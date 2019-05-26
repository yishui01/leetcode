func combinationSum(candidates []int, target int) [][]int {
    //对于这种穷举题，我只说一句，回溯大法好!
    sort.Ints(candidates)
    lens := len(candidates)
    tmpSlice := make([]int, 0)
    res := make([][]int, 0)
    recursion(candidates, lens, 0,tmpSlice, 0, target, &res)
    return res
}

func recursion(nums []int, lens int, start int, tmpSlice []int, tmpRes int, target int, res *[][]int) {
   
    for i:=start; i < lens; i++ {
        if nums[i] > target {
            break
        }
        if nums[i] + tmpRes < target {
            t := make([]int, len(tmpSlice) + 1)
            copy(t, append(tmpSlice, nums[i]))
            recursion(nums, lens, i,t, tmpRes + nums[i], target, res)
        } else if nums[i] + tmpRes == target {
            *res = append(*res, append(tmpSlice, nums[i]))
        }
        
    }
}
