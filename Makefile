build:
	@go build

debug:
	@go-bindata -debug=true -ignore=\\.bower\\.json static/...
	build

release:
	@go-bindata -ignore=\\.bower\\.json static/...
	@CGO_ENABLED=0 go build -a -ldflags '-s'
