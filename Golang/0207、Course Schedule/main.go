func canFinish(numCourses int, prerequisites [][]int) bool {
    if len(prerequisites) == 0 {
        return true
    }
    mapsEnd := make(map[int]bool)
    mapsIng := make(map[int]bool)
    for _,v := range prerequisites{
        if !dfs(numCourses,prerequisites,v[0],mapsEnd,mapsIng){
            return false
        }
    }
    return true
}

func dfs(count int, nums [][]int, nowFind int, mapsEnd,mapsIng map[int]bool) bool {
    if _,ok:=mapsEnd[nowFind];ok {
        return true
    }
    //如果当前要找的值没有,找出所有依赖项
    allDepend := make([]int, 0)
    for _,v:=range nums {
        if v[0] == nowFind{
            allDepend = append(allDepend, v[1])
        }
    }
    if len(allDepend) == 0 {
        mapsEnd[nowFind]=true
        return true
    }
    //有依赖项
    mapsIng[nowFind] = true
    
    for _,v := range allDepend {
        if _,ok := mapsIng[v];ok {
            return false
        }
        if _,ok := mapsEnd[v];!ok{
            if !dfs(count,nums,v,mapsEnd,mapsIng) {
                return false
            }
        }
    }
    
    //走到这里就是全部找到了
    delete(mapsIng,nowFind)
    mapsEnd[nowFind]=true
    
    return true
    
}



/***  下面大佬解法只要4ms
func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make(map[int][]int)
	for i := 0; i < len(prerequisites); i++ {
		c1, c2 := prerequisites[i][0], prerequisites[i][1]
		adj[c2] = append(adj[c2], c1)
	}

    visited := make(map[int]bool)
    recStack := make(map[int]bool)
	for c := range adj {
        if visited[c] {
            continue
        }
        if isCycle(c, adj, visited, recStack) {
            return false
        }
	}
	return true
}

func isCycle(c int, adj map[int][]int, visited map[int]bool, recStack map[int]bool) bool {
    visited[c] = true
    recStack[c] = true
    
    for _, nbr := range adj[c] {
        if recStack[nbr] {
            return true
        }
        if visited[nbr] == true {
            continue
        }
        if isCycle(nbr, adj, visited, recStack) {
            return true
        }
    }
    recStack[c] = false
    return false
}

***/

