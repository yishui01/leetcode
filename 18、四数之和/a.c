/**
 * Return an array of arrays of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
int** fourSum(int* nums, int numsSize, int target, int* returnSize) {
    int i = 0, j = 0, z = 0, left = 0, right = 0, target1 = 0,  target2 = 0, target3 = 0, length = 0;
    int *current, **res;
    res = (int **)malloc(sizeof(int*)*numsSize*10);
    //和3sum思路一样
    //先从小到大排序
    for(i = 0; i < numsSize - 1; i++) {
        for(j = 0; j < numsSize-1-i; j++) {
            if(nums[j] > nums[j+1]) {
                z = nums[j];
                nums[j] = nums[j+1];
                nums[j+1] = z;
            }
        }
    }

    //再用双指针慢慢找，去重复
    for (i = 0; i < numsSize - 3; i++) {
        target1 = target - nums[i];
        
        for (j = i+1; j < numsSize - 2; j++) {
            target2 = target1 - nums[j];
            left = j+1;
            right = numsSize - 1;
            while(left < right) {
                target3 = nums[left] + nums[right];
                if (target3 > target2) {
                   
                    right--;
                } else if (target3 < target2) {
                  
                    left++;
                }else{
                   // printf("i是%d,target1是%d,target2是%d,left是%d，right是%d\n", i,target1, target2,left,right);
                    current = (int*)malloc(sizeof(int)*4);
                    current[0] = nums[i];
                    current[1] = nums[j];
                    current[2] = nums[left];
                    current[3] = nums[right];
                    res[length] = current;
                    length++;
                    //先去掉两指针的重复，如果下一个指针还和现在的值一样，那组合肯定也一样了
                    while(nums[left+1] == nums[left]) {
                        left++;
                    }
                    while(nums[right-1] == nums[right]) {
                        right--;
                    }
                    left++;
                    right--;
                }
            }
            //搞完一次循环再去掉target2重复的部分，如果下一个target2和现在的target2一样，那就是说前面的i对应的值，以及左右指针也肯定一样
            while(nums[j+1] == nums[j]) {
                j++;
            }
        }
           //搞完一次循环再去掉target1重复的部分，target1重复了那接下来的也一样，所以要去掉
        while(nums[i+1] == nums[i]) {
                i++;
            }
    }
    *returnSize = length;
    return res;
}