package euler

/*
Longest Collatz sequence

The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

func P14() uint64 {
    var starting_number, longest_chain uint64 = 1, 0;
    for i := uint64(0); i < 1e6; i++ {
    	count := chain_length(i);
    	if(count > longest_chain) {
    		longest_chain = count;
    		starting_number = i;
    	}
    }
    return starting_number;
}

func chain_length(n uint64) uint64 {
	var count uint64 = 1;
	for n > 1 {
		if n % 2 == 0 {
			n /= 2; 
		} else {
			n = 3*n + 1;
		}
		count++;
	}
	return count;
}