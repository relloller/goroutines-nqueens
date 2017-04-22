/*
https://github.com/relloller/goroutines-nqueens
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

const nq_num = 14

var solsArr []int

func sumSols(nq int, arr []int) int {
	var sumATemp = 0
	arrL := len(arr)
	for i := 0; i < arrL-1; i++ {
		sumATemp += (arr[i] * 2)
	}
	if nq%2 == 1 {
		sumATemp += arr[arrL-1]
	} else {
		sumATemp += (arr[arrL-1] * 2)
	}
	return sumATemp
}

func abs_v(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func floor_half(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return (n - 1) / 2
}

func ceil_half(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return (n + 1) / 2
}

func checkColumn(a [2]int, b [2]int) bool {
	return a[1] != b[1]
}

func checkDiagonal(a [2]int, b [2]int) bool {
	return abs_v(a[0]-b[0]) != abs_v(a[1]-b[1])
}

func checkSpace(a [2]int, b [2]int) bool {
	return checkColumn(a, b) && checkDiagonal(a, b)
}

func checkSpaceEach(arr [][2]int, b [2]int) bool {
	l := len(arr)
	for i := 0; i < l; i++ {
		if !checkSpace(arr[i], b) {
			return false
		}
	}
	return true
}

func nqrec(n int, arr [][2]int) int {
	arrL := len(arr)
	if arrL == n {
		solsArr[arr[0][1]] += 1
	} else {
		for i := 0; i < n; i++ {
			if checkSpaceEach(arr, [2]int{arrL, i}) {
				nqrec(n, append(arr, [2]int{arrL, i}))
			}
		}
	}
	return solsArr[arr[0][1]]
}

func nqueens(n1 int, arr1 [][2]int) int {
	return nqrec(n1, arr1)
}

func closureCount(n int, lim int) func() int {
	var c = -1
	var nq_num = n
	tmr_i := time.Now()
	return func() int {
		c += 1
		if c == lim {
			fmt.Println("Duration:", time.Since(tmr_i), ", Solutions:", sumSols(nq_num, solsArr))
		}
		return c
	}
}

func arrLenFill(arrL int, fillVal int) []int {
	var arrT []int
	for i := 0; i < arrL; i++ {
		arrT = append(arrT, fillVal)
	}
	return arrT
}

func main() {
	var wg sync.WaitGroup
	var middlePos = floor_half(nq_num)
	solsArr = arrLenFill(ceil_half(nq_num), 0)
	goCounter := closureCount(nq_num, ceil_half(nq_num))
	goCounter()

	for i := 0; i < middlePos; i++ {
		wg.Add(1)
		go func(n int, arr [][2]int, ii int) {
			defer wg.Done()
			nqueens(n, arr)
			fmt.Println(ii, solsArr[arr[0][1]], solsArr)
			goCounter()
		}(nq_num, [][2]int{[2]int{0, i}}, i)
	}

	if nq_num%2 == 1 {
		wg.Add(1)
		go func(n int, arr [][2]int) {
			defer wg.Done()
			nqueens(n, arr)
			fmt.Println(middlePos, solsArr[arr[0][1]], solsArr)
			goCounter()
		}(nq_num, [][2]int{[2]int{0, middlePos}})
	}

	wg.Wait()
	fmt.Println("# of queens: ", nq_num)
}
