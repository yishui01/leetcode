func reconstructQueue(people [][]int) [][]int {
    //身高降序，前面人数升序
    sort.Slice(people,func(i, j int)bool{
        return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
    })

    res := make([][]int,0)

    for _,v := range people {
        if v[1] >= len(res) {
            res = append(res,v)
        } else {
            pos := v[1]
            //最后面补入一个，任意值
			res = append(res, res[len(res)-1])
            //将p位置移动到p+1，即将p位置空出来
			copy(res[pos+1:], res[pos:])
            //将p插入
			res[pos] = v
        }
    }

    return res

}
