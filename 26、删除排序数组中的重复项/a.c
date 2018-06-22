int removeDuplicates(int* nums, int numsSize) {
    if (numsSize <= 1)return numsSize;
    int i = 0, current = 0, place = 0;
    for (i = 0; i < numsSize; i++) {
        if(i == 0) {
            current = nums[i];
            place++;
        } else {
            if (nums[i] == current) {
                //如果是重复的
                continue;
            } else {
                //不是重复的
                current = nums[i];
                nums[place] = current;
                place++;
            }
        }
    }
    
    return place;
}