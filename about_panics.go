package goKoans

func divideFourBy(i int) int {
	return 4 / i
}

const divisor = 0

func aboutPanics() {
	panicRecover := func() {
		if r := recover(); r != nil {
			println("recovered")
		}
	}
	defer panicRecover()
	assert(true) // panics are exceptional errors at runtime

	n := divideFourBy(divisor)
	assert(n == 2) // panics are exceptional errors at runtime
}
