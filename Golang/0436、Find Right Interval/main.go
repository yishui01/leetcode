type Node struct {
    ints []int
    id int
}

type lists []*Node

func (self lists)Len()int{
    return len(self)
}
func (self lists)Less(i,j int)bool{
    return self[i].ints[0] < self[j].ints[0]    
}
func (self lists)Swap(i,j int){
    self[i],self[j] = self[j],self[i]
}
//题干有点坑
//你需要存储的满足条件的区间 j 的最小索引，这意味着区间 j 有最小的起始点可以使其成为“右侧”区间
//前半段是说最小索引，那这个最小索引是不是可以理解为intervals中最小的索引值？
//后半段又说j有最小起点值
//[ [3,4], [2,3], [1,2] ]  对于第二个元素[1,2],按照前半段题干，最小索引应该为0号位的[3,4], 按照后半段，是按最小起点值，那应该是2号位的[2,3]
//最终看了示例才知道是按照后半段最小起点值来决定，这样简单了很多

 //先做成Node，将ID绑上去，再从起点进行排序
 //排完序后二分
func findRightInterval(intervals [][]int) []int {
    lists := make(lists,len(intervals))
    for k,v := range intervals {
        lists[k] = &Node{
            ints:v,
            id:k,
        }
    }

    sort.Sort(lists)

    res := make([]int,len(intervals))

    for k,v := range intervals {
        res[k] = binarySearch(v[1],k,lists)
    }

    return res
}

func binarySearch(target int, index int,lists []*Node ) int {
    l,r := 0,len(lists)-1
    res := -1
    for l <= r {
        mid := l + (r-l)/2
        if lists[mid].ints[0] >= target{
            res = lists[mid].id
            r = mid-1 //不断往左靠近，寻找最靠近target的右边的起点值
        } else {
            l = mid+1
        }
    }
    return res    
}









