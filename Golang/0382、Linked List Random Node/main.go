type Solution struct {
	head *ListNode
}

/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	return this.poolAlgorithm(1)[0]
}

// 蓄水池算法，等概选择k个元素入蓄水池
func (this *Solution) poolAlgorithm(k int) []int {
	pool := make([]int, k) // 大小为k的蓄水池
	pointer := 0
	for cur := this.head;cur != nil;cur=cur.Next {
		random := rand.Intn(pointer + 1)
		if random < k { //当前随机数字打中蓄水池，那就替换蓄水池中的数字为当前链表val
			pool[random] = cur.Val
		}
		pointer ++
	}
	return pool
}

