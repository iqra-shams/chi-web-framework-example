// mypackage/summary.go
package pkg

type Summary struct {
	LineCount        int
	WordsCount       int
	VowelsCount      int
	PuncuationsCount int
}

func Counts(data string, channal chan Summary) {
	DocCounts := Summary{}
	for _, char := range data {
		switch {
		case char == '\n':
			DocCounts.LineCount++
		case char == ' ':
			DocCounts.WordsCount++
		case char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U':
			DocCounts.VowelsCount++
		case (char < 48 && char > 32) || (char < 65 && char > 57) || (char < 97 && char > 90) || (char < 127 && char > 122):
			DocCounts.PuncuationsCount++
		}
	}

	channal <- DocCounts
}
