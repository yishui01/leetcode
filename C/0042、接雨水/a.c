int trap(int* height, int heightSize) {
    //找出每个坐标的左边最大值和右边最大值，取其中较小的一个和当前坐标比较，如果比当前坐标大，差值就是雨水量
    if(heightSize == 0)return 0;
    int i = 0, res = 0, max = 0;
    int tmp[heightSize];
    for(i = 0; i < heightSize; i++) {
        tmp[i] = max;
        max = height[i] >max ? height[i] : max; //找出当前左边最大值
    }
    
    max = 0;
    
    for(i = heightSize -1; i >= 0; i--) {
        tmp[i] = tmp[i] > max ? max : tmp[i]; //找出左右两边的较小的一端
        max = max > height[i] ? max : height[i]; //找出当前右边的最大值
        if(tmp[i] > height[i])res+=tmp[i] - height[i];
    }
    
    return res;
}


//下面的很难理解，暂时留在这里先
/*int trap(int* height, int heightSize) {
    
    思路：双指针法，左右两头各放一个指针，矮的一端朝中间移动，假设左边矮，移动左边的指针，移动之前判断这个最大左边和当前的这个左边谁高，
    如果最大的左边高，那么当前的这个坐标就可以装水，为什么，首先移动的是矮指针，就是说右边的高度当前是最高的，没有之一，比你左边的任何一个
    坐标都高，那么也包括你的最大左边，判断一个坐标能不能装水主要看这个坐标的左右两边是不是都有比他高的坐标，现在右边比你高了，然后如果这个
    最大左边也比当前的左边高，那左右两边都有比自己高的了，自然可以装水，装水量就是矮的那一个围墙高度减去当前坐标的高度，也就是最大左边减去当前
    左边的高度。
    如果移动的是右指针，也一样的
    一句话：到了一个坐标，找出这个坐标的左边最大值和右边最大值，两个都比当前坐标高的话，取其中小的一个减去当前坐标高度就是这个坐标的盛水量
    
    int left = 0, right = heightSize -1, res = 0, max_left = 0, max_right = 0;
    while(left < right) {
        if(height[left] < height[right]) {
            //如果左边的比右边的矮,移动左指针
            max_left = height[left] > max_left ? height[left] : max_left; //找出最大左边
  printf("左指针是%d,右指针是%d,max_left是%dheight[left]是%d左边小于右边%d\n", left,right,max_left,height[left],max_left - height[left]);            res += max_left - height[left++];
        } else {
            max_right = height[right] > max_right ? height[right] : max_right;
            printf("右边小于等于左边%d\n", max_right - height[right]);
            res += max_right - height[right--];
        }
    }
    return res;
}*/