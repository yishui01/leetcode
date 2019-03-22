int jump(int* nums, int numsSize) {
        int ret = 0; //当前步数
        int last = 0; //当前步数能跳到的最远距离的下标
        int curr = 0; //当前i节点能跳到的最远距离的下标
        if(nums[0] == 0 || numsSize == 0 || numsSize == 1)return 0;
        for (int i = 0; i < numsSize; ++i) {
            if (i > last) {
                //遍历到的下标超过了当前步数的最远距离，那没办法，只能往前跳一步咯，既然跳了一步，就要更新下last
                last = curr;
                ++ret;
            }
            
            curr = curr > i+nums[i] ? curr : i+nums[i];
             if(curr >= numsSize-1){
                   return ret+1; 
                  /*下一步步数的范围覆盖全部了就直接返回
                 为什么可以在这里直接返回res+1，因为last记录着当前步数内的所有i节点，并且对应着步数，假设last
                 覆盖的范围内有能够直接跳到末尾的i节点存在，那么很遗憾，这个i节点在之前一定被curr提前发现了，为什么？last本来就是curr赋值而来，
                 last的范围就是上一个curr的范围，那么当curr第一次遍历到这个节点的时候，程序就直接返回了，不会再对last进行赋值，所以last的范围是不可能
                 包含可以跳到末尾的数组下标。last只记录着上一个curr的值，还记录着上一次的步数，此时curr发现了结果节点，说明再跳一步，就可以直接到数组
                 尾，那么步数就直接是last的步数+1，返回就行了*/
                 }
           
        }

        return ret;
}