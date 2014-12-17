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

func SumDivisors(n int64) int64 {
    var sum int64 = 0;
    for i := int64(1); i < n; i++ {
        if n % i == 0 {
            sum += i;
        }
    }
    return sum;
}

func CountDivisors(n int64) int64 {
    var count, length int64 = 0, Sqrt(n);
    for i := int64(1); i < length; i++ {
        if n % i == 0 {
            count += 2;
        }
    }
    if(length * length == n) { //perfect square
        count++;
    }
    return count;
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