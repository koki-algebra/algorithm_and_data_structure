package sort

func InsertionSort(A []int) []int {
	for i := 1; i < len(A); i++ {
		v := A[i]
		j := i - 1
		for ; j >= 0 && A[j] > v; j-- {
			A[j+1] = A[j]
		}
		A[j+1] = v
	}

	return A
}