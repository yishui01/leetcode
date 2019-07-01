/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    //这题有什么难的啊，直接硬写啊，移呗
    //好吧这题我提交了4次才过，打脸了 ╮(╯▽╰)╭
    if head == nil {
        return nil
    }
    var res *ListNode = nil
    var tmp *ListNode = nil
    tmpVal := head.Val - 1 //重复值的赋值很有讲究，这个不好初始化，只能先赋值为头部节点值-1，保证和头部节点不一样就行，等找到第一个重复node之后再赋值
    init := false //是否已经初始化好了tmpVal,没初始化好的话，tmpVal只能保证和当前节点值不一样，这样才能不影响判断
    
    for head != nil {
        if init == false {
            tmpVal = head.Val - 1
        }
        if (head.Next == nil && head.Val != tmpVal) || (head.Next != nil && head.Val != head.Next.Val && head.Val != tmpVal) {
            //有效
            if res == nil {
                res = head
                tmp = head
            } else {
                tmp.Next = head
                tmp = tmp.Next
            }
            
        } else {
            tmpVal = head.Val //找到一个重复的Node，就把重复值赋值给tmpVal，加强防御，否则仅凭head.Val != head.Next.Val还不行，因为最后一个重复Node就会漏掉
            init = true
            if tmp != nil {
                tmp.Next = nil
            }
        }
        head = head.Next
    }
    
    return res
    
}
