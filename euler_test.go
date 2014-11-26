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