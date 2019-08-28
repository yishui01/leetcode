type ListNode struct {
     Key int
     Val int
 //    cacheNode *LRUCache
     Next *ListNode
     Prev *ListNode
}
type LRUCache struct {
    Head *ListNode
    Tail *ListNode
    Map map[int]*ListNode
//    Node *ListNode
    Capacity int
}


func Constructor(capacity int) LRUCache {
    cache := LRUCache {
        Head : nil,
        Tail : nil,
        Capacity: capacity,
    }
    cache.Map = make(map[int]*ListNode)
    return cache
}

func (this *LRUCache) MoveToHead (node *ListNode){
    if node == this.Head {
        return
    }
    if node.Next == nil {
//        fmt.Println("before moving tail ", this.Tail,"node", node, "node.prev", node.Prev)
        this.Tail = node.Prev
        this.Tail.Next = nil
    } else{
        node.Next.Prev = node.Prev
    }
    if node.Prev != nil {
        node.Prev.Next = node.Next
    }    
    node.Prev = nil
    node.Next = this.Head
    this.Head.Prev = node
    this.Head = node
    
//    fmt.Println("after move: head", this.Head, "tail ", this.Tail )        

}

func (this *LRUCache) Get(key int) int {
//    fmt.Println("get ",key)
    node,ok := this.Map[key]
    if !ok {
//        fmt.Println("not found")
        return -1
    }
    // move to head of queue
//    fmt.Println("moving to head of queue", node)

    this.MoveToHead(node)        
    
    return node.Val
}


func (this *LRUCache) Put(key int, value int)  {
//    fmt.Println("Put ", key, value, "len", len(this.Map), "cap", this.Capacity)
    
    node, ok := this.Map[key]
    if ok {
        node.Val = value
//        fmt.Println("update existing key", key, value)
        this.MoveToHead(node)
    } else {
        if len(this.Map) == this.Capacity {
    //        fmt.Println("kickout", this.Tail.Val)

            delete(this.Map, this.Tail.Key) 
     //       fmt.Println("tail in kickout", this.Tail, this.Tail.Prev)

            if this.Tail.Prev != nil {
                this.Tail.Prev.Next = nil
            }
            this.Tail = this.Tail.Prev
        }
    //    fmt.Println("after kickout", "head", this.Head, "tail ", this.Tail)
        node = new(ListNode)
        node.Val = value
        node.Key = key
        node.Next = this.Head
        this.Map[key] = node
    //    fmt.Println("map", this.Map[key], this.Map[key].Val)

        if this.Head == nil{
            this.Head = node
            this.Tail = node
            fmt.Println("init tail", this.Tail)
        } else {
            this.Head.Prev = node
        }
    }    
    this.Head = node
    
//    fmt.Println("head", this.Head, "tail ", this.Tail)
}

