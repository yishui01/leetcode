
func numSquares(n int) int {
    //其实这题最适合的是bfs，是一道经典bfs问题，一个数可以拆成多个子数，每一个子数又可以拆，完美bfs
    //队列从第一轮只有n一个元素开始，由n演化出(n-1*1),(n-2*2),(n-3*3),(n-4*4)...把这些演化出来的数字组合成一个新的队列，用于下一轮循环
    //那么其实每一个数字都代表一条演化路径，当前分解的次数大家都是一样的，就是外层循环的次数（队列迭代的次数），那么一旦演化出一个可以被直接
    //开方的数字，那么演化结束，最短路径就是这个能直接开方的这条路径，count就是个数。
    //注意：队列迭代的时候要用个东西记录下当前数字是否之前加入过队列，加入过就不要再加入队列了
    //ps：根据4平方和定理，每个正整数均可表为四个整数的平方和(其中有些整数可以为零)。
    
    if n == 0 {
		return 0
	}
	var count int
   // seen := make(map[int]bool) //64ms
    seen := make([]bool,n+1) //4ms  奇怪，换成切片快这么多的？？？
	q := []int{n}
    
	for len(q) > 0 {
		count++
		var newQ []int
        for k := 0; k < len(q); k++ {
            x := q[k]
			for i := 1; i*i <= x; i++ {
				if x == i*i {
					return count
				}
				next := x - i*i
				if !seen[next] {
					newQ = append(newQ, next)
					seen[next] = true
				}
			}
		}
		q = newQ
	}
	return count
}


// func numSquares(n int) int {
//     //这题很经典，自顶向下，dfs 记忆递归，每次尝试减掉一个数的平方和，把减到的值传递给下一层递归，递归返回值与当前的res比较，取其中较小的那个
//     res := make([]int,n+1)
//     return dfs(n,res)
// }

// func dfs(n int,res []int) int {
//     if n <= 1 {
//         return n
//     }
//     if res[n] != 0 {
//         return res[n]
//     }
    
//     sq := int(math.Sqrt(float64(n)))
//     if sq * sq == n {
//         res [n] =1
//         return 1
//     }
    
//     min := n
//     for i:=1;i<=sq;i++{
//         min = Min(min,dfs(n-i*i,res)+1)
//     }
//      res[n] = min
//     return min
// }

// func Min(a,b int) int {
//     if a <= b {
//         return a
//     }
//     return b
// }
