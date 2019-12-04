/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * public interface NestedInteger {
 *
 *     // @return true if this NestedInteger holds a single integer, rather than a nested list.
 *     public boolean isInteger();
 *
 *     // @return the single integer that this NestedInteger holds, if it holds a single integer
 *     // Return null if this NestedInteger holds a nested list
 *     public Integer getInteger();
 *
 *     // @return the nested list that this NestedInteger holds, if it holds a nested list
 *     // Return null if this NestedInteger holds a single integer
 *     public List<NestedInteger> getList();
 * }
 */
public class NestedIterator implements Iterator<Integer> {

    private List<Integer>list;
    private int index;

    public NestedIterator(List<NestedInteger> nestedList) {
        list = integerIterator(nestedList);
        index = -1;
    }

    private static List<Integer> integerIterator(List<NestedInteger> nestedIntegerList){
            ArrayList<Integer> list = new ArrayList<>(nestedIntegerList.size());
            for(NestedInteger tmp : nestedIntegerList){
                    if (tmp.isInteger()) 
                    list.add(tmp.getInteger());
                else 
                    list.addAll(integerIterator(tmp.getList()));
            }
            return list;
    }

    @Override
    public Integer next() {
        if (this.hasNext()){
            return list.get(++index);
        }
        return null;
    }

    @Override
    public boolean hasNext() {
        if(index+1 < list.size()){
            return true;
        }
        return false;
    }
}

/**
 * Your NestedIterator object will be instantiated and called as such:
 * NestedIterator i = new NestedIterator(nestedList);
 * while (i.hasNext()) v[f()] = i.next();
 */
