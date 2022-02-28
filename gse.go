package main

import (
	"github.com/blugelabs/bluge/analysis"
	"github.com/go-ego/gse"

	"github.com/prabhatsharma/zinc/pkg/plugin"
	zincanalysis "github.com/prabhatsharma/zinc/pkg/plugin/analysis"
	"github.com/prabhatsharma/zinc/pkg/zutils"

	"github.com/hengfeiyang/gse-zinc-plugin/analyzer"
	"github.com/hengfeiyang/gse-zinc-plugin/token"
	"github.com/hengfeiyang/gse-zinc-plugin/tokenizer"
)

var ZINC_PLUGIN_TYPE = plugin.TYPE_ANALYSIS
var ZINC_PLUGIN_NAME = "gse_analyzer"

func Load() *zincanalysis.Analysis {
	t := new(zincanalysis.Analysis)
	t.Analyzer = make(map[string]*analysis.Analyzer)
	t.Analyzer["gse_standard"] = analyzer.NewStandardAnalyzer(seg)
	t.Analyzer["gse_search"] = analyzer.NewSearchAnalyzer(seg)

	t.Tokenizer = make(map[string]analysis.Tokenizer)
	t.Tokenizer["gse_standard"] = tokenizer.NewStandardTokenizer(seg)
	t.Tokenizer["gse_search"] = tokenizer.NewSearchTokenizer(seg)

	t.TokenFilter = make(map[string]analysis.TokenFilter)
	t.TokenFilter["gse_stop"] = token.NewStopTokenFilter(seg, nil)

	return t
}

var seg *gse.Segmenter

func init() {
	seg = new(gse.Segmenter)
	seg.LoadDictEmbed("zh_s")
	seg.LoadStopEmbed()
	seg.Load = true
	seg.SkipLog = false

	dataPath := zutils.GetEnv("ZINC_PLUGIN_PATH", "./plugins")
	dataPath += "/" + ZINC_PLUGIN_NAME + "/"
	if ok, _ := zutils.IsExist(dataPath); !ok {
		dataPath = "./data/"
	}
	userDict := dataPath + "dict/zh/user.txt"
	if ok, _ := zutils.IsExist(userDict); ok {
		seg.LoadDict(userDict)
	}
	stopDict := dataPath + "dict/zh/stop.txt"
	if ok, _ := zutils.IsExist(stopDict); ok {
		seg.LoadStop(stopDict)
	}
}
