int jump(int* nums, int numsSize) {
        int ret = 0; //当前步数
        int last = 0; //当前步数能跳到的最远距离的下标
        int curr = 0; //下一步能跳到的最远距离的下标
        for (int i = 0; i < numsSize; ++i) {
            if (i > last) {
                //遍历到的下标超过了当前步数的最远距离，那没办法，只能往前跳一步咯，既然跳了一步，就要更新下last
                last = curr;
                ++ret;
                if(last >= numsSize-1){
                    //当前步数的范围覆盖全部了就直接返回
                     return ret;
                 }
            }
            
            curr = curr > i+nums[i] ? curr : i+nums[i];
           
        }

        return ret;
}