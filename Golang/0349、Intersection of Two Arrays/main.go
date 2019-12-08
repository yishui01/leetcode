func intersection(nums1 []int, nums2 []int) []int {
    //这....将其中一个数组建立一个map如何...
    res := make([]int,0)
    maps := make(map[int]bool)
    for _,v:=range nums1{
        maps[v]=true
    }
    for _,v := range nums2{
        if val,ok:=maps[v];ok && val{
            res = append(res,v)
            maps[v]=false
        }
    }

    return res
}
