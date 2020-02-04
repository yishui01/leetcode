/*
// Definition for a QuadTree node.
class Node {
    public $val;
    public $isLeaf;
    public $topLeft;
    public $topRight;
    public $bottomLeft;
    public $bottomRight;

    @param Boolean $val 
    @param Boolean $isLeaf 
    @param Node $topLeft 
    @param Node $topRight 
    @param Node $bottomLeft 
    @param Node $bottomRight 
    function __construct($val, $isLeaf, $topLeft, $topRight, $bottomLeft, $bottomRight) {
        $this->val = $val;
        $this->isLeaf = $isLeaf;
        $this->topLeft = $topLeft;
        $this->topRight = $topRight;
        $this->bottomLeft = $bottomLeft;
        $this->bottomRight = $bottomRight;
    }
}
*/
class Solution {

    /**
     * @param Integer[][] $grid
     * @return Node
     */
    function construct($grid) {
        if (count($grid) == 0) {
            return new Node(false,true,null,null,null,null);
        }
        return $this->helper($grid,0,0,count($grid));
    }

    public function helper($grid,$x,$y,$len)
     {
        if ($len <= 0){
            return null;
        }
        $sub = (int)($len/2);
        for ($i=$x;$i<$x+$len;$i++){
            for ($j=$y;$j<$y+$len;$j++){
                if ($grid[$i][$j] != $grid[$x][$y]){
                    return new Node(true,false,
                    $this->helper($grid,$x,$y,$sub),
                    $this->helper($grid,$x,$y+$sub,$sub),
                    $this->helper($grid,$x+$sub,$y,$sub),
                    $this->helper($grid,$x+$sub,$y+$sub,$sub));
                }
            }
        }

        return new Node($grid[$x][$y] == true,true,null,null,null,null);
    }

}
