/**
思路：
1.用slice存值，用map保存值在slice中的index；
2.每次删除时，为了避免移动元素，用数组末尾元素覆盖需要删除的元素，然后删除数组末尾元素；

*/

type RandomizedSet struct {
	vals []int
	_map map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
        rand.Seed(time.Now().Unix())
	return RandomizedSet{
		_map: make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {

	if _, ok := this._map[val]; ok {
		return false
	}
	this.vals = append(this.vals, val)
	this._map[val] = len(this.vals) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this._map[val]; !ok {
		return false
	}

	i := this._map[val]
	last := len(this.vals) - 1

	this.vals[i] = this.vals[last]
	this._map[this.vals[last]] = i

	delete(this._map, val)
	this.vals = this.vals[:last]

	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(this.vals))
	return this.vals[i]
}

