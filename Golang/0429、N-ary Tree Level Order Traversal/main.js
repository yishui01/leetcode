/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */
/**
 * @param {Node} root
 * @return {number[][]}
 */
var levelOrder = function(root) {
    var res = new Array();
    var f = (node,level)=>{
        if (!node){
            return 
        }
        if (!res[level]){
            res[level] = new Array();
        } 
         res[level].push(node.val)
        for(let i=0;i<node.children.length;i++){
            f(node.children[i],level+1)
        }
    }

    f(root,0)

    return res
};

