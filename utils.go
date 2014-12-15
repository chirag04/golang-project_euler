package euler

import (
    "math"
)

func IsPrime(n int64) bool {
    length := Sqrt(n);
    for i := int64(2); i <= length; i++ {
        if n % i == 0 {
            return false;
        }
    }
    return true;
}

func SmallestFactor(n int64) int64 {
	length := Sqrt(n);
	for i := int64(2); i <= length; i++ {
		if n % i == 0 {
			return i;
		}
	}
	return n;
}

func Sqrt(n int64) int64 {
    return int64(math.Sqrt(float64(n)));
}