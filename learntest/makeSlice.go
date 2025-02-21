package main

func makeSliceLearn() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	b = b[:cap(b)]
	printSlice(b)

	b = b[1:]
	printSlice(b)
}
