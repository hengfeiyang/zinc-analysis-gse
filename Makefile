gse_analyzer.so: *.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -buildmode=plugin -o gse_analyzer.so

run: *.go
	go build &&  ./gse-zinc-plugin

clean:
	rm -f gse_analyzer.so
