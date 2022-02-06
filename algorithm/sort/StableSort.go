package sort

type Suit int

const (
	S Suit = iota
	H
	C
	D
)

func (s Suit) String() string {
	switch s {
	case S:
		return "S"
	case H:
		return "H"
	case C:
		return "C"
	case D:
		return "D"
	default:
		return "Unknown"
	}
}

type Trump struct {
	Suit Suit
	Value int
}

func TrumpBubbleSort(C []Trump) {
	for i := 0; i < len(C); i++ {
		for j := len(C)-1; j > i; j-- {
			if C[j].Value < C[j-1].Value {
				tmp := C[j]
				C[j] = C[j-1]
				C[j-1] = tmp
			}
		}
	}
}

func TrumpSelectionSort(C []Trump) {
	for i := 0; i < len(C); i++ {
		minj := i
		for j := i; j < len(C); j++ {
			if C[j].Value < C[minj].Value {
				minj = j
			}
		}
		tmp := C[i]
		C[i] = C[minj]
		C[minj] = tmp
	}
}