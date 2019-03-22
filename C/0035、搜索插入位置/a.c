int searchInsert(int* nums, int numsSize, int target) {
    //for循环遍历啊，遍历数大于目标数字的时候停止，返回前一个下标，太简单了吧
    int i = 0;
    for(i = 0; i < numsSize; i++) {
        if(nums[i] == target) {
            return i;
        }
        if(nums[i] > target)return i;
    }
    
    return i;
}