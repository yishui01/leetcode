import "sort"
func fourSum(nums []int, target int) [][]int {
    //先排序，再确定一个数字，再从这个数字的后面确定一个数字，再从后来确定的数字后面用用双指针
    lens := len(nums)
    if lens <= 3 {
        return [][]int{}
    }
    sort.Ints(nums)
    res := make([][]int, 0)
    for i:=0; i < lens-3;i++ {
        if (i > 0 && nums[i] == nums[i-1]){
            continue
        }
        //第一个数字为nums[i]
        for j:=i+1;j < lens-2;j++ {
            if (j > i+1 && nums[j] == nums[j-1]){
            continue
             }
            //第二个数字为nums[j]
            l:= j+1
            r:= lens-1
            for l < r {
                tmpSum := nums[i] + nums[j] + nums[l] + nums[r]
                if tmpSum == target {
                    res = append(res,[]int{nums[i], nums[j], nums[l],nums[r]})
                    for r > l && nums[r] == nums[r-1] {
                         r--
                    }
                    for r > l && nums[l] == nums[l+1] {
                         l++
                    }
                    r--
                    l++
                } else if tmpSum > target {
                    for r > l && nums[r] == nums[r-1] {
                         r--
                    }
                    r--
                } else {
                    for r > l && nums[l] == nums[l+1] {
                         l++
                    }
                    l++
                }
            }
        }
    }
    
    return res
}
