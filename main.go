package main

func main() {
	Solution([]int{2, 3, 1, 5})
}

func Solution(A []int) int {
	expected := ((len(A)+2) * (len(A)+1)) / 2
	sum := 0
	for _, a := range A {
		sum += a
	}
	return expected - sum
}

