package primes

func countPrimes(n int) int {
	isPrim := make([]bool, n)
	for i := range isPrim {
		isPrim[i] = true
	}
	for i := 2; i*i < n; i++ {
		if isPrim[i] {
			for j := i * i; j < n; j += i {
				isPrim[j] = false
			}
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrim[i] {

			count++
		}
	}
	return count
}
