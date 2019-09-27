func findOrder(numCourses int, prerequisites [][]int) []int {
    result := make([]int, 0, numCourses)
    
    graph  := make([][]int, numCourses)
    degree := make([]int, numCourses)
    
    for _, v := range prerequisites {
        src := v[1]
        dest := v[0]
        
        graph[src] = append(graph[src], dest)
        degree[dest] += 1 //记录了每一个有依赖项的课程
    }
    
    var queue []int
    //fill in the queue
    for i := 0; i < numCourses; i++ {
        if degree[i] == 0 {
            queue = append(queue, i)    //没有依赖项的入队
        }
	}
    
    for len(queue) > 0 {
        node := queue[0]
        result = append(result, node)
        
        for _, v := range graph[node] { //我服了，这里node原本是队列中的值，既然到了队列中说明已经全部解除依赖，然后这个点指向的所有节点都可以--，dist减到0
                                        //也就是依赖项减到0的时候，那就是解除全部依赖了就可以入队了，用于帮助其他人解除依赖，如果有环，那是不可能入队的
            degree[v] -= 1
            if degree[v] == 0 {
                queue = append(queue, v)
            }
        }
        
        queue = queue[1:]   // Dequeue
    }
    
    if len(result) == numCourses {
        return result    
    }
    
    return make([]int, 0)
}

//自己写的dfs判断是否有环
// func findOrder(numCourses int, prerequisites [][]int) []int {
    
//     res := make([]int, 0)
    
//     ready,find := make(map[int]bool),make(map[int]bool)
    
//      for i:=0;i<numCourses;i++{
//         if !helper(i,ready,find,&res,prerequisites){
//             return []int{}
//         }
//     }
    
//     return res
// }

// func helper(findNum int,ready map[int]bool,find map[int]bool, res *[]int, nums [][]int) bool {
//     if _,ok := ready[findNum];ok{
//         return true
//     }
//     if _,ok:=find[findNum];ok {
//         return false
//     }
    
//     find[findNum] = true
    
//     allDepend := make([]int,0)
//     for _,v := range nums{
//         if v[0] == findNum {
//             allDepend = append(allDepend, v[1])
//         }
//     }
    
//     for _,v := range allDepend{
//         if !helper(v,ready,find,res,nums){
//             return false
//         }
//     }
    
//     delete(find,findNum)
//     ready[findNum] = true
//     *res = append(*res,findNum)
    
//     return true
    
// }
