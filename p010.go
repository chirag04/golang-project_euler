package euler

/*
Summation of primes

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

func P10() int64 {
	var i, sum int64 = 3, 2;
    for ; i < 2e6; i++ {
        if IsPrime(i) {
            sum += i;
        }
    }
    return sum;
}