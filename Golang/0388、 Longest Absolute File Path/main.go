func lengthLongestPath(input string) int {
    res := 0
    maps := make(map[int]int)
    s := input
    //每一层都在上一层的基础上进行拼接
    level := 0
    i:=0
    flag := false
    for i < len(s) {
        flag=false
        if i < len(s) && s[i] == '\n' {
            level = 0
            i++
            for s[i] == '\t'{
                i++
                level++
            }
        } else {
            maps[level]=0
            for i < len(s) && s[i] != '\n'  {
                if s[i] == '.'{
                    flag = true
                }
                maps[level]++
                i++
            }

            if flag {
                     //将该阶段和之前阶段的总和加起来
                    tmp,index := level,level
                    for index >= 0 {
                        tmp += maps[index]
                        index--
                    }
                    //fmt.Println(level,maps,tmp)

                    if tmp > res {
                        res = tmp
                    }
            }
           
        }
    }

    if !flag {
        return 0
    }

    return res


}
