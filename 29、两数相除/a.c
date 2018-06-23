#define ABS(x) ((x)>0)?(x):(-x)
int divide(int dividend, int divisor) {
    if(divisor==0)
        return INT_MAX;
    if(divisor==1)
        return dividend;
    if(divisor==-1)
    {
        if(dividend==INT_MIN)
            return INT_MAX;
        return -dividend;
    }
    long int absdibidend = ABS((long)dividend);
    long int absdibisor = ABS((long)divisor);
    long int ret=0;
    while(absdibidend >= absdibisor)
    {
        long int temp = absdibisor;
        long int count = 1;
        while(absdibidend > temp)
        {
            temp<<=1;
            if(temp<absdibidend)
                count<<=1;
        }
        ret+=count;
        absdibidend-=(temp>>1);
    }
    if(ret>INT_MAX || ret<INT_MIN)
        return INT_MAX;
    if((dividend<0 && divisor>0)||(dividend>0 && divisor<0))
        return -ret;
    else
        return ret;
}