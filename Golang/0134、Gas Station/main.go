func canCompleteCircuit(gas []int, cost []int) int {
    lens := len(gas)
    now := 0
    start := 0
    //轮询每一个i作为起点，找到符合条件的i，O(n2)
    for i:=0; i < lens; i++ {
       start = i
        now = 0
        step := 0 
        for j:=0; j < lens; j++ {
            nowIndex := start+j
            if nowIndex > lens-1 {
                nowIndex -= lens
            }
            now += gas[nowIndex]
            if now  < cost[nowIndex] {
                break
            } else {
                now -= cost[nowIndex]
                step++
            }
        }
        if step == lens {
            return i
        }
    }
    
    return -1
}


/*
//看了大佬的解法，有个思路就是start点不行的时候，start+1肯定也不行，因为start到start+1肯定是有剩余的，至少不可能为负数，start都不行了，start+1，start+2一直到当前的i都是不行的， 那就直接将start设置为i+1，然后记录下到达i+1时欠下的债务，最后完成遍历时，看最后一步的剩下来的结余和之前需要偿还的债务是否能够抵消，能的话就ok不行就-1
时间复杂度：O(n)，完美
func canCompleteCircuit(gas []int, cost []int) int {
	remain, start, debt := 0, 0, 0
    len := len(gas)
    
	for i := 0; i <  len; i++ {
		remain += gas[i] - cost[i]
		if remain < 0 {
			start = i + 1
			debt += remain
			remain = 0
		}
	}
	if remain + debt >= 0  {
		return start
	}
	return -1
}

**/
