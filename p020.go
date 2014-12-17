package euler

/*
Factorial digit sum

n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

func P20() int {
    sum := 0;
    digits := [200]int{};
    digits[0] = 1;
    for i := 2; i <= 100; i++ {
    	for j := 0; j < len(digits); j++ {
    		digits[j] *= i;
    		if j > 0 && digits[j - 1] > 9 {
    			digits[j] += int(digits[j - 1] / 10);
    			digits[j - 1] %= 10;
    		}
    	}
    }
    for i := 0; i < len(digits); i++ {
    	sum += digits[i];
    }
    return sum;
}