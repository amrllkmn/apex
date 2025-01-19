APP_NAME := apex-app
run:
	go run .

build:
	go build -o $(APP_NAME) .

clean:
	rm -rf $(APP_NAME)
