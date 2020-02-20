/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

/**
 * Encodes a tree to a single string.
 *
 * @param {TreeNode} root
 * @return {string}
 */
var serialize = function(root) {
    if (!root){
        return "#"
    }
    var str = root.val
    str += ","+serialize(root.left)
    str += ","+serialize(root.right)
    return str
};

/**
 * Decodes your encoded data to tree.
 *
 * @param {string} data
 * @return {TreeNode}
 */
var deserialize = function(data) {
    var arr = data.split(",")
    var index = 0
    var lens = arr.length

    function dfs(arr,length){
        if (index >= length){
            return null
        }
        let tmp = arr[index]
        index++
        if (tmp == "#"){
            return null
        }
        var node = new TreeNode(tmp)
        node.left = dfs(arr,length)
        node.right = dfs(arr,length)
        return node
    }

    return dfs(arr,lens)

};


/**
 * Your functions will be called as such:
 * deserialize(serialize(root));
 */
