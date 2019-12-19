/* The guess API is defined in the parent class GuessGame.
   @param num, your guess
   @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
      int guess(int num); */

public class Solution extends GuessGame {

    public int guessNumber(int n) {
        int start = 0;
        int end = n;
        while(start <= end){
            int mid = start + (end-start)/2;
            //int mid = (start+end)/2; //这种写法在start和end的和足够大时导致整形溢出，造成数据混乱，GG
            int res = guess(mid);
            if (res == 0) {
                return mid;
            }
            if (res == -1) { //My number is lower说的是他自己的数字小一些，而不是我猜的这个数字小，-_-||
                end = mid-1;
            } else {
                start = mid+1;
            }
        }
        return n;
    }

}

