func combinationSum2(candidates []int, target int) [][]int {
    lens := len(candidates)
    sort.Ints(candidates)
    tmpSlice := make([]int, 0)
    res := make([][]int, 0)
    recursion(candidates, 0,0,lens, tmpSlice, 0, target, &res)
    
    return res
}

func recursion(nums[]int, start int, level int, lens int,  tmpSlice []int, tmpRes int, target int, res *[][]int) {
    // if level==start && start > 0 && nums[start] == nums[start -1] {
    //     return
    // }
    for i:=start; i < lens; i++ {
        if(i != start && nums[i] == nums[i-1]){ 
            //这里i != start 是精髓好吧，配合nums[i] == nums[i-1],直接去重
            //i != start 保证了当前的第一个元素是可以用的，因为我们传递到下一层的起点，也就是start，是下一个元素，
            //那么保证[1,1,2]这种两个1的能全部起作用，又要去重复
            //不为start，保证每个元素都能起作用，
            //nums[i] == nums[i-1] 保证起过作用的元素不会重复
            //就是说每个相同的元素只有在start的时候才能起作用，其他时间直接被过滤掉了，我每次传到下一层的都是不同数字，
            //然后下一层就靠着这个start ！= i 减小过滤条件才能拼成1,1,1这种组合，这里面每一个1都是靠着i == start才进来的
            continue; 
        }
        if tmpRes + nums[i] < target {
            t := make([]int, len(tmpSlice)+1)
            copy(t, append(tmpSlice, nums[i]))
            recursion(nums, i+1,level+1,lens, t, tmpRes+nums[i], target, res)
            
        } else if tmpRes + nums[i] == target {
            *res = append(*res,  append(tmpSlice, nums[i]))
        }
    }
}
