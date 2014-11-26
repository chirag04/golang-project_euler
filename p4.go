package euler

/*
Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

func P4() int {
	maxPalindrome := 0;
    for i := 100; i < 1000; i++ {
    	for j := 100; j < 1000; j++ {
    		if k := i * j;  isPalindrome(k) && k > maxPalindrome {
    			maxPalindrome = k;
    		}
    	}
    }
    return maxPalindrome;
}

func isPalindrome(n int) bool {
	reversed := 0;
	i := n;
	for ; i > 0; i = i/10 {
		reversed = 10 * reversed + i % 10
	}
	return n == reversed;
}