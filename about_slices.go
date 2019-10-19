package go_koans

func aboutSlices() {
	fruits := []string{"apple", "orange", "mango"}

	assert(fruits[0] == "apple") // slices seem like arrays
	assert(len(fruits) == 3)     // in nearly all respects

	tastyFruits := fruits[1:3]         // we can even slice slices
	assert(tastyFruits[0] == "orange") // slices of slices also share the underlying data

	pregnancySlots := []string{"baby", "baby", "lemon"}
	assert(cap(pregnancySlots) == 3) // the capacity is initially the length

	pregnancySlots = append(pregnancySlots, "baby!")
	assert(len(pregnancySlots) == 4) // slices can be extended with append(), much like realloc in C
	assert(cap(pregnancySlots) == 6) // but with better optimizations

	pregnancySlots = append(pregnancySlots, "another baby!?", "yet another, oh dear!", "they must be Catholic")

	assert(len(pregnancySlots) == 7)  // append() can take N arguments to append to the slice
	assert(cap(pregnancySlots) == 12) // the capacity optimizations have a guessable algorithm
}
