package goKoans

func isPrimeNumber(possiblePrime int) bool {
	for underPrime := 2; underPrime < possiblePrime; underPrime++ {
		if possiblePrime%underPrime == 0 {
			return false
		}
	}
	return true
}

func findPrimeNumbers(channel chan int) {
	for i := 2; i < 12; /* infinite loop */ i++ {
		if isPrimeNumber(i) {
			channel <- i
		}
		assert(i < 100) // i is afraid of heights
	}
}

func aboutConcurrency() {
	ch := make(chan int, 5)

	assert(len(ch) == 0) // concurrency can be almost trivial
	// your code goes here
	findPrimeNumbers(ch)

	assert(<-ch == 2)
	assert(<-ch == 3)
	assert(<-ch == 5)
	assert(<-ch == 7)
	assert(<-ch == 11)
}
