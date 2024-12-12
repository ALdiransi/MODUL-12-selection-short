package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text()) // Jumlah daerah

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		m, _ := strconv.Atoi(input[0]) // Jumlah rumah di daerah ini
		houses := make([]int, m)

		for j := 0; j < m; j++ {
			houses[j], _ = strconv.Atoi(input[j+1])
		}

		selectionSort(houses)

		// Cetak hasil pengurutan
		for _, house := range houses {
			fmt.Printf("%d ", house)
		}
		fmt.Println()
	}
}
