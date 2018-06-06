int threeSumClosest(int* nums, int numsSize, int target) {
    //反正这种题先把数组排序,然后用双指针
    int i = 0, j = 0, z = 0, res = 0, curr = 0,left = 0, right = 0, mytar = 0, maxabs = 0;
    for(i = 0; i < numsSize - 1; i++) {
        for (j = 0; j < numsSize - i -1; j++) {
            if(nums[j] > nums[j+1]) {
                z = nums[j];
                nums[j] = nums[j+1];
                nums[j+1] = z;
            }
        }
    }
    
    res = nums[0] + nums[1] + nums[2]; //当前的和
    maxabs = abs(target-res); //当前的和与target相差的值
    
    for(i = 0; i < numsSize-2; i++) { //遍历到倒数第三个数时就可以了

        left = i+1;
        right = numsSize - 1;
        mytar = target - nums[i]; //要追求的目标, 如果left指向的值+right指向的值能=mtar最好，不能等于的话，取绝对值相差最小的那个
        while(left < right) {
            curr = nums[left] + nums[right];
            if (curr == mytar) {
                return curr + nums[i];
            }else{
                if(abs(target-(curr+nums[i])) < maxabs) {
                res = curr + nums[i];
                maxabs = abs(target-res);
                }
            } 
            
            if(curr > mytar) {
                right--;
            } else {
               left++; 
            }

            
        }
        
    }
    
    return res;
}