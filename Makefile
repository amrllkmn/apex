APP_NAME := apex-app
run:
	go run ./cmd

build:
	go build -o $(APP_NAME) ./cmd

clean:
	rm -rf $(APP_NAME)
