func permute(nums []int) [][]int {
    return per(nums, 0, len(nums), nil)
}

func per(nums []int, l,r int, res [][]int) [][]int {
    if l == r {
        return append(res, dup(nums))
    }
    for i := l; i < r; i++ {
        if i!=l{
            swap(l, i, nums)
        }
        res = per(nums, l+1, r, res)
        if i!=l{
            swap(l, i, nums)
        }
    }
    return res
}

func swap(a,b int, ar []int) {
    ar[a], ar[b] = ar[b], ar[a]
}

func dup(nums []int) []int {
    cp := make([]int, len(nums))
    copy(cp, nums)
    return cp
}
