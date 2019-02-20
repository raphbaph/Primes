// Implementation of the Sieve of Eratosthenes for finding prime numbers
// See: https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes for more information
// by Raphael Spannocchi - rs@dubidu.io - Learning Go in 2019

package main

import fmt "fmt"

const scale = 400 // Find primes in [1..scale+1]

func main() {
	var start_set [scale]int // this is the array to sieve through, consisting of numbers 1 to scale
	primes := make([]int, 1) // this slice will hold the prime numbers found

	for i := 2; i < scale+2; i++ { // populate the array to be sieved starting at 2, finding prime 1 is for free
		start_set[i] = i
	}

	div := 2                // Start with 2 as the first divider to check for primes
	inner := 1              // inner counts iterations of inner for loop
	outer := 0              // outer for loop
	first := true           // first flags first non-divisble number found
	primes[0] = 1           // 1 is the lowest prime number :)
	var candidate int       // storing a prime candidate
	var prime_pos int       // at position prime_pos in the array to be sieved, no need to start at 0 for every iteration, get some efficiency here :)
	var found_prime = false //flag if a prime was found during an inner iteration

	for outer < scale {
		for inner < scale {
			if start_set[inner]%div != 0 && first {
				candidate = start_set[inner]
				primes = append(primes, start_set[inner])
				first = false
				inner++
				prime_pos = inner
				found_prime = true
				continue
			}
			if start_set[inner]%div == 0 { // divisible number, set to 0, eliminate from sieve
				start_set[inner] = 0
				inner++
				continue
			}
			if start_set[inner] == 0 { // already flagged as non-prime, continue
				continue
			}
			inner++
		}
		div = candidate       // next prime needs to be the divider now
		first = true          // reset first non-divisible number found flag
		inner = prime_pos + 1 // start search at position of next prime plus 1
		if found_prime {      // if a prime was found in the last inner loop, jump to that position in the outer for loop.
			outer = candidate
			found_prime = false // reset the flag
		} else {
			outer++
		}
	}

	// output what we found and how many iterations were necessary.
	fmt.Printf("Prime numbers: %v\n", primes)
	fmt.Printf("Iterations: %v", iter)
}
