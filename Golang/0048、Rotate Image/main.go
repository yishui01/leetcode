func rotate(matrix [][]int)  {
    //确定四条边，每条边对应不同的映射
    //i代表行，j代表列
    //左边 -j为0是左边
    //上边 -i为0是上边
    //右边 -j为Colens-1是右边
    //下边 -i为Rowlens-1为下边
    
    //替换是从最外圈开始，一圈一圈朝内替换
    lens := len(matrix)
    if lens == 0 {
        return
    }
    //i代表圈数
    for i:=0; i < lens-1;i++ {
        //j代表里面要进行替换的轮数
        for j:=0; j < lens-1-i*2; j++ {
            //上换右
            //最右边的点保存下来
            tmp := matrix[i+j][lens-1-i]
            //最左边的点替换最右边的点
            matrix[i+j][lens-1-i] = matrix[i][i+j]
            
            //右换下
            tmp2 := matrix[lens-1-i][lens-1-i-j]
            matrix[lens-1-i][lens-1-i-j] = tmp
            
            //下换左
            tmp3 := matrix[lens-1-i-j][i]
            matrix[lens-1-i-j][i] = tmp2
            
            //左换上
            matrix[i][i+j] = tmp3
        }
       
    }
    
    
}
