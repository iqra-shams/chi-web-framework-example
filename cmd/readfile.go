package cmd

import (
	// "net/http"
"fmt"
	"github.com/iqra-shams/chi/pkg"
	"github.com/iqra-shams/chi/utils"
)
func ProcessFile(FileData string,routines int) {
	channal := make(chan pkg.Summary)
	
	chunk := len(FileData) / routines
	startIndex := 0
	endIndex := chunk
	for iterations := 0; iterations < routines; iterations++ {
		go pkg.Counts(FileData[startIndex:endIndex], channal)
		startIndex = endIndex
		endIndex += chunk
	}
	utils.GetChunksSummary(channal , routines)
	
	fmt.Printf("total number of lines : %d \n", utils.Lines)
	fmt.Printf("total number of words : %d \n", utils.Words)
	fmt.Printf("total number of vowels: %d \n", utils.Vowels)
	fmt.Printf("total number of puncuations : %d \n", utils.Puncuations)
}
