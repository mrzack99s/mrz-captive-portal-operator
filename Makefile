build:
	go build -o bin/mrz-captive-portal-operator .

run:
	go run main.go

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/mrz-captive-portal-operator-linux-amd64 .