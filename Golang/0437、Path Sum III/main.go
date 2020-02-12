/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 //方法一、暴力dfs，以每个节点为根节点，找出所有包含该根节点的有效路径  执行用时： 32ms
func pathSum(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }
    //包含root的有效路径 + 左半部分所有有效数 + 右半部分所有有效数
    return helper(root,sum)+pathSum(root.Left,sum)+pathSum(root.Right,sum)

    //错误：这里3个helper，代表包含root的有效路径和包含root的两个子节点的有效路径，然后root的子节点的子节点就漏掉了.....
    //return helper(root,sum)+helper(root.Left,sum)+helper(root.Right,sum)
}

//找出以root为根节点的有效路径是多少条
func helper(root *TreeNode,sum int)int {
    if root == nil{
        return 0
    }
    now := 0
    if root.Val == sum {
        now = 1 //这里不能直接返回1，因为有可能后面的子节点加起来一正一负抵消再加上父节点刚好也等于sum
    }
    return now + helper(root.Left,sum-root.Val)+helper(root.Right,sum-root.Val)
}

//方法二：map记录当前单条，注意是单条路径上所有节点值以及出现的次数，因为每次dfs完之后再回溯清空map，所以是单路径
 // 另外传一个当前单路径的节点和，并把每次产生的节点和存入map，出现相同的节点和就计数+1，
 // 到当前路径为止的节点和（总节点和）- sum = 我们需要找的节点和
 // 需要的节点和有几个，res就加几

 // 假设A-B-C-D-E-F 当前到F，总结点和 - sum的值在map中出现次数为2，一次是A-C的节点和，一次是A-E的节点和
 //那么合法路径就有两条，一次是 C（不包含C）-F ，一次是E（不包含E）-F，因为这两段距离肯定都是sum
 // 为毛不同节点数之和会一样？因为节点值可能有负数，所以会前后抵消的可能

//  执行用时： 4ms
/*func pathSum(root *TreeNode, sum int) int {
    hash := make(map[int]int)
    cur := 0
    count := 0
    hash[0] = 1
    dfs(root,cur,sum,hash,&count)
    return count
}

func dfs(root *TreeNode,cur int,sum int,hash map[int]int,count *int){
    if root == nil{
        return
    }
    cur += root.Val
    v,res := hash[cur - sum]
    if res {
        *count = *count + v
    }
    //我草下面这段也太猛了老铁，hash中记录的始终是当前这一条路径的节点上的值
    //先添加进去，最后回溯，这种方法怎么想到的，还有这种操作？
    hash[cur]++
    dfs(root.Left,cur,sum,hash,count)
    dfs(root.Right,cur,sum,hash,count)
    hash[cur] -= 1
}*/
