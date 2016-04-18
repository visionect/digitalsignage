debug:
	@go-bindata -debug=true -ignore=\\.bower\\.json static/...
	@go build

bindata.go:
	@go-bindata -ignore=\\.bower\\.json static/...

release/digitalsignage_osx_64:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-s -w' -o release/digitalsignage_osx_64

release/digitalsignage_linux_64:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o release/digitalsignage_linux_64

release/digitalsignage_windows_64.exe:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-s -w' -o release/digitalsignage_windows_64.exe

build: release/digitalsignage_osx_64 release/digitalsignage_linux_64 release/digitalsignage_windows_64.exe

release:
	$(MAKE) bindata.go
	$(MAKE) build

compress_osx:
	$(MAKE) bindata.go
	$(MAKE) release/digitalsignage_osx_64
	upx -7 release/digitalsignage_osx_64

compress_linux:
	$(MAKE) bindata.go
	$(MAKE) release/digitalsignage_linux_64
	upx -7 release/digitalsignage_linux_64

compress_windows:
	$(MAKE) bindata.go
	$(MAKE) release/digitalsignage_windows_64.exe
	upx -7 release/digitalsignage_windows_64.exe

compress: compress_osx compress_linux compress_windows
