
CLIENT_PATH = cmd/client
SERVER_PATH = cmd/server
client = client
server = server
binaries = $(client) $(server)

.PHONY: clean

build-client:
	go build -o $(client) $(CLIENT_PATH)/client.go

build-server:
	go build -o $(server) $(SERVER_PATH)/server.go

build: build-client build-server

clean:
	rm -f $(binaries)
