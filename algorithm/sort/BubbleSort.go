package sort

func BubbleSort(A []int) (count int) {
	count = 0
	flag := true

	for i := 0; flag; i++ {
		flag = false
		for j := len(A)-1; j >= i+1; j-- {
			if A[j] < A[j-1] {
				// 隣接要素を交換する
				tmp := A[j]
				A[j] = A[j-1]
				A[j-1] = tmp

				flag = true
				count++
			}
		}
	}

	return
}