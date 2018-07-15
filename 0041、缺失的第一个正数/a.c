int firstMissingPositive(int* nums, int numsSize) {
    //借鉴了大神思路，把数组自己当做一个正的自然数序列hash表，也就是值是1,2,3,4,5...  对应的数组下标是0,1,2,3,4...
    //我们遍历数组的时候，把符合条件的值放到对应的下标当中（和原来的值做一个交换），如果位置上已经有一个符合条件的数了，那就不用交换
    int i = 0,j = 0, z = 0;
    for(i = 0; i < numsSize; i++) {
         //如果不是正数直接跳过，不符合序列要求没什么可交换的，
        //大于最大值也跳过，没有位置给你放，
        //如果位置上已经有个符合要求的数，那也不用交换了
            while(nums[i] > 0 && nums[i] <= numsSize && nums[nums[i]-1] != nums[i]){
                //之前是用的if判断，然而各种BUG，看了答案改成了while循环，我是有多chun
                z = nums[nums[i] -1 ];
                nums[nums[i] -1 ] = nums[i];
                nums[i] = z;
            }

    }
    
    //现在交换完了，就把这个交换完的数组当做是一个完美正自然数序列，遍历的时候发现哪个位置不对的话，直接返回那个位置该有的自然数，那个数就是最小的
    for(i = 0;i < numsSize; i++) {
        if(nums[i] != i+1)return i+1;
    }
    
    return i+1;
}