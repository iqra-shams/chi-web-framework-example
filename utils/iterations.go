package utils
import (
	"fmt"
	"github.com/iqra-shams/chi/pkg"

)
func GetChunksSummary(channal chan pkg.Summary , routines  int){
	for iterations := 0; iterations < routines; iterations++ {
	counts := <-channal
	fmt.Printf("number of lines of chunk %d: %d \n", iterations+1, counts.LineCount)
	fmt.Printf("number of words of chunk %d: %d \n", iterations+1, counts.WordsCount)
	fmt.Printf("number of vowels of chunk %d: %d \n", iterations+1, counts.VowelsCount)
	fmt.Printf("number of puncuations of chunk %d: %d \n", iterations+1, counts.PuncuationsCount)
	Lines = Lines + counts.LineCount
	Words = Words + counts.WordsCount
	Vowels = Vowels + counts.VowelsCount
	Puncuations = Puncuations + counts.PuncuationsCount
} 
}