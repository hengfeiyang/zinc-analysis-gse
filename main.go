package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	text := "《复仇者联盟3：无限战争》是全片使用IMAX摄影机拍摄制作的科幻片好吗好的."
	analysis := Load()

	// Analyzer
	tokens1 := analysis.Analyzer["gse_standard"].Analyze([]byte(text))
	jstr1, _ := json.MarshalIndent(tokens1, "", "  ")
	fmt.Println(string(jstr1))

	tokens2 := analysis.Analyzer["gse_search"].Analyze([]byte(text))
	jstr2, _ := json.MarshalIndent(tokens2, "", "  ")
	fmt.Println(string(jstr2))
}
