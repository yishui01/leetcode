func intersect(nums1 []int, nums2 []int) []int {
    //和前一题一样，多了个计数
    res := make([]int,0)
    maps := make(map[int]int)
    for _,v:=range nums1{
        maps[v]++
    }
    for _,v := range nums2{
        if val,ok:=maps[v];ok && val > 0{
            res = append(res,v)
            maps[v]--
        }
    }

    return res
}


/**
还有一个在排序状态下，两个数组单指针从0开始，相等则存入res，nums1的值大则nums2的指针++，反之则nums1的指针++，其中任何一个数组遍历完成则停止，输出res
**/
