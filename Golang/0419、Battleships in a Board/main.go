func countBattleships(board [][]byte) int {
    //follow up不会，看了题解，遇到X就无脑+1再判断上左是否是X，是的话再-1
    res := 0
    for i:=0;i<len(board);i++{
        for j:=0;j<len(board[0]);j++{
            if board[i][j] != 'X'{
                continue
            }
            res++
            if i > 0 && board[i-1][j] == 'X'{
                res--
            }
            //这里题目说了不会给出无效的用例，所以不用担心会有同时符合i和j的情况导致重复减少res，不会的
            if j > 0 && board[i][j-1] == 'X' {
                res--
            }
        }
    }

    return res
}
