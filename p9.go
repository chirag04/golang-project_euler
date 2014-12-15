package euler

/*
Special Pythagorean triplet

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

func P9() int {
	product := 0;
    for a := 1; a < 1000; a++ {
    	for b := a+1; b < 1000; b++ {
    		c := 1000 - a - b;
    		if(a*a + b*b == c*c) {
    			product = a * b * c;
    			break;
    		}
    	}
    	if(product != 0) {
    		break;
    	}
    }
    return product;
}