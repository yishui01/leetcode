//大佬的代码就是整洁+清晰 6666


//首先构建图映射
func buildGraph(n int, edges [][]int) [][]int {
    graph := make([][]int, n)
    for i := 0; i < n; i++ {
        graph[i] = make([]int, 0)
    }
    for _, e := range edges {
        graph[e[0]] = append(graph[e[0]], e[1])
        graph[e[1]] = append(graph[e[1]], e[0])
    }
    return graph
}

//构建出每个结点当前的深度
func computeDegrees(graph [][]int) []int {
    degrees := make([]int, len(graph))
    for i, adj := range graph {
        degrees[i] = len(adj)
    }
    return degrees
}

//从最外层的叶子结点开始剥洋葱，剥完一轮找出深度为1的结点（新的叶子结点），进入下一轮剥洋葱【滑稽】
func findMinHeightTrees(n int, edges [][]int) []int {
    graph := buildGraph(n, edges) //构建一个图
    degrees := computeDegrees(graph)
    ans := make([]int, 0)
    for i, d := range degrees {
        if d == 1 {
            ans = append(ans, i) //找出深度为1的结点，就是最外层的结点入队
        }
    }
    
    for n > 2 {
        n -= len(ans) //减掉叶子结点数量
        next := make([]int, 0)
        for _, i := range ans {
            //degrees[i]-- //成功发发现大佬代码冗余【滑稽】【滑稽】
            for _, j := range graph[i] {
                degrees[j]-- //把与这个点连接的点的深度也减1
                if degrees[j] == 1 { //筛选出下一轮叶子结点
                    next = append(next, j)
                }
            }
        }
        ans = next
    }
    if n > 0 && len(ans) == 0 {
        ans = append(ans, 0)
    }
    return ans
}
