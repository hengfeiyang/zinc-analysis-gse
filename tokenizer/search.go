package tokenizer

import (
	"github.com/blugelabs/bluge/analysis"
	"github.com/go-ego/gse"
)

type SearchTokenizer struct {
	seg *gse.Segmenter
}

func NewSearchTokenizer(seg *gse.Segmenter) *SearchTokenizer {
	return &SearchTokenizer{seg}
}

func (t *SearchTokenizer) Tokenize(input []byte) analysis.TokenStream {
	result := make(analysis.TokenStream, 0, len(input))
	segments := t.seg.ModeSegment(input, true)
	for _, seg := range segments {
		result = append(result, &analysis.Token{
			Term:         []byte(seg.Token().Text()),
			Start:        seg.Start(),
			End:          seg.End(),
			PositionIncr: 1,
			Type:         analysis.Ideographic,
		})
	}
	return result
}
