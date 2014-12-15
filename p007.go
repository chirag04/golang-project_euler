package euler

/*
10001st prime

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

func P7() int64 {
    var count, prime int64 = 0, 0;
    for i := int64(2); ; i++ {
    	if IsPrime(i) {
    		count++;
    	}
    	if(count == 10001) {
    		prime = i;
    		break;
    	}
    }
    return prime;
}