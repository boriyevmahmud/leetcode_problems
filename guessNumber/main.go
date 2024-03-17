package main

func main() {

	guessNumber()

}

func guessNumber(n int) int {

	return binarySearch(0, n)
}
func binarySearch(start, finish int) int {

	if start > finish {
		return -1
	}

	dot := (finish-start)/2 + start

	switch guess(dot) {
	case -1:
		return binarySearch(start, dot-1)
	case 1:
		return binarySearch(dot+1, finish)
	default:
		return dot
	}

}
