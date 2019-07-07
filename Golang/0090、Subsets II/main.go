func subsetsWithDup(nums []int) [][]int {
    lens := len(nums)
    if lens == 0 {
        return [][]int{}
    }
    sort.Ints(nums)
    res := make([][]int, 0)
    res = append(res, []int{})
    for i := 0; i < lens; i++ {
        recusrsion(nums, []int{}, 0, i+1, 0, lens, &res)
    }
    
    return res
    
}

func recusrsion(nums,tmpRes []int, level,target, start, lens int, res *[][]int) {
    if level == target {
        dist := make([]int, len(tmpRes))
        copy(dist, tmpRes)
        *res = append(*res, dist)
        return 
    }
    
    for i:=start; i < lens; i++ { 
        //注意这个过滤条件是要在数组排序之后才可生效的，也就是重复的元素要相邻才行，不然没法过滤
        //i > start 可以放行那些起点值，起点值不管是不是重复，因为没有被使用过，都应该放行
        //nums[i] == nums[i-1] 这就是去重复的关键，当前值和之前的值是一样的，那就不必要进入下一层了，下一层已经收到你的心意了，不要传重复的心意给下一层
        if i > start && nums[i] == nums[i-1] {
            continue
        }
        recusrsion(nums, append(tmpRes, nums[i]), level+1, target, i+1, lens, res)
    }
    
}
