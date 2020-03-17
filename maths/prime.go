/*
@Author: Piyus Gupta
@Date  : March 17, 2020
@Title : Quick way to get primes till a given number in golang

Future Scope :
We can do make use of channels to make this code concurrent to
enable parallization and have it as faster as possible.

[piyusg@piyusg-devbox]$ ./prime --help
Usage of ./prime:
  -debug
        a bool flag for debug
  -last_number int
        an int till which we want to check primes (default 100)
[piyusg@piyusg-devbox]$ time ./prime --last_number=100000000 --debug=false
Total primes below 100000000 is 5761455

real    1m42.410s
user    2m25.012s
sys     0m5.212s

It has been tested in a quad-core vm for first 100M numbers, took 1:42 min.
*/

package main

import (
	"flag"
	"fmt"
	"math"
)

var primes []int

func main() {
	max_check := flag.Int("last_number", 100,
		`an int till which we want to check primes`)
	debug := flag.Bool("debug", false, "a bool flag for debug")
	flag.Parse()
	// first add 2 as the first prime and then check only for odds
	if *max_check >= 2 {
		primes = append(primes, 2)
	}
	for n := 3; n <= *max_check; n += 2 {
		if f := checkIfPrime(n); f == 0 {
			primes = append(primes, n)
		}
	}
	if *debug {
		fmt.Println(primes)
	}
	fmt.Printf("Total primes below %d is %d", *max_check, len(primes))
	fmt.Println()
}

func checkIfPrime(n int) int {
	last_check := int(math.Sqrt(float64(n)))
	// numbers to check for primability is primes till root of n
	for _, p := range primes {
		if n%p == 0 {
			return p
		} else if p > last_check {
			break
		}
	}
	return 0
}
