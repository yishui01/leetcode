func findKthNumber(n int, k int) int {
    //字典序，前缀i开头的数字必定在i+1开头的数字之前
    //定义一个函数，确定两个前缀之间间隔的数字有多少,[1,2)之间的数字，比如1,10,2,3.. 这里1和2之间间隔了2个数字
    getCount := func (prefix,n int)int { //prefix可以是任意1-n之间的数字，
        count := 0
        a,b:=prefix,prefix+1
        for a<=n {
            count += min(n+1,b)-a  //n要+1是因为a是10的次幂扩大的，假设n为10 && 小于b，那么此时a已经是10了，n-a=0 不对，要把10这个数字算上，所以要+1，其他同理
            //下面的扩大十倍是前缀里下一个小区间的数字的个数，比如1和2,第一轮 count += 2-1,第二轮 count += 20-10，这样一直加下去，直到a>=n
            b*=10
            a*=10
        }
        return count
    }

    //确定第k个数字,前缀从1开始
    prefix:=1
    cnt:=1 //当前第几个数字
    for cnt < k {
        count:=getCount(prefix,n)
        if cnt + count > k {
            //k在当前prefix到prefix+1的数字之间
            prefix *= 10 //缩小一个区间，比如一开始是1，现在变成10，一开始是10，现在变成100，无论怎样，prefix和prefix*10在序列中都是前后相邻的数字
            cnt++ //由于prefix变化了，变成下一位asc序列数了，所以cnt要计数+1
        } else {
            prefix++ //这个每+1，count就成倍加，因为每一个prefix之间间隔了count个数
            cnt += count
        }
    }

    return prefix
}

func min(a,b int)int{
    if a <= b {
        return a
    }
    return b
}


