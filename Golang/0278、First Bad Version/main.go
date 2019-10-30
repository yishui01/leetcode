bool isBadVersion(int version);


//好吧这题没有golang版本的什么鬼

int firstBadVersion(int n) {
    //这种直接二分了，要找的点它本身是true然后他前面必须为false，这个点就是最小的最差版本
    int res = 0;
    int left = 1;
    int right = n;
    int mid = 0;
    
    while(left <=right ){
        mid = left+(right-left)/2; //这一手学到了，（left+right）/2会导致在left+right的时候爆int
        if (isBadVersion(mid)) {
            if (mid == 0 || !isBadVersion(mid-1)) {
                res = mid;
                return res;
            }
            right = mid-1;
        } else {
            left = mid+1;
        }
    }
    
    return res;
}
