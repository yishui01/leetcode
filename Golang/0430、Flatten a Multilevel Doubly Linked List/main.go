/**
 * // Definition for a Node.
 * function Node(val,prev,next,child) {
 *    this.val = val;
 *    this.prev = prev;
 *    this.next = next;
 *    this.child = child;
 * };
 */
/**
 * @param {Node} head
 * @return {Node}
 */
var flatten = function(head) {
    var res = []
    var dfs = (node)=>{
        if (!node || node==undefined){
            return ;
        }
        res.push(node.val);
        if (node.child){
           dfs(node.child);
        }
        if (node){
          dfs(node.next);
        }
    }

    dfs(head);

    dummy = new Node(-1,null,null,null)
    tmpHead = dummy
    prev = null
    for (let i=0;i<res.length;i++){
        tmpHead.next = new Node(res[i],prev);
        tmpHead = tmpHead.next;
        prev = tmpHead;
    }

    return dummy.next;
};

