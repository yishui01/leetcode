func numberOfBoomerangs(points [][]int) int {
    //遍历每一个点，找出能和这个点组成回旋镖的所有点，等距离点为一组，计算每组能组成回旋镖的个数
    res := 0

    getDis := func(a,b []int)int{
        x := a[0]-b[0]
        y := a[1]-b[1]
        return x*x + y*y
    }

    for i:=0;i<len(points);i++{
        maps := make(map[int]int,len(points)) //提前分配好map大小，可以节约很多时间
        for j:= 0;j<len(points);j++{
            if i != j {
                 maps[getDis(points[i],points[j])]++
            }
        }
        for _,v := range maps {
            res += v*(v-1) //排列组合  A32  = 3*2  A42=4*3
        }
    }

    return res
}
