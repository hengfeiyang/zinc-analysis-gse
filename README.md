# zinc-analysis-gse

it's a plugin of zinc to support Chinese analyzer.

Analyzer: `gse_standard` , `gse_search`

Tokenizer: `gse_standard` , `gse_search`

TokenFilter: `gse_stop`

> build has embed dictionary of `zh/s_1.txt`, `zh/stop_tokens.txt`.

you can find it: https://github.com/go-ego/gse/tree/master/data/dict

> also you can custom dictionary follow [custom user dictionary](#custom-user-dictionary)

after custom, you need restart zinc.

## gse

https://github.com/go-ego/gse

Go efficient multilingual NLP and text segmentation; support english, chinese, japanese and other.

## usage

### 1. build as go module

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildmode=plugin -o gse_analyzer.so
```

### 2. install to zinc plugin

```
mkdir -p ${ZINC_PLUGIN_PATH}/gse_analyzer
cp gse_analyzer.so ${ZINC_PLUGIN_PATH}/gse_analyzer/
cp -R data/dict    ${ZINC_PLUGIN_PATH}/gse_analyzer/
```

### 3. use gse to analyze or mappings

POST http://localhost:4080/
```
{
  "analyzer": "gse_search",
  "text": "《复仇者联盟3：无限战争》是全片使用IMAX摄影机拍摄制作的的科幻片."
}
```

## custom user dictionary

add your words append to the file `${ZINC_PLUGIN_PATH}/gse_analyzer/dict/zh/user.txt`

format:

```
分词文本  频率        词性
word    frequency   property
```

like:

```
复仇者联盟 100 n
```

## custom stop tokens

add your words append to the file `${ZINC_PLUGIN_PATH}/gse_analyzer/dict/zh/stop.txt`

format:

```
停止词
word
```

like:

```
哈哈
```

## Credit

* https://github.com/prabhatsharma/zinc
* https://github.com/blugelabs/bluge
* https://github.com/go-ego/gse
