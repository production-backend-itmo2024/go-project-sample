LOCAL_BIN:=$(CURDIR)/bin

build_server:
	go build -o="$(LOCAL_BIN)/serv" ./cmd/server/main.go


build_client:
	go build -o="$(LOCAL_BIN)/client" ./cmd/client/main.go