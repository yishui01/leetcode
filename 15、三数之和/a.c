/**
 * Return an array of arrays of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
int** threeSum(int* nums, int numsSize, int* returnSize) {
    /*这道题看完毫无头绪，借鉴了网上的解题思路，先对数组进行排序，然后确定一个值，接下来对后面的数采用双指针，两者加起来=确定的值的相反数即可*/
    /*数组为从小到大排序，双指针一左一右只想确定的值之后的序列的两个端点，值大于target时，右指针左移，小于target时左指针右移，还要进行去重复操作*/
    int i = 0, j = 0, z = 0, left = 0, right = 0, target = 0, count = 0;

    int ** res = (int **)malloc(sizeof(int *) * numsSize*10); 
    int *n;
    //先把数组进行从小到大的排序
    for(i = 0; i < numsSize - 1; i++) {
        for( j = 0; j < numsSize - 1 - i; j++ ) {
            if (nums[j] > nums[j+1]) {
                z = nums[j];
                nums[j] = nums[j+1];
                nums[j+1] = z;
            }
        }
    }
    for(i = 0; i<numsSize-2; i++) {
        //遍历到倒数第三个的时候就停止
        if(nums[i] > 0)break;
        left = i + 1;
        right = numsSize - 1;
        target = 0 - nums[i];
        
        while(left < right) {
             if ( (nums[left] + nums[right]) > target ) {
               //  printf("到第一\n");
                right--;
            } else if ((nums[left] + nums[right]) < target ) {
                 // printf("到第二\n");
                left++;
            } else {
                 // printf("到第三\n");
                 n = (int *)malloc(sizeof(int) * 3);
                 n[0] = nums[i]; 
                 n[1] = nums[left];
                 n[2] = nums[right];
                 res[count] = n;
                 
                 count++;
                 left++;
                 right--;
                 
                 while(nums[left] == nums[left-1] && left < right) {
                     left++;
                 }
                 while(nums[right] == nums[right+1] && left < right) {
                     right--;
                 }
                 
            }
        }
        
        while(nums[i] == nums[i+1]) {
            //printf(去重复\n");
            i++;
        }
       
    }
    
    *returnSize = count;
    return res;
    
    
}