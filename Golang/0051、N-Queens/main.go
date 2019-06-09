func solveNQueens(n int) [][]string {
//初始化一个一维数组用于判断
nums := make([]int, n)
for k, _ := range nums {
nums[k] = -1
}

res := make([][]string, 0)

recursion(nums, 0, n, &res)

return res
}


//看了大佬的检查方式使用一个一维数组，index为row，value为col，然后叼就叼在TMD可以用 横坐标和纵坐标的的绝对值是否相等来判断是否在一个斜线上，膜拜╮(╯▽╰)╭
func recursion(nums []int, level int, n int, res *[][]string) {
if level == n {
//有解法了，根据一维数组还原出需要的答案
t := make([]string, n)
for i := 0; i < n; i++ {
tmp := make([]byte, n)
for j := 0; j < n; j++ {
if nums[i] == j {
tmp[j] = 'Q'
} else {
tmp[j] = '.'
}
}
t[i] = string(tmp)
}

// fmt.Println(t,nums)
*res = append(*res, t)
return
}

for i := 0; i < n;i++ {
if check(level, i, nums, n) {
//合适的话就进入下一层
nums[level] = i
recursion(nums, level+1, n, res)
//回溯
nums[level] = -1
}
}

}

//检查当前位置是否合法
func check(row int, col int, nums []int, n int) bool {

for i := 0; i < len(nums); i++ {
//这里只要检查列和斜就行了，行不用检查，因为每一个元素就是一行，不存在冲突
if nums[i] == col {
return false
}

if nums[i] != -1 && (math.Abs(float64(col-nums[i])) == math.Abs(float64(row-i))){
return false
}
}
return true
}

//func solveNQueens(n int) [][]string {
//    //很明显，回溯穷举了解一下
//
//    nums := make([][]byte, n)
//    full := make([]byte,n)
//
//    for i:=0; i < n; i++ {
//        full[i] = '.'
//    }
//
//    for i:=0; i < n; i++ {
//        tmp := make([]byte, n)
//        copy(tmp, full)
//        nums[i] = tmp
//    }
//
//    res := make([][]string, 0)
//
//    recursion(nums,n,  0, &res)
//
//    return res
//}
//
//func recursion(nums [][]byte,n int,level int, res *[][]string) {
//    if n == level {
//        dist := make([]string, n)
//        for k,v := range nums {
//            dist[k] = string(v)
//        }
//        *res = append(*res, dist)
//        return
//    }
//
//    //先添加一个，看有没有问题，没有问题就进入到下一层
//    for i:=0; i < n; i++ {
//        nums[level][i] = 'Q'
//        if check(nums, n, level, i) {
//            recursion(nums, n, level+1, res)
//        }
//        nums[level][i] = '.' //回溯归位
//    }
//}
//
//
////检查皇后当前的位置是否合法
//func check(nums [][]byte, n int, row int, col int) bool {
//    //检查面积区块
//    if row-1 >=0 {
//        if col -1 >=0 && nums[row-1][col-1] == 'Q' {
//            return false
//        }
//        if nums[row-1][col] == 'Q' {
//            return false
//        }
//        if col + 1 < n && nums[row-1][col+1] == 'Q' {
//            return false
//        }
//    }
//
//    if col - 1 >= 0 && nums[row][col - 1] == 'Q' {
//            return false
//    }
//
//    if col + 1 < n && nums[row][col+1] == 'Q' {
//            return false
//    }
//
//    if row+1 < n {
//        if col -1 >=0 && nums[row+1][col-1] == 'Q' {
//            return false
//        }
//        if nums[row+1][col] == 'Q' {
//            return false
//        }
//        if col + 1 < n && nums[row+1][col+1] == 'Q' {
//            return false
//        }
//    }
//
//    //再检查横
//    for i:=0; i < n;i++ {
//        if i == col {
//            continue
//        }
//        if nums[row][i] == 'Q' {
//            return false
//        }
//    }
//     //再检查竖
//    for i:=0; i < n;i++ {
//        if i == row {
//            continue
//        }
//        if nums[i][col] == 'Q' {
//            return false
//        }
//    }
//    //再检查斜，他吗的斜着有两个方向，你到底要闹哪样，调了我好久，我TM真是蛋疼
//    for i:=1; i < n; i++ {
//        if row -i >=0 {
//            if col - i >= 0 && nums[row-i][col-i] == 'Q'{
//                return false
//            }
//            if col + i < n && nums[row-i][col + i] == 'Q'{
//                return false
//            }
//        }
//
//        if row+i <n {
//             if col - i >= 0 && nums[row+i][col-i] == 'Q'{
//                return false
//            }
//            if col + i < n && nums[row+i][col + i] == 'Q'{
//                return false
//            }
//        }
//
//    }
//
//    return true
//
//
//}
