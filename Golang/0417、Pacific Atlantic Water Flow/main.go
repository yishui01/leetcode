func pacificAtlantic(matrix [][]int) [][]int {
    if len(matrix) == 0 {
        return [][]int{}
    }
    //这种求全路径的，dfs没得跑了,每个点上下左右扩散
    res := make([][]int,0)
    for i:=0;i < len(matrix);i++{
        for j:=0;j<len(matrix[0]);j++{
            if leftTop(i,j,matrix) && rightBottom(i,j,matrix) {
                res = append(res,[]int{i,j})
            }
        }
    }

    //还有一种方法，从终点开始进行扩散，所到之处进行标记，最后找出左上和右下都能到达的点即可

    return res
}
var direction [][]int = [][]int{{1,0},{-1,0},{0,-1},{0,1}}

//是否能到达左上
func leftTop(i,j int,matrix [][]int) bool {
    if i ==0 || j == 0 {
        return true
    }
    tmp := matrix[i][j]
    matrix[i][j] = math.MinInt32
    for _, v := range direction {
        if i+v[0] >= 0 && j+v[1] >= 0 && i+v[0] < len(matrix) && j+v[1] <len(matrix[0]) && matrix[i+v[0]][j+v[1]] != math.MinInt32  && matrix[i+v[0]][j+v[1]]  <= tmp  && leftTop(i+v[0],j+v[1],matrix){
            matrix[i][j] = tmp
            return true
        }
    }
     matrix[i][j] = tmp

    return false
}

//是否能到达右下
func rightBottom(i,j int,matrix [][]int) bool {
     if i ==len(matrix)-1 || j == len(matrix[0])-1 {
        return true
    }
    tmp := matrix[i][j]
    matrix[i][j] = math.MinInt32
    for _, v := range direction {
        if i+v[0] >= 0 && j+v[1] >= 0 && i+v[0] < len(matrix) && j+v[1] <len(matrix[0]) && matrix[i+v[0]][j+v[1]] != math.MinInt32  && matrix[i+v[0]][j+v[1]]  <= tmp  && rightBottom(i+v[0],j+v[1],matrix){
            matrix[i][j] = tmp
            return true
        }
    }
    matrix[i][j] = tmp
    return false
}


