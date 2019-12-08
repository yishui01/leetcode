func topKFrequent(nums []int, k int) []int {
    //直接不要想太多，map计数后排序后取出前k个
    maps := make(map[int]int)

    for _,v := range nums {
        maps[v]++
    }

    ints := make([]int,0)

    for _,v := range maps {
        ints = append(ints,v)
    }
    sort.Ints(ints)
    min := ints[len(ints)-k]

    res := make([]int,0)
    for k,v := range maps {
        if v >= min {
            res = append(res,k)
        }
    }

    return res
}
