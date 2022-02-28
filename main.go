package main

import "fmt"

func main() {
	text := "《复仇者联盟3：无限战争》是全片使用IMAX摄影机拍摄制作的科幻片好吗好的."
	analysis := Load()

	// Analyzer
	tokens1 := analysis.Analyzer["gse_standard"].Analyze([]byte(text))
	fmt.Println(tokens1)

	tokens2 := analysis.Analyzer["gse_search"].Analyze([]byte(text))
	fmt.Println(tokens2)

	// Cut
	fmt.Println("cut: ", seg.Cut(text, true))
	fmt.Println("cut all: ", seg.CutAll(text))
	fmt.Println("cut for search: ", seg.CutSearch(text, true))
	fmt.Println(seg.String(text, true))

	// Segment
	segment1 := seg.Segment([]byte(text))
	fmt.Println("cut for Segment: ")
	for _, word := range segment1 {
		fmt.Println(string(word.Token().Text()))
	}
	segment2 := seg.ModeSegment([]byte(text), true)
	fmt.Println("cut for ModeSegment: ")
	for _, word := range segment2 {
		fmt.Println(string(word.Token().Text()))
	}
}
