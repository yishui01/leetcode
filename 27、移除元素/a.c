int removeElement(int* nums, int numsSize, int val) {
    int  i = 0, place = 0;
    for(i = 0; i < numsSize; i++) {
        if(nums[i] != val){
            nums[place++] = nums[i];
        }
    }
    return place;
}