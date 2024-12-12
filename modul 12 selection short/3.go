func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := []int{}

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())

		if num == -5313 { // Penanda akhir
			break
		} else if num == 0 {
			// Urutkan data dan hitung median
			insertionSort(data)
			n := len(data)
			median := data[n/2]
			if n%2 == 0 {
				median = (data[n/2-1] + data[n/2]) / 2
			}
			fmt.Println(median)
		} else {
			data = append(data, num)
		}
	}
}
