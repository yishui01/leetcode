func isSelfCrossing(x []int) bool {
//这题不会做，看的答案，有三种情况，看着图对着图写代码即可
    /**
         x(1)
    ┌───┐
x(2)│   │x(0)
    └───┼──>
    x(3)│

x(1)
    ┌──────┐
    │      │x(0)
x(2)│      ^
    │      │x(4)
    └──────│
      x(3)

      x(1)
    ┌──────┐
    │      │x(0)
x(2)│     <│────│
    │       x(5)│x(4)
    └───────────│
        x(3)
**/
    for i:=3; i < len(x);i++ {
        if x[i] >= x[i-2] && x[i-3] >= x[i-1] {
            return true
        }
        if i >= 4 && x[i-1] == x[i-3] && x[i] >= x[i-2] - x[i-4] {
            return true
        }
        if (i >= 5 && x[i-2] >= x[i-4] && x[i-3] >= x[i-1] && x[i-1] >= x[i-3] - x[i-5] && x[i] >= x[i-2] - x[i-4]) {
                return true;
        }
    }
    return false
}