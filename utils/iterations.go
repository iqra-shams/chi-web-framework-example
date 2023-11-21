package utils

import (
	"github.com/iqra-shams/chi/pkg"
)

func GetChunksSummary(channal chan pkg.Summary, routines int) pkg.Summary {

	var FileSummary pkg.Summary
	for iterations := 0; iterations < routines; iterations++ {
		counts := <-channal
		FileSummary.LineCount += counts.LineCount
		FileSummary.WordsCount += counts.WordsCount
		FileSummary.VowelsCount += counts.VowelsCount
		FileSummary.PuncuationsCount += counts.PuncuationsCount
	}
	return FileSummary
}
