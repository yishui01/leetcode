func canCross(stones []int) bool {
    maps := make(map[string]bool)
    return helper(0,1,stones,maps)
}


//记忆化dfs，[index和step拼接字符串]组成key来存储结果
func helper(startIndex int,step int,stones []int,maps map[string]bool) bool {
    if startIndex == len(stones)-1 {
        return true
    }

    if val,ok := maps[buildStr(startIndex,step)];ok {
        return val
    }

    //找当前跳完有没有落点
    r := stones[startIndex] + step
    newStart := -1
    for i:=startIndex+1;i<len(stones);i++{
        if stones[i] == r{
            newStart = i
            break
        }
    }
    if newStart == -1 {
        return false
    }

    //到这里就是说当前跳法是合格的
    //总共有三种法
    if step > 1 {
        if helper(newStart,step-1,stones,maps) {
            return true
        }
        maps[buildStr(newStart,step-1)] = false
    }

    if helper(newStart,step,stones,maps){
        return true
    }
    maps[buildStr(newStart,step)] = false

    if helper(newStart,step+1,stones,maps){
        return true
    }
    maps[buildStr(newStart,step+1)] = false

    return false
}

func buildStr(a,b int)string{
    return strconv.Itoa(a)+strconv.Itoa(b)
}
