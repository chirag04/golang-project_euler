package euler

import "math"

/*
Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

func P3(n int64) int64 {
	for {
		p := smallestFactor(n);
		if p < n {
			n /= p;
		} else {
			return n;
		}
	}
}

func smallestFactor(n int64) int64 {
	length := sqrt(n);
	for i := int64(2); i <= length; i++ {
		if n % i == 0 {
			return i;
		}
	}
	return n;
}

func sqrt(n int64) int64 {
	return int64(math.Sqrt(float64(n)));
}