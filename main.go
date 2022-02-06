package main

import (
	"algorithm_and_data_structure/algorithm/sort"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printTrumps(trumps []sort.Trump) {
	for i := 0; i < len(trumps); i++ {
		if i != len(trumps)-1 {
			fmt.Printf("%s%d ", trumps[i].Suit.String(), trumps[i].Value)
		} else {
			fmt.Printf("%s%d\n", trumps[i].Suit.String(), trumps[i].Value)
		}
	}
}

func main() {
	var (
		N int
		err error
		suit sort.Suit
		value int
		trumps []sort.Trump
	)
	
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	// カードの枚数読み込み
	if sc.Scan() {
		N, err = strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
	}

	// カード読み込み
	for i := 0; i < N; i++ {
		if sc.Scan() {
			str := sc.Text()
			if len(str) != 2 {
				fmt.Println("invalid input")
				os.Exit(1)
			}

			// 絵柄(suit)を判別
			switch string(str[0]) {
			case "S":
				suit = sort.S
			case "H":
				suit = sort.H
			case "C":
				suit = sort.C
			case "D":
				suit = sort.D
			}

			// 数字(value)を取得
			value, err = strconv.Atoi(string(str[1]))
			if err != nil {
				panic(err)
			}

			trumps = append(trumps, sort.Trump{Suit: suit, Value: value})
		}
	}

	// スライスをコピー
	trumps2 := make([]sort.Trump, len(trumps))
	copy(trumps2, trumps)

	// バブルソート実行
	sort.TrumpBubbleSort(trumps)
	printTrumps(trumps)
	fmt.Println("Stable")

	// 選択ソート実行
	sort.TrumpSelectionSort(trumps2)
	printTrumps(trumps2)
	if &trumps != &trumps2 {
		fmt.Println("Not stable")
	} else {
		fmt.Println("Stable")
	}
}