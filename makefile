all: setup
	~/go/bin/swag init
	go build

run: all
	LISTEN_ADDR=:8080 CONNECTION_STRING=db.db HOST_URL=localhost:8080 ./web_practicum

setup:
	mkdir -p ./qr-codes
