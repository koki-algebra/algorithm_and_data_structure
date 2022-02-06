package sort


func SelectionSort(A []int) (count int) {
	count = 0
	for i := 0; i < len(A); i++ {
		// i 〜 len(A) - 1 で最小値を探索
		minj := i
		for j := i; j < len(A); j++ {
			if A[j] < A[minj] {
				minj = j
			}
		}
		
		// A[i] と A[minj] を交換
		tmp := A[i]
		A[i] = A[minj]
		A[minj] = tmp

		// 交換回数をインクリメント
		if i != minj {
			count++
		}
	}

	return
}