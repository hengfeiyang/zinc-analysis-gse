gse_analyzer.so: *.go
	GOOS=darwin GOARCH=amd64 go build -buildmode=plugin -o gse_analyzer.so

run: *.go
	go build &&  ./gse-zinc-plugin

clean:
	rm -f gse_analyzer.so
