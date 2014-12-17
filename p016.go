package euler

/*
Power digit sum

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

func P16() int {
    sum := 0;
    // 2^1000 = 10 ^ x. Solve for x to get the number of digits.
    var num = [305]int{};
    num[0] = 1;
    for i := 1; i <= 1000; i++ {
    	for j := 0; j < len(num); j++ {
    		num[j] *= 2;
    		if j > 0 && num[j - 1] > 9 {
    			num[j] += int(num[j - 1] / 10);  
    			num[j - 1] %= 10;
    		}
    	}
    }
    for i := 0; i < len(num); i++ {
    	sum += num[i];
    }
    return sum;
}