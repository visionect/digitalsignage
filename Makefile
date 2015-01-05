build:
	@go build

debug:
	@go-bindata -debug=true -ignore=\\.bower\\.json static/...
	build

release:
	@go-bindata -ignore=\\.bower\\.json static/...
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-s' -o release/digitalsignage_osx_64
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o release/digitalsignage_linux_64
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-s' -o release/digitalsignage_windows_64.exe
