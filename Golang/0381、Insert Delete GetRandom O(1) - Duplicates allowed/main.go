type RandomizedCollection struct {
    vals []int
    maps map[int][]int
}


/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
    rand.Seed(time.Now().Unix()) 
    return RandomizedCollection {
        vals:make([]int,0),
        maps:make(map[int][]int),
    }
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
    res := false
    this.vals = append(this.vals,val)
    if _,ok := this.maps[val];!ok {
        res = true
    }
    this.maps[val] = append(this.maps[val],len(this.vals)-1)

    return res
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
    res := false
    if _,ok := this.maps[val];ok {
        res =  true

        //先移除那个要移除的元素的val的值，以及在maps中保持的索引
        //把要移除的值和数组最后一个元素交换位置，移除数组最后一个元素即可
        indexArr := this.maps[val]
        removeIndex := indexArr[0]

        this.vals[removeIndex] = this.vals[len(this.vals)-1]
        this.vals = this.vals[:len(this.vals)-1]

        if len(indexArr) == 1 {
            delete(this.maps,val)
        } else {
            this.maps[val] = indexArr[1:]
        }

        //现在要移除的那个元素已经从vals中删除了，并且其对应的maps也清理了，但是那个辅助移除的数组最后一个元素，他的vals没问题，但他的索引变了
        //因此要调整其maps中的值
        
        if len(this.vals) > removeIndex { //要移除的那个元素可能本身就是最后一个元素，那么就上面已经在maps中移除了，不用接下来的替换了
        //接下来把最后一个元素对应的索引调整成removeIndex，因为别人本来是在最后的，你把他调到removeIndex了，那他对应的maps中保持的值也肯定要变
                for k,v := range this.maps[this.vals[removeIndex]] {
                    if v == len(this.vals){
                        this.maps[this.vals[removeIndex]][k] = removeIndex
                        break
                    }
                }
        }
    }

    return res
}


/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
    i:=rand.Intn(len(this.vals))
    return this.vals[i]
}


/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
