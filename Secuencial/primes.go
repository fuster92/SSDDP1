package main

func IsPrime(n int) (foundDivisor bool) {
	foundDivisor = false

	for i := 2; (i < n) && !foundDivisor; i++ {
		foundDivisor = (n % i == 0)
	}
	return !foundDivisor
}

func FindPrimes(n int) (primes []int){
	for i := 1; i < n; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}