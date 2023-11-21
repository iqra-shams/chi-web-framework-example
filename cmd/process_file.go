package cmd

import (
	// "net/http"

	"github.com/iqra-shams/chi/pkg"
	"github.com/iqra-shams/chi/utils"
)

func ProcessFile(FileData string, routines int) pkg.Summary {

	channal := make(chan pkg.Summary)

	chunk := len(FileData) / routines
	startIndex := 0
	endIndex := chunk
	for iterations := 0; iterations < routines; iterations++ {
		go pkg.Counts(FileData[startIndex:endIndex], channal)
		startIndex = endIndex
		endIndex += chunk
	}
	FileSummary := utils.GetChunksSummary(channal, routines)

	return FileSummary
}
