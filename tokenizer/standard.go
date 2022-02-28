package tokenizer

import (
	"github.com/blugelabs/bluge/analysis"
	"github.com/go-ego/gse"
)

type StandardTokenizer struct {
	seg *gse.Segmenter
}

func NewStandardTokenizer(seg *gse.Segmenter) *StandardTokenizer {
	return &StandardTokenizer{seg}
}

func (t *StandardTokenizer) Tokenize(input []byte) analysis.TokenStream {
	result := make(analysis.TokenStream, 0, len(input))
	segments := t.seg.Segment(input)
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
