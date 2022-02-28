package analyzer

import (
	"github.com/blugelabs/bluge/analysis"
	"github.com/go-ego/gse"

	"github.com/hengfeiyang/gse-zinc-plugin/token"
	"github.com/hengfeiyang/gse-zinc-plugin/tokenizer"
)

func NewSearchAnalyzer(seg *gse.Segmenter) *analysis.Analyzer {
	return &analysis.Analyzer{
		Tokenizer:    tokenizer.NewSearchTokenizer(seg),
		TokenFilters: []analysis.TokenFilter{token.NewStopTokenFilter(seg, nil)},
	}
}
