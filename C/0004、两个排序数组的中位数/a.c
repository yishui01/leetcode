double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size) {
    //先分配一个大数组，把两个数组的值全部加到大数组中，再把大树组排序，去中位数即可
   
    int *arr = (int *) malloc(sizeof(int) * (nums1Size + nums2Size));
    double res;
    int i, j, z;
    for (i = 0; i < nums1Size; i++) {
        arr[i] = nums1[i];
    }
    for ( i = 0; i < nums2Size; i++) {
        arr[nums1Size + i] = nums2[i];
    }
    for ( i = 0; i < nums1Size + nums2Size; i++) {
        for (j = 0; j < nums1Size + nums2Size - 1 - i; j++) {
            if (arr[j] > arr[j+1]) {
                z = arr[j];
                arr[j] = arr[j+1]; 
                arr[j+1] = z;
            }
        }
    }
    for(i = 0; i < nums1Size + nums2Size; i++) {
        printf("%d\n", arr[i]);
    }
    i = (nums1Size + nums2Size) / 2;
    if((nums1Size + nums2Size) % 2 == 0) {
        //偶数个
        res = (arr[i-1] + arr[i]) /2.0;

    } else {
        //奇数个
         res = arr[i];
    }
    free(arr);
    return res;
}