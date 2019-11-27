func isValidSerialization(preorder string) bool {
    //利用满二叉树性质，叶子节点的总和是除了叶子结点的其他节点的总和多1
    //题目中这棵树就是满二叉树，那么只要用两个变量记录叶子节点和非叶子节点即可

    leaf := 0
    node := 0

    arrs := strings.Split(preorder,",")

    for i:=0;i<len(arrs);i++{
        if arrs[i] == "," { 
            continue
        }

        if arrs[i] == "#" {
            leaf++
        } else {
            node++
        }
        if leaf > node+1 || (leaf == node+1 && i != len(arrs)-1) {
            return false
        }

    }

    return leaf == node+1

}
