func wiggleMaxLength(nums []int) int {
    lens := len(nums)
    if lens < 2 {
        return lens
    }
    res := 1
    prev := -1
    curr := 0

    for i:=1;i<lens;i++{
        tmp := nums[i] - nums[i-1]
        if tmp == 0 {
            continue
        }
        if tmp > 0 {
            curr = 1
        } else {
            curr = 0
        }
        if curr != prev || prev == -1{
            res += 1
        }
        prev = curr
    }

    return res
}
