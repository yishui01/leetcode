func minMutation(start string, end string, bank []string) int {
    //直接用bfs，双向bfs（单向bfs面临到后面基数很大的问题），从src队列向dist转，再另一个队列向src转，每次遍历小的那个队列
    src,dist := make(map[string]struct{}),make(map[string]struct{})

    src[start],dist[end] = struct{}{},struct{}{}

    res := 0
    maps := []byte{'A','C','G','T'}

    bankMap := make(map[string]struct{})
    for _,v := range bank{
        bankMap[v] = struct{}{}
    }

    if start == end {
        return 0
    }

    if _,ok:=bankMap[end];!ok{ //目标基因不合法，直接-1
        return -1
    }

    for len(src) != 0 && len(dist) != 0{
        //选少的那个队列
        if len(src) > len(dist) {
            src,dist = dist, src
        }
        res++
        newSrc := make(map[string]struct{})
        for k,_ := range src {
            delete(src,k)
            //把k的每个字符都转换一下，如果再dist中就直接return，不在的话看是否合法，合法就继续入队
            bys := []byte(k)
            for i:=0;i<len(bys);i++{
                changeByte := bys[i]
                for _,mv := range maps {
                    if changeByte == mv {
                        continue
                    }
                    bys[i] = mv
                    newstr := string(bys)
                    //先看是否合法
                    if _,ok := bankMap[newstr];ok {
                        //再看是否达到目标
                        if _,ok := dist[newstr];ok {
                            return res
                        }
                         newSrc[newstr] = struct{}{}
                    }
                   
                }
                bys[i] = changeByte
            }

        }

        src = newSrc
    }

    return -1
}
