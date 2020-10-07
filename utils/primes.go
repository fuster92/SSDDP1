// AUTORES: Rafael Tolosana Calasanz
// FICHERO: primes.go
// FECHA: 01-oct-2020
// TIEMPO: 30'
// DESCRIPCIÓN: Funciones para el cálculo de primos

package utils

func IsPrime(n int) (foundDivisor bool) {
	foundDivisor = false

	for i := 2; (i < n) && !foundDivisor; i++ {
		foundDivisor = (n%i == 0)
	}
	return !foundDivisor
}

func FindPrimes(n int) (primes []int) {
	for i := 1; i < n; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}
