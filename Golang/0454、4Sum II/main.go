func fourSumCount(A []int, B []int, C []int, D []int) int {
    //分成两半，A+B == -（C+D）
    res:= 0
    maps := make(map[int]int,len(A)+len(B))
    sort.Ints(A)
    sort.Ints(B)
    for i:=0;i<len(A);i++{
        l1 := i+1
        for ;l1<len(A)&&A[l1]==A[i];l1++{}
        for j:=0;j<len(D);j++{
            j1 := j+1
            for ;j1<len(B)&&B[j1] == B[j];j1++{}
            maps[A[i]+B[j]] += (l1-i) * (j1-j)
            j=j1-1
        }
        i=l1-1
    }

    for c:=0;c<len(C);c++{
        for d:=0;d<len(D);d++{
            if val,ok := maps[-(C[c]+D[d])];ok {
                res += val
            }
        }
    }

    return res
}
