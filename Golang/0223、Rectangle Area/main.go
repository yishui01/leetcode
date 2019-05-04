func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
    area := (C-A)*(D-B) + (G-E)*(H-F)
    //从宽的角度来看是否有交集
    //A的左边如果小于B的左边，那A的右边一定要大于B的左边
    //A的左边如果大于B的左边，那A的左边一定要小于B的右边
    
    //从高的角度来看是否有交集
    //A的上面如果高于B的上面，那A的下面一定要矮于B的上面
    //A的上面如果矮于B的上面，那A的上面一定要高于B的下面
    
    width := 0
    height := 0
    if A <= E && C >= E{
        if C>=G{
            //包含
            width = G-E
        } else {
             width = C - E
        }
    } else if A >= E &&  A <= G {
        //看是否是包含关系
        if C < G {
            //包含
             width = C - A
        } else {
             width = G - A
        }
    }
    
    if D >= H && B<=H{
        if F >= B {
            height = H-F
        } else {
            height = H-B
        }
    } else if D<=H && D>=F {
        if B >= F {
            height = D-B
        } else {
            height = D-F
        }
    }
    return area - width*height
    
}
