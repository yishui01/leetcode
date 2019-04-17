func reverse(x int) int {
    
    const INT_MAX = int32(^uint32(0) >> 1)
    const INT_MIN = ^INT_MAX
    
    //一个个的截取，放到数组中，再把数组转换成数字
    flag := -1
    if x == 0 {
        return x
    } else if x < 0 {
        x = -x
        flag = 1 //代表之前的x为负数
    }
    if x < 10 {
        return x
    }
    numArr := make([]int, 0)
    for {
        if x >= 10 {
            tmp := x/10
            numArr = append(numArr, x-tmp*10)
            x /= 10
        } else if x > 0 && x < 10{
            numArr = append(numArr, x)
            x /= 10
        } else {
            break
        }
    }
    var res int
    numLen := len(numArr)
    for i:=0;i < numLen; i++ {
         res += numArr[i]*pow(10, numLen-1-i)
        if res > int(INT_MAX) {
            return 0
        }
    }
    
    return res * (flag*-1)
    
}

func pow(x,n int)int{
    pre := x
    if n == 0 {
        return 1
    }
    for ;n > 1; n--{
        x =x*pre
    }
    return x
}


/**
//看着大神的算法，给跪了
import "math"
// 문제를 끝까지 읽는 연습 
func reverse(x int) int {
    res := 0
    xInt := x
    for xInt != 0 {
        pop := xInt % 10
        xInt /= 10
        res = res*10 + pop
        if res > math.MaxInt32 {
            return 0
        }
        if res < math.MinInt32 {
            return 0
        }
    }
    return res
}
**/
