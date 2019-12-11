func maxEnvelopes(envelopes [][]int) int {
    //先排好序，全部按照升序排列，第一个数字相等按照第二个元素升序
    //排完之后双重循环，找出前面元素的两个数字都比当前数字小的，dp
    


     /****
     解法一、全部升序排列，O(n2)迭代DP....耗时280ms
    sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
    res := 0
    dp := make([]int,len(envelopes))
    
    for i:=0;i<len(envelopes);i++{
        dp[i] = 1
        for j:=0;j<i;j++{
            if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1]{
                dp[i] = max(dp[i],dp[j]+1)
            }
        }
        res = max(res,dp[i])
    }

    return res
    ***/

    /***
    看了下大佬还有个标准LIS解法。精髓在于第一个数组升序，第二个数字降序，然后对所有第二个数字组成的一维数组进行LIS
    **/
    
    //解法二、第一个数字升序，第二个数字降序，降序是为了防止干扰对其进行LIS，只要是递增的，第一个数字肯定不相同，因为第一个数字相同的那些，第二个数字都是降序
    
     sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

    tails := make([]int,0) //注意此处不能命名为LIS，因为该数组并不是准确的LIS，只是一个保持递增的辅助数组而已

    for i:=0;i<len(envelopes);i++{
        number := envelopes[i][1] //第二个数字
        if i == 0 || tails[len(tails)-1] < number{
            tails = append(tails,number)
        } else {
            //来了个小数，这个小数的作用就是要找到他合适的位置，将原本在那个位置上比他大的数替换掉，最终的目的就是要让该数组整体尽可能的小
            j := len(tails)-1
            for ;j>=0;j--{
                if tails[j] < number {
                    break;
                }
            }
            tails[j+1] = number
        }
    }
    
    return len(tails)
}


func max (a,b int)int{
    if a > b {
        return a
    }
    return b
}
