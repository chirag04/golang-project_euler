package euler

/*
Smallest multiple

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

func P5() int {
	_lcm := 1;
    for i := 1; i <= 20; i++ {
    	_lcm = Lcm(i, _lcm);
    }
    return _lcm;
}

