package euler

import "testing"

func TestP1(t *testing.T) {
	const out = 233168;
	if x := P1(); x != out {
		t.Errorf("Sum of Multiples of 3 and 5 upto 1000 = %v, want %v", x, out);
	}
}

func TestP2(t *testing.T) {
	const out = 4613732;
	if x := P2(); x != out {
		t.Errorf("Sum of even fibo numbers upto 4000000 = %v, want %v", x, out);
	}
}

func TestP3(t *testing.T) {
	const in, out = 600851475143, 6857;
	if x := P3(in); x != out {
		t.Errorf("Largest prime factor of %v = %v, want %v", in, x, out);
	}
}

func TestP4(t *testing.T) {
	const out = 906609;
	if x := P4(); x != out {
		t.Errorf(" largest palindrome made from the product of two 3-digit numbers = %v, want %v", x, out);
	}
}

func TestP6(t *testing.T) {
	const out = 25164150;
	if x := P6(); x != out {
		t.Errorf(" the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum = %v, want %v", x, out);
	}
}

func TestP7(t *testing.T) {
	const out = 104743;
	if x := P7(); x != out {
		t.Errorf(" the 10 001st prime number = %v, want %v", x, out);
	}
}

func TestP9(t *testing.T) {
	const out = 31875000;
	if x := P9(); x != out {
		t.Errorf(" the product abc = %v, want %v", x, out);
	}
}

func TestP10(t *testing.T) {
	const out = 142913828922;
	if x := P10(); x != out {
		t.Errorf(" the sum of all the primes below two million = %v, want %v", x, out);
	}
}

func TestP12(t *testing.T) {
	const out = 76576500;
	if x := P12(); x != out {
		t.Errorf(" the value of the first triangle number to have over five hundred divisors = %v, want %v", x, out);
	}
}

func TestP13(t *testing.T) {
	const out = 5537376230;
	if x := P13(); x != out {
		t.Errorf(" the first ten digits of the sum of the following one-hundred 50-digit numbers = %v, want %v", x, out);
	}
}

func TestP14(t *testing.T) {
	const out = 837799;
	if x := P14(); x != out {
		t.Errorf(" starting number, under one million, produces the longest chain = %v, want %v", x, out);
	}
}

func TestP20(t *testing.T) {
	const out = 648;
	if x := P20(); x != out {
		t.Errorf(" the sum of the digits in the number 100! = %v, want %v", x, out);
	}
}