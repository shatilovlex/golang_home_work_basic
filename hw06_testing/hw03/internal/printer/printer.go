package printer

func GenerateChessboard(size int) string {
	var result string
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			if (i+j)%2 == size%2 {
				result = result + " "
			} else {
				result = result + "#"
			}
		}
		result = result + "\n"
	}
	return result
}
