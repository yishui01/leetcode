/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//迭代流
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	toLeft := false
	ret := [][]int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		values := []int{}
		levelLength := len(q)
		for i := levelLength - 1; i >= 0; i-- {
			values = append(values, q[i].Val)
			if toLeft {
				if q[i].Right != nil {
					q = append(q, q[i].Right)
				}
				if q[i].Left != nil {
					q = append(q, q[i].Left)
				}
			} else {
				if q[i].Left != nil {
					q = append(q, q[i].Left)
				}
				if q[i].Right != nil {
					q = append(q, q[i].Right)
				}
			}
		}
		ret = append(ret, values)
		q = q[levelLength:]
		toLeft = !toLeft
	}

	return ret
}





//以下为本人的渣渣思路，先使用递归全部按照从左往右的顺序排好，再处理结果数组，把该反转的反转
// func zigzagLevelOrder(root *TreeNode) [][]int {
//     res := make([][]int, 0)
//     recursion(root, 0, &res)
//     if res != nil {
//         for i:=0; i < len(res); i++ {
//         if i %2 == 0 {
//             continue
//         } else {
//             subLens := len((res)[i])
//             for j:=0; j < subLens/2; j++ {
//                 (res)[i][j], (res)[i][subLens-1-j] = (res)[i][subLens-1-j], (res)[i][j]
//             }
//           }
//         }
//     }
    
    
//     return res
// }

// func recursion(node *TreeNode, level int, res *[][]int){
//     if node == nil {
//         return 
//     }
    
//     if len(*res) < level+1 {
//         *res = append(*res, []int{node.Val})
//     } else {
//         (*res)[level] = append((*res)[level], node.Val)
//     }
    
//     if node.Left != nil {
//         recursion(node.Left, level + 1, res)
//     }
    
//     if node.Right != nil {
//         recursion(node.Right, level + 1 ,res)
//     }
    
// }
