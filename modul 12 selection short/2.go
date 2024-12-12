func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text()) // Jumlah daerah

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		m, _ := strconv.Atoi(input[0]) // Jumlah rumah di daerah ini
		houses := make([]int, m)

		odds := []int{}
		evens := []int{}

		for j := 0; j < m; j++ {
			num, _ := strconv.Atoi(input[j+1])
			houses[j] = num
			if num%2 == 0 {
				evens = append(evens, num)
			} else {
				odds = append(odds, num)
			}
		}

		selectionSort(odds)           // Ganjil ascending
		selectionSort(evens)          // Genap ascending
		for i, j := 0, len(evens)-1; i < j; i, j = i+1, j-1 { // Reverse genap
			evens[i], evens[j] = evens[j], evens[i]
		}

		// Cetak hasil pengurutan
		for _, odd := range odds {
			fmt.Printf("%d ", odd)
		}
		for _, even := range evens {
			fmt.Printf("%d ", even)
		}
		fmt.Println()
	}
}
